package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"net/http"
	"time"

	"k8s-deploy/config"
	"k8s-deploy/middleware"
	"k8s-deploy/model"
	"k8s-deploy/pkg/kubectl"
	"k8s-deploy/pkg/logger"
	"k8s-deploy/pkg/prometheus"
	"k8s-deploy/pkg/swagger"
	"k8s-deploy/query"
	"k8s-deploy/routes"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate gengin k8s-deploy.api
//go:generate gen-swagger --local_api= k8s-deploy.api

func main() {
	flag.Parse()

	configFile := fmt.Sprintf("etc/k8s-deploy-%s.yaml", config.Env)

	e := gin.New()
	e.Use(
		middleware.RequestId,
		middleware.RequestLog,
		gin.Recovery(),
		middleware.CorsMiddleware,
		middleware.ApiHitRecord,
	)

	pprof.Register(e)
	config.Setup(configFile)
	logger.Setup()
	kubectl.Setup()
	prometheus.Setup(e)
	//redis.Setup()
	model.Setup()
	query.SetDefault(model.DB())
	routes.Setup(e)
	swagger.Setup(e)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ServerConf.Host, config.ServerConf.Port),
		Handler: e,
	}
	go func() {
		logrus.Info("listen to ", server.Addr)
		server.ListenAndServe()
	}()
	wait := config.RegisterCloseFn(func() {
		defer logrus.Warning("closed api server")

		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()
		server.Shutdown(ctx)
	})
	wait()
}
