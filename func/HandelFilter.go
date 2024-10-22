package groupie

import (
	"net/http"
	"strconv"
	"strings"
)

func SearchName(str, sub string) bool {
	if strings.Contains(strings.ToLower(str), sub) {
		return true
	}
	return false
}

func SearchYear(yearData int, yearClient string) bool {
	yearC, err := strconv.Atoi(yearClient)
	if err != nil {
		return false
	}
	if yearC == yearData {
		return true
	}
	return false
}

func SearchFirstAlbum(yearData, yearClient string) bool {
	if yearData == yearClient {
		return true
	}
	return false
}

func SearchMember(Members []string, Member string) bool {
	for _, meb := range Members {
		if strings.Contains(strings.ToLower(meb), Member) {
			return true
		}
	}
	return false
}

func SearchLocation(Locations []string, local string) bool {
	for _, loc := range Locations {
		if strings.Contains(strings.ToLower(loc), local) {
			return true
		}
	}
	return false
}

func HandelFilter(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}
	search := req.FormValue("Search")
	filter := req.FormValue("filter")
	if search == "Search" {
		Data.Filters = nil
		keyword := req.FormValue("keyword")
		keyword = strings.ReplaceAll(keyword, ", ", "-")
		keyword = strings.ToLower(keyword)
		for i, artist := range Data.Arts {
			if SearchName(artist.Name, keyword) || SearchYear(artist.CreationDate, keyword) ||
				SearchFirstAlbum(artist.FirstAlbum, keyword) || SearchMember(artist.Members, keyword) || SearchLocation(Data.DataLocals["index"][i].Locations, keyword) {
				Data.Filters = append(Data.Filters, artist)
			}
		}

	} else if filter == "Filter" {
		Data.Filters = nil
		FromCreationDate := req.FormValue("FromCreationDate")
		ToCreationDate := req.FormValue("ToCreationDate")
		FromFirsetAlbum := req.FormValue("FromFirsetAlbum")
		ToFirsetAlbum := req.FormValue("ToFirsetAlbum")
		local := req.FormValue("local")
		local = strings.ReplaceAll(local, ", ", "-")
		local = strings.ToLower(local)
		req.ParseForm()
		Members := req.Form["members"]
		for i, artist := range Data.Arts {
			Form, err := strconv.Atoi(FromCreationDate)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			To, err := strconv.Atoi(ToCreationDate)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}

			FormAlbum, err := strconv.Atoi(FromFirsetAlbum)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			ToAlbum, err := strconv.Atoi(ToFirsetAlbum)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			FirstAlbum, err := strconv.Atoi(artist.FirstAlbum[6:])
			if err != nil {
				for _, nMembers := range Members {
					num, err := strconv.Atoi(nMembers)
					if err != nil {
						Error(res, 400, "Oops!! Bade Request")
						return
					}
					if num == len(artist.Members) {
						for _, location := range Data.DataLocals["index"][i].Locations {
							if strings.Contains(location, local) {
								Data.Filters = append(Data.Filters, artist)
								break
							}
						}
					}
				}
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			if (artist.CreationDate >= Form && artist.CreationDate <= To) && (FirstAlbum >= FormAlbum && FirstAlbum <= ToAlbum) {
				for _, nMembers := range Members {
					num, err := strconv.Atoi(nMembers)
					if err != nil {
						Error(res, 400, "Oops!! Bade Request")
						return
					}
					if num == len(artist.Members) {
						for _, location := range Data.DataLocals["index"][i].Locations {
							if strings.Contains(location, local) {
								Data.Filters = append(Data.Filters, artist)
								break
							}
						}
					}
				}
				if len(Members) == 0 {
					for _, location := range Data.DataLocals["index"][i].Locations {
						if strings.Contains(location, local) {
							Data.Filters = append(Data.Filters, artist)
							break
						}
					}
				}
			}

		}
	} else {
		Data.Filters = Data.Arts
	}
}
