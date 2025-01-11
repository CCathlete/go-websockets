package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	// Setting up a filesystem with root at projec_root/html.
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap,
) (err error) {

	// Loading the template from our filesystem.
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return
	}

	// Executing the template.
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
