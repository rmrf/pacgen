package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rmrf/pacgen/gener"
)

func main() {
	var dbFile string
	var listenAddr string
	flag.StringVar(&dbFile, "db-file", "sql.db", "SQLite db file")
	flag.StringVar(&listenAddr, "listen", "0.0.0.0:8001", "which host:port will we listen on")
	flag.Parse()

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	g := gener.NewGener(db)

	router := gin.Default()
	router.LoadHTMLFiles("template/template.tmpl")
	router.GET("/proxy.pac", g.GetPac)
	//router.GET("/admin", g.Admin)

	router.Run(listenAddr)
}
