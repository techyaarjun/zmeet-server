package gin

import (
	"github.com/gin-gonic/gin"
	"zmeet/pkg/api"
	"zmeet/pkg/store"
	"zmeet/pkg/ws"
)

func Initilize(store *store.Store) {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", api.HandlePing)
	r.GET("/ws", func(c *gin.Context) {
		ws.HandleSocketConnection(c, store.CustomLogger())
	})

	handshake := r.Group(api.HandShake)
	handshake.GET("/offer/:id", func(context *gin.Context) {
		id := context.Param("id")
		api.Offer(id, store, context)
	})
	handshake.POST("/answer/:id", func(context *gin.Context) {
		id := context.Param("id")
		api.Answer(id, store, context)
	})

	panic(r.Run())
}
