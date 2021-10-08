package gener

import (
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/rmrf/pacgen/config"
)

type Gener struct {
	ProxyMap   map[string][]string
	ListenAddr string
}

func NewGener(confFile string) *Gener {
	var pMap = make(map[string][]string)
	var conf config.C
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatalln(err)
	}
	for name, _ := range conf.Proxies {
		pMap[name] = []string{}
	}
	return &Gener{pMap, conf.Listen}
}

func (g *Gener) GetPac(gctx *gin.Context) {
	var proxyMap = make(map[string]string)
	proxyMap["proxyGFW"] = "192.168.100.14:3128"
	proxyMap["proxyInternal"] = "192.168.100.12:3128"
	gctx.HTML(http.StatusOK, "pac.tmpl", proxyMap)

}

func (g *Gener) Admin(gctx *gin.Context) {
	var proxyMap = make(map[string]string)
	proxyMap["proxyGFW"] = "192.168.100.14:3128"
	proxyMap["proxyInternal"] = "192.168.100.12:3128"
	gctx.HTML(http.StatusOK, "index.tmpl", proxyMap)

}

func (g *Gener) AddTargetDomain(domain, proxy string) {

}
