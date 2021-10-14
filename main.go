package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/rmrf/pacgen/gener"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "config", "config.toml", "configuration file")
	flag.Parse()

	g := gener.NewGener(confFile)

	router := gin.Default()
	router.GET("/proxy.pac", g.GetPac)
	router.GET("/admin", g.Admin)

	router.Run(g.ListenAddr)
}
