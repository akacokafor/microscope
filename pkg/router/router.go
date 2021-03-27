package router

import (
	"net/http"

	"github.com/akacokafor/microscope/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Configure(routePrefix, gocraftNamespace string, gocraftRedisPool *redis.Pool) http.Handler {
	return api.NewHTTPRouter(routePrefix, true, api.GoCraftOptions{
		Namespace: gocraftNamespace,
		Pool:      gocraftRedisPool,
	})
}

func ConfigureRoutes(router gin.IRoutes, routePrefix, gocraftNamespace string, gocraftRedisPool *redis.Pool) {
	api.RegisterRoutesOnGin(router, routePrefix, true, api.GoCraftOptions{
		Namespace: gocraftNamespace,
		Pool:      gocraftRedisPool,
	})
}
