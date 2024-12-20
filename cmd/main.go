package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"liliya/api/controller"
	"liliya/api/logic"
	"liliya/api/plugin"
	"liliya/initialize"
	"liliya/models"
	"liliya/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.Info("start server")

	type serverInstance server.Server
}

func init() {
	logrus.Info("start init server")

	var serverInstance server.Server

	if err := initialize.LoadConfiguration(); err != nil {
		logrus.Fatal(err.Error())
	}
	if err := initialize.RunLogger(); err != nil {
		logrus.Fatal(err.Error())
	}

	plugins := plugin.NewPlugin(&initialize.Config, &initialize.Env)
	logics := logic.NewLogic(plugins)
	controllers := controller.NewController(logics, plugins)

	go runServer(serverInstance, controllers, &initialize.Env, initialize.Config.Server)
	runChannelStopServer()

	serverInstance.Shutdown(context.Background())
}

func runServer(server server.Server, controller *controller.Controller, env *models.Env, cfg models.ServerConfig) {
	ginEngine := controller.InitHTTPRoutes(env)
	if err := server.Run(cfg.Port, ginEngine); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", nil, err.Error())
	}
}
func runChannelStopServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}
