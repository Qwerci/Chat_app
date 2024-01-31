package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrade web socket connection

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(c *gin.Context){
	// Upgrade get request to websocket protocol
	ws, err :=upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil{
		log.Print("upgrade:", err)
		return
	}

	defer ws.Close()

	for {
		// read data from ws
		mt, message, err := ws.ReadMessage()
		if err != nil{
			log.Print("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		// write ws data
		err = ws.WriteMessage(mt, message)
		if err != nil{
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	fmt.Println("Websocket Server!")

	r := gin.Default()
	r.GET("/ws", Handler)
	r.Run(":8448")
}