package main

import (
	"inssa_club_waitlist_backend/cmd/server/docs"
	"inssa_club_waitlist_backend/cmd/server/middlewares"
	"inssa_club_waitlist_backend/cmd/server/routes"
	"inssa_club_waitlist_backend/cmd/server/utils"
	"inssa_club_waitlist_backend/configs"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRoutes(engine *gin.Engine) {
	IS_ENABLE_SWAGGER := configs.Envs["IS_ENABLE_SWAGGER"]
	if IS_ENABLE_SWAGGER == "true" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	for _, controller := range routes.GetRoutes() {
		engine.Handle(controller.Method, controller.Path, controller.Handler)
	} // setup routes
}

func setupDocuments() {
	docs.SwaggerInfo.Title = "waitlist-api.inssa.club"
	docs.SwaggerInfo.Description = "The REST API for waitlist service of api.inssa.club"
	docs.SwaggerInfo.Host = "api.inssa.club"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.BasePath = "/waitlist"
	docs.SwaggerInfo.Schemes = []string{"https"}
}

func runServer(engine *gin.Engine) {
	IS_SERVERLESS := configs.Envs["IS_SERVERLESS"]
	PORT := ":" + configs.Envs["SERVER_PORT"]

	if IS_SERVERLESS == "true" {
		gateway.ListenAndServe(PORT, engine)
	} else {
		engine.Run(PORT)
	}
}

func main() {
	utils.InitDB()
	engine := gin.New()
	middlewares.Setup(engine)
	setupDocuments()
	setupRoutes(engine)
	runServer(engine)
}
