package gener

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gener struct {
	SQLiteDB *sql.DB
}

func NewGener(db *sql.DB) *Gener {
	return &Gener{SQLiteDB: db}
}

func (g *Gener) GetPac(gctx *gin.Context) {
	var proxyMap = make(map[string]string)
	proxyMap["proxyGFW"] = "192.168.100.14:3128"
	proxyMap["proxyInternal"] = "192.168.100.12:3128"
	gctx.HTML(http.StatusOK, "template.tmpl", proxyMap)

}

func (g *Gener) AddTargetDomain(domain, proxy string) {

}
