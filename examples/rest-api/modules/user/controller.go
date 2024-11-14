package user

import (
	"net/http"

	"github.com/huboh/vara"
	"github.com/huboh/vara/pkg/modules/json"
)

type controller struct {
	users *service
	json  *json.Service
}

func newController(s *service) *controller {
	return &controller{
		users: s,
	}
}

func (c *controller) Config() *vara.ControllerConfig {
	return &vara.ControllerConfig{
		Pattern: "/users",
		RouteConfigs: []*vara.RouteConfig{
			{
				Pattern: "/",
				Method:  http.MethodPost,
				Handler: http.HandlerFunc(c.handleGetUser),
			},
			{
				Pattern: "/id",
				Method:  http.MethodPost,
				Handler: http.HandlerFunc(c.handleGetUsers),
			},
		},
	}
}

func (c *controller) handleGetUser(w http.ResponseWriter, r *http.Request) {
	c.json.Write(w, json.Response{
		Data: c.users.getUser(),
	})
}

func (c *controller) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	c.json.Write(w, json.Response{
		Data: map[string]string{"message": "user controller is working!"},
	})
}
