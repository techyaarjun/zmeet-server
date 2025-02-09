package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"zmeet/pkg/pion"
	"zmeet/pkg/store"
	"zmeet/pkg/user"
)

func ListAllUsers(c *gin.Context, s *store.Store) {
	users := s.GetAllZMeetUsers()
	if len(users) == 0 {
		users = []*user.ZMeetUser{}
	}

	userResponses := make([]map[string]interface{}, 0, len(users))
	for _, u := range users {
		userResponses = append(userResponses, map[string]interface{}{
			"id":   u.ID(),
			"name": u.Name(),
		})
	}

	c.JSON(http.StatusOK, userResponses)
}

func IsReady(pc *pion.PC, u *user.ZMeetUser) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if pc.ICEConnected() && pc.ConnectionState() {
				u.SetConnected(true)
				return
			}
		}
	}
}
