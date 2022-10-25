package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin/Databases/Mysql/Famey"
	"gin/Router"
	"gin/pkg/setting"

	"github.com/apex/gateway"
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

// https://ibbv4y3f1j.execute-api.us-east-1.amazonaws.com/default/hello-world
// 本地测试 停止 command/control + c
func main() {
	defer Famey.DB.Close()

	if inLambda() {

		fmt.Println("running aws lambda in aws")
		// fmt.Println(os.Getenv("AWS_REGION"))
		log.Fatal(gateway.ListenAndServe(":8080", Router.InitRouter()))

	} else {

		srv := &http.Server{
			Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
			Handler:        Router.InitRouter(),
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}

		fmt.Printf("port %d", setting.HTTPPort)
		srv.ListenAndServe()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")
	}
}
