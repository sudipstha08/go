package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Running websocket 🚌 ")

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		log.Println("New connection")

		s.Join("chat")
		return nil
	})

	server.OnEvent("/chat messsage", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		s.Emit("Joined")
		return "recv " + msg
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)

	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
