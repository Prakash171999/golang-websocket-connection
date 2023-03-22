package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func ws(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	ws, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("Upgrade Error:", err)
		return
	}
	defer ws.Close()

	//The for loop is the main loop of the WebSocket handler function. It reads data from the WebSocket connection ws and then writes the same data back to the connection.
	for {
		//read data from ws
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read Message Error:", err)
			break
		}
		log.Printf("Recieved Message: %s", message)

		//write ws data
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write Message:", err)
			break
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/ws", ws)
	r.Run(":8080")
}
