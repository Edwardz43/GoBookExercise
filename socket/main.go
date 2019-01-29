package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

//Custom server which basically only contains a socketio variable
//But we need it to enhance it with functions
type customServer struct {
	Server *socketio.Server
}

type member struct {
	id   string
	name string
}

type onlinemembers struct {
	members map[string]member
}

type Post struct {
	ID  string
	Msg string
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

			om.AddMember(so.Id())

			o := om.members[so.Id()]

			server.BroadcastTo("clients", "newmember", o.name+" joined clients.")

			//In case you want to send a custom emit directly after the client connected.
			//If you fire an emit directly after the connection event it won't work therefore you need to wait a bit
			//In this case two seconds.
			ticker := time.NewTicker(500 * time.Millisecond)
			go func() {
				for {
					select {
					case <-ticker.C:
						so.Emit("onlinemembers", len(om.members))
						server.BroadcastTo("clients", "onlinemembers", len(om.members))
						ticker.Stop()
						return
					}
				}
			}()
		})

		//What will happen if clients disconnect
		so.On("disconnection", func() {
			om.RemoveMember(so.Id())
			server.BroadcastTo("clients", "onlinemembers", len(om.members))
		})

		//Custom event as example
		so.On("job", func(msg string) {
			log.Println("job : " + msg)
			server.BroadcastTo("clients", "onlinemembers", om.members[so.Id()].name+" : "+msg)
		})

		so.On("post", func(msg string) {

			name := om.members[so.Id()].name

			data := Post{name, msg}

			log.Println("member : " + data.ID + " => " + data.Msg)

			server.BroadcastTo("clients", "post", data)
		})

	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}

func (o *onlinemembers) AddMember(id string) {

	name := fmt.Sprintf("member%d", len(om.members)-1)

	m := member{id, name}

	o.members[m.id] = m

	log.Printf("member : " + m.name + " added")

	//o.Foreach()
}

func (o *onlinemembers) RemoveMember(id string) {

	delete(o.members, id)

	//log.Printf("member : " + id + " removed")

	//o.Foreach()
}

func (o *onlinemembers) Foreach() {
	for k, v := range o.members {
		log.Printf("key : %s, value : %s", k, v.name)
	}
}

func timer(t chan string) {
	t <- string(time.Now().Format("2019-05-14 12:23:56"))
}
