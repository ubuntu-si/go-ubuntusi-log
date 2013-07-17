package rest

import (
	"github.com/trevex/golem"
	"net/http"
)

// Create room manager
var myroommanager = golem.NewRoomManager()

func join(conn *golem.Connection) {
	myroommanager.Join("updates", conn)
}

// On leave, leave the specified room.
func leave(conn *golem.Connection) {
	myroommanager.Leave("updates", conn)
}

func Notify(msg string) {
	myroommanager.Emit("updates", "update", msg)
}

func connClose(conn *golem.Connection) {
	myroommanager.LeaveAll(conn)
}

func RegisterWS() {

	// Create a router
	myrouter := golem.NewRouter()
	myrouter.On("join", join)
	myrouter.On("leave", leave)
	myrouter.OnClose(connClose)

	// Serve the public files
	//http.Handle("/", http.FileServer(http.Dir("./public")))
	// Handle websockets using golems handler
	http.HandleFunc("/ws", myrouter.Handler())

}
