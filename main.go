package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Main(); err != nil {
		fmt.Print("error: ", err)
		os.Exit(1)
	}
}

func Main() error {
	r := gin.Default()
	r.GET("/access", AskRights())
	return r.Run()
}
