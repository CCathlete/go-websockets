package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

// For handling incoming messages.
var wsChan = make(chan WsPayload)

var clients = make(map[WebSocketConnection]string)

var views = jet.NewSet(
	// Setting up a filesystem with root at projec_root/html.
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// Upgrading a connection to websocket.
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
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

// We want to inherit all fields and methods & maybe add methods of our own so we need to wrap in a struct.
type WebSocketConnection struct {
	*websocket.Conn
}

// The payload of the ws request body.
type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"` // Ommitting from JSON
}

// ---------------------------------------------------------------------
// ------------- Defining handlers as handlerFuncs. --------------------
// ---------------------------------------------------------------------

// When the /ws endpoint is hit we trigger a goroutine that listens for messages in a ws and send them to the ws channel, which is being read by another goroutine that runs the ReadFromWsChannel function.
var WsEndpoint http.HandlerFunc = func(
	w http.ResponseWriter, r *http.Request,
) {

	ws, err := connectionUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected to ws endpoint.")

	response := WsJsonResponse{}
	response.Message = `<em><small>Connected to server</small></em>`

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn)

}

// Listens for ws messages and send them to the ws channel.
func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: %v\n", r)
		}
	}()

	var payload WsPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// No payload, do nothing.
		} else {
			//
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ReadFromWsChannel() {
	response := WsJsonResponse{}

	// Getting a list of users and broacasting it back.
	for {
		event := <-wsChan

		switch event.Action {
		case "username":
			// For each event we're reading from the channel, we're registering the connection and username in the clients map.
			clients[event.Conn] = event.Username
			response.ConnectedUsers = refreshUsersList()
			response.Action = "list_users"
			log.Println("Reading from ws channel. Response: ", response)
			broadcastToAll(response)
		case "left":
			response.Action = "list_users"
			delete(clients, event.Conn)
			response.ConnectedUsers = refreshUsersList()
			log.Println("Reading from ws channel. Response: ", response)
			broadcastToAll(response)
		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s",
				event.Username, event.Message)
			broadcastToAll(response)
		}
	}
}

// Refreshes the users list that are currently registered in the clients map.
func refreshUsersList() (users []string) {

	for _, username := range clients {
		if username != "" {
			users = append(users, username)
		}
	}
	sort.Strings(users)

	return
}

func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			// If we can't send a response to the client we close the connection.
			log.Println("Error broadcasting to client.")
			client.Close()
			delete(clients, client)
		}
	}
}
