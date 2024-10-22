package groupie

import (
	"net/http"
	"os"
)

// Handler css files
func CssHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}
	filePath := "res/css/" + req.URL.Path[len("/css/"):]
	if filePath == "res/css/" {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	_, err := os.Stat(filePath)
	if err != nil {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	http.StripPrefix("/css/", http.FileServer(http.Dir("res/css/"))).ServeHTTP(res, req)
}
