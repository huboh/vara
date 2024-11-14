package auth

import (
	"github.com/huboh/vara"
)

type guard struct{}

func newGuard() *guard {
	return &guard{}
}

func (g *guard) Allow(gCtx vara.GuardContext) (bool, error) {
	var isPublicRoute bool

	if gCtx.RouteConfig.Metadata != nil {
		// get metadata from route and determine if its a public route
		isPublicRoute = true
	}

	return isPublicRoute, nil
}
