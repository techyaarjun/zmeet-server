package gin

import (
	"github.com/gin-gonic/gin"
	"zmeet/pkg/api"
	"zmeet/pkg/store"
)

func Initilize(store *store.Store) {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/ping", api.HandlePing)

	handshake := r.Group(api.HandShake)
	handshake.GET("/offer/:id", func(context *gin.Context) {
		id := context.Param("id")
		api.Offer(id, store, context)
	})
	handshake.POST("/answer/:id", func(context *gin.Context) {
		id := context.Param("id")
		api.Answer(id, store, context)
	})
	handshake.POST("/ice-candidate/:id", func(context *gin.Context) {
		id := context.Param("id")
		api.ICECandidate(id, store, context)
	})

	users := r.Group(api.Users)
	users.GET("", func(c *gin.Context) {
		api.ListAllUsers(c, store)
	})

	panic(r.Run())
}
