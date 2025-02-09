package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"zmeet/pkg/logger"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleSocketConnection(c *gin.Context, logger *logger.Logger) {

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Debug("WebSocket upgrade failed")
		return
	}
	defer func() {
		panic(conn.Close())
	}()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			logger.Debug("WebSocket read failed")
			break
		}
		log.Printf("recv: %v %s", mt, message)
		//err = c.WriteMessage(mt, message)
		//if err != nil {
		//	log.Println("write:", err)
		//	break
		//}
	}
}
