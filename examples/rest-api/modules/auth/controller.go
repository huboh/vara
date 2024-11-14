package auth

import (
	"net/http"

	"github.com/huboh/vara"
	"github.com/huboh/vara/pkg/modules/json"
)

type controller struct {
	auth *service
	json *json.Service
}

func newController(s *service, j *json.Service) *controller {
	return &controller{
		auth: s,
		json: j,
	}
}

func (c *controller) Config() *vara.ControllerConfig {
	return &vara.ControllerConfig{
		Pattern: "/auth",
		RouteConfigs: []*vara.RouteConfig{
			{
				Pattern:  "/signin",
				Method:   http.MethodPost,
				Handler:  http.HandlerFunc(c.handleSignin),
				Metadata: map[string]string{},
			},
			{
				Pattern:  "/signup",
				Method:   http.MethodPost,
				Handler:  http.HandlerFunc(c.handleSignup),
				Metadata: map[string]string{},
			},
		},
		GuardConstructors: []vara.GuardConstructor{
			newGuard,
		},
	}
}

func (c *controller) handleSignin(w http.ResponseWriter, r *http.Request) {
	c.json.Write(w, json.Response{
		Data: c.auth.signin(),
	})
}

func (c *controller) handleSignup(w http.ResponseWriter, r *http.Request) {
	c.json.Write(w, json.Response{
		Data: c.auth.signup(),
	})
}
