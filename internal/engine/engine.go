package engine

import (
	v1 "cld-quicker/internal/api/v1"
	"cld-quicker/pkg"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CreateServer  The method for creating the service
func CreateServer() {
	e := gin.New()
	apiV1 := v1.NewApiV1()
	apiV1.LaunchOnline(e)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", pkg.Config.App.Port),
		Handler: e,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err := server.Shutdown(timeoutCtx)
	if err != nil {
		zap.L().Info(err.Error())
	}
}
