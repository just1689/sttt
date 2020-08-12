package main

import (
	"fmt"
	"github.com/just1689/sttt/api"
	"github.com/just1689/sttt/web"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	listen := getListen()
	api.RunQuickGame()
	logrus.Println("listening on", listen)
	web.Setup(listen)
}

func getListen() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprint(":", port)
}
