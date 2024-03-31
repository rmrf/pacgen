package main

import (
	"flag"

	"pacgen/gener"

	"github.com/gin-gonic/gin"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "config", "config.toml", "configuration file")
	flag.Parse()

	g := gener.NewGener(confFile)

	router := gin.Default()
	router.GET("/", g.GetPac)

	router.Run(g.ListenAddr)
}
