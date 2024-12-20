package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"liliya/api/logic"
	"liliya/api/plugin"
	"liliya/models"
	"liliya/server/ServerMode"
	"liliya/utils"
	"net/http"
	"strings"
	"time"
)

type Controller struct {
	logic  *logic.Logic
	plugin *plugin.Plugin
}

func NewController(logic *logic.Logic, plugin *plugin.Plugin) *Controller {
	return &Controller{logic: logic, plugin: plugin}
}

func (con *Controller) InitHTTPRoutes(env *models.Env) *gin.Engine {
	ginSetMode(env.ServerMode)
	router := gin.Default()
	allowOrigins := strings.Split(env.Domain, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{http.MethodPut, http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Status", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin",
			utils.HeaderAuthorization, utils.HeaderClientRequestId},
		ExposeHeaders: []string{"Content-Length", utils.HeaderTimestamp,
			utils.HeaderClientRequestId, utils.HeaderRequestId},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	coinGecko := router.Group("/api/v1")
	{
		coin := coinGecko.Group("/coins")
		{
			coin.GET("/:currency/:page", con.getAllTokens)
		}
		market := coinGecko.Group("/market")
		{
			market.GET("/:id", con.getAllMarketById)
		}
		category := coinGecko.Group("/category")
		{
			category.GET("/:id", con.getAllCategoriesById)
			category.GET("", con.getAllCategories)
		}
	}
	return router
}

func ginSetMode(serverMode string) {
	if serverMode == ServerMode.DEVELOPMENT {
		gin.SetMode(gin.ReleaseMode)
	}
}
