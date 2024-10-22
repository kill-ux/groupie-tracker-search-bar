package groupie

import (
	"net/http"
	"strings"
)

var Data = &Page{}

// Home page contain data about artits
func HandelHome(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
		return
	}
	if req.URL.Path != "/" {
		if strings.ContainsRune(req.URL.Path[1:], '/') {
			http.Redirect(res, req, "/notFound", http.StatusFound)
			return
		}
		Error(res, 404, "Oops!! Page Not Found")
		return
	}
	Data.DataList = map[any]string{}
	for _, artist := range Data.Arts {
		Data.DataList[artist.Name] = " - artist/band"
		Data.DataList[artist.CreationDate] = " - creation date"
		Data.DataList[artist.FirstAlbum] = " - first album date"
		for _, member := range artist.Members {
			Data.DataList[member] = " - member"
		}
	}
	for _, value := range Data.DataLocals["index"] {
		for _, local := range value.Locations {
			Data.DataList[local] = " - locations"
		}
	}

	HandelFilter(res, req)
	RenderPage("index", res)
}
