package routes

import (
	"inssa_club_waitlist_backend/cmd/server/controllers"

	"github.com/gin-gonic/gin"
)

// RouteInfo is a struct for a route
type RouteInfo struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// GetRoutes is a function which returns the information of routes
func GetRoutes() []RouteInfo {
	controller := controllers.Controller{}
	routeInfo := []RouteInfo{
		{Method: "POST", Path: "/interest", Handler: controller.AddInterest},
		{Method: "DELETE", Path: "/interest", Handler: controller.DeleteInterest},
	}
	return routeInfo
}
