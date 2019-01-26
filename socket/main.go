package main

import (
	"log"
	"net/http"
	"time"

	"github.com/googollee/go-socket.io"
)

//Custom server which basically only contains a socketio variable
//But we need it to enhance it with functions
type customServer struct {
	Server *socketio.Server
}

type member struct {
	index uint
	id    string
}

type onlinemembers struct {
	members map[string]member
	count   int
}

var om onlinemembers

//Header handling, this is necessary to adjust security and/or header settings in general
//Please keep in mind to adjust that later on in a productive environment!
//Access-Control-Allow-Origin will be set to whoever will call the server
func (s *customServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.Server.ServeHTTP(w, r)
}

func main() {

	om.members = make(map[string]member)

	//get/configure socket.io websocket for clients
	ioServer := configureSocketIO()

	wsServer := new(customServer)
	wsServer.Server = ioServer

	//HTTP settings
	println("Core Service is listening on port 8081...")
	http.Handle("/socket.io/", wsServer)
	http.ListenAndServe(":8081", nil)
}

func configureSocketIO() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	//Client connects to server
	server.On("connection", func(so socketio.Socket) {

		//What will happen as soon as the connection is established:
		so.On("connection", func(msg string) {
			so.Join("clients")

			m := member{uint(om.count), so.Id()}

			addMember(&om, m)

			println(so.Id() + " joined clients.")

			//In case you want to send a custom emit directly after the client connected.
			//If you fire an emit directly after the connection event it won't work therefore you need to wait a bit
			//In this case two seconds.
			ticker := time.NewTicker(2 * time.Second)
			go func() {
				for {
					select {
					case <-ticker.C:
						so.Emit("onlinemembers", len(om.members))
						ticker.Stop()
						return
					}
				}
			}()
		})

		//What will happen if clients disconnect
		so.On("disconnection", func() {
			log.Println(so.Id() + " on disconnect")
			removeMember(&om, so.Id())
			it(&om)
		})

		//Custom event as example
		so.On("hello", func(msg string) {
			log.Println("received request (hello): " + msg)

			server.BroadcastTo("clients", so.Id()+" : ", msg)

			sm := "How can I help you?"

			so.Emit("Hi", sm)

			server.BroadcastTo("clients", "server : ", sm)
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}

func addMember(o *onlinemembers, m member) {

	o.members[m.id] = m

	o.count++

	log.Printf("member : " + m.id + " added")

	it(o)
}

func removeMember(o *onlinemembers, id string) {

	delete(o.members, id)

	log.Printf("member : " + id + " removed")

	it(o)
}

func it(o *onlinemembers) {
	for k, v := range o.members {
		log.Printf("key : %s, value : %s", k, v.id)
	}
}
