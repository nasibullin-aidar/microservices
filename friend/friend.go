package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func main() {
	serviceName := "friend"

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)

	gin.SetMode(gin.ReleaseMode)

	h := http.NewServeMux()

	h.HandleFunc("/healthz", handler())
	h.HandleFunc("/readyz", handler())

	addr := "0.0.0.0:8080"

	s := http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Println("err listen and serve", err)
			return
		}
	}()

	fmt.Printf("listening %s addr: %s \n", serviceName, addr)

	<-ctx.Done()

	_ = s.Shutdown(ctx)

	fmt.Println("stop ", serviceName)
}

func handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
