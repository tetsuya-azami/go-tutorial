package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-tutorial/chapter8/api"
	"go-tutorial/chapter8/app/models"
	"go-tutorial/chapter8/configs"
	"go-tutorial/chapter8/controllers"
	"go-tutorial/chapter8/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

func main() {
	if err := models.SetDatabase(models.InstanceMySQL); err != nil {
		logger.Fatal(err.Error())
	}
	router := gin.Default()
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Swaggerの準備
	if configs.Config.IsDevelopment() {
		swaggerJson, _ := json.Marshal(swagger)
		var SwaggerInfo = &swag.Spec{
			InfoInstanceName: "swagger",
			SwaggerTemplate:  string(swaggerJson),
		}
		swag.Register(SwaggerInfo.InfoInstanceName, SwaggerInfo)
		router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// APIのルーティング
	apiGroup := router.Group("/api")
	{
		v1 := apiGroup.Group("/v1")
		{
			v1.Use(middleware.OapiRequestValidator(swagger))
			albumHandler := &controllers.AlbumHandler{}
			api.RegisterHandlers(v1, albumHandler)
		}
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	defer logger.Sync()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}
	<-ctx.Done()
	logger.Info("Shutdown")
}
