package gener

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"pacgen/config"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

type Gener struct {
	ProxyMap    map[string]Proxy
	ListenAddr  string
	PacTemplate string
}

type Proxy struct {
	Address    string
	Targert    []string
	TargertStr string
}

func NewGener(confFile string) *Gener {
	var pMap = make(map[string]Proxy)
	var conf config.C
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatalln(err)
	}
	for name, p := range conf.Proxies {
		var proxy Proxy
		proxy.Targert = getTargetDomain(p.TargetFile)
		proxy.Address = p.Address
		proxy.TargertStr = genTargetStr(proxy.Targert)
		pMap[name] = proxy
	}
	return &Gener{pMap, conf.Listen, conf.PacTemplate}

}

func getTargetDomain(fn string) []string {
	var allDomains []string
	dat, err := os.ReadFile(fn)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	for _, d := range strings.Split(string(dat), "\n") {
		if d != "" {
			allDomains = append(allDomains, d)
		}
	}
	return allDomains

}
func genTargetStr(targets []string) string {
	var osb strings.Builder
	for _, ot := range targets {
		osb.WriteByte('"')
		osb.WriteString(ot)
		osb.WriteByte('"')
		osb.WriteByte(',')
	}
	return strings.TrimSuffix(osb.String(), ",")
}

func (g *Gener) GetPac(gctx *gin.Context) {

	pacString, err := g.FormatPacTmpl(g.PacTemplate)
	if err != nil {
		log.Printf("GetPac err = %+v\n", err)
		gctx.JSON(http.StatusServiceUnavailable, gin.H{"status": "failed"})
		return
	}

	gctx.String(http.StatusOK, pacString)

}

func (g *Gener) Admin(gctx *gin.Context) {
	var proxyMap = make(map[string]string)
	gctx.HTML(http.StatusOK, "index.tmpl", proxyMap)

}

func (g *Gener) AddTargetDomain(domain, proxy string) {

}

func (g *Gener) FormatPacTmpl(pacFile string) (string, error) {
	pacDat, err := os.ReadFile(pacFile)
	if err != nil {
		log.Printf("FormatPacTmpl err = %+v\n", err)
		return "", err
	}

	var result bytes.Buffer

	type data struct {
		OuterProxy      string
		OuterTargets    string
		InternalProxy   string
		InternalTargets string
	}

	var d = &data{OuterProxy: g.ProxyMap["outer"].Address,
		OuterTargets:    g.ProxyMap["outer"].TargertStr,
		InternalProxy:   g.ProxyMap["internal"].Address,
		InternalTargets: g.ProxyMap["internal"].TargertStr}

	tmpl, err := template.New("pacTmpl").Parse(string(pacDat))
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&result, d)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
