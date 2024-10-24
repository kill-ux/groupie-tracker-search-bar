package groupie

import (
	"net/http"
	"strconv"
	"strings"
)

func SearchName(str, sub string) bool {
	return strings.Contains(strings.ToLower(str), sub)
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
	return yearData == yearClient
}

func SearchLoop(slice []string, value string) bool {
	for _, element := range slice {
		if strings.Contains(strings.ToLower(element), value) {
			return true
		}
	}
	return false
}

func HandelFilter(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}
	action := req.FormValue("action")
	req.ParseForm()
	if len(req.Form) > 0 && (len(action) == 0 || (len(action) > 0 && action != "Filter" && action != "Search")) {
		Error(res, 400, "Oops!! Bad Request")
		return
	}

	if action == "Search" {
		Data.Filters = nil
		keyword := req.FormValue("keyword")
		if _, ok := req.Form["keyword"]; !ok {
			Error(res, 400, "Oops!! Bad Request")
			return
		}
		keyword = strings.ReplaceAll(keyword, ", ", "-")
		keyword = strings.ToLower(keyword)
		keyword = strings.TrimSpace(keyword)
		for i, artist := range Data.Arts {
			if SearchName(artist.Name, keyword) || SearchYear(artist.CreationDate, keyword) ||
				SearchFirstAlbum(artist.FirstAlbum, keyword) || SearchLoop(artist.Members, keyword) ||
				SearchLoop(Data.DataLocals["index"][i].Locations, keyword) {
				Data.Filters = append(Data.Filters, artist)
			}
		}

	} else if action == "Filter" {
		Data.Filters = nil
		FromCreationDate := req.FormValue("FromCreationDate")
		ToCreationDate := req.FormValue("ToCreationDate")
		FromFirsetAlbum := req.FormValue("FromFirsetAlbum")
		ToFirsetAlbum := req.FormValue("ToFirsetAlbum")
		local := req.FormValue("local")
		local = strings.ReplaceAll(local, ", ", "-")
		local = strings.ToLower(local)
		Members := req.Form["members"]
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
		for i, artist := range Data.Arts {

			FirstAlbum, err := strconv.Atoi(artist.FirstAlbum[6:])
			if err != nil {
				Error(res, 400, "Oops!! Bad Request")
				return
			}
			if (artist.CreationDate >= Form && artist.CreationDate <= To) && (FirstAlbum >= FormAlbum && FirstAlbum <= ToAlbum) {
				for _, nMembers := range Members {
					num, err := strconv.Atoi(nMembers)
					if err != nil {
						Error(res, 400, "Oops!! Bad Request")
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
