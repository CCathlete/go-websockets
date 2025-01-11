package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	// Setting up a filesystem with root at projec_root/html.
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var connectionUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// For some reason the teacher wanted to bypass the request origin being allowed to be sent from localhost only.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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

// We convert a function with the appropriate signature to a handlerFunc type which is a funciton that satisfies the http.Handler interface.

var Home http.HandlerFunc = func(
	w http.ResponseWriter, r *http.Request,
) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
		return
	}

}

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `josn:"message_type"`
}

// Upgrading connection to websocket.
var WsEndpoint http.HandlerFunc = func(
	w http.ResponseWriter, r *http.Request,
) {

	ws, err := connectionUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to ws endpoint.")

	var response WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

}
