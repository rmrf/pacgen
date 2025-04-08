package gener

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"pacgen/config"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

type Gener struct {
	ProxyMap    map[string]Proxy
	ListenAddr  string
	PacTemplate string
	C           config.C
}

type Proxy struct {
	Address    string
	Targert    []string
	TargertStr string
}

func NewGener(confFile string) *Gener {
	var conf config.C
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatalln(err)
	}
	proxyMap := generateProxyMap(conf)
	return &Gener{proxyMap, conf.Listen, conf.PacTemplate, conf}
}

func generateProxyMap(conf config.C) map[string]Proxy {
	proxyMap := make(map[string]Proxy)
	for name, p := range conf.Proxies {
		var proxy Proxy
		proxy.Targert = getTargetDomain(p.TargetFile)
		proxy.Address = p.Address
		proxy.TargertStr = genTargetStr(proxy.Targert)
		proxyMap[name] = proxy
	}
	return proxyMap
}

func (g *Gener) WatchProxyMap(quit chan struct{}) {
	ticker := time.NewTicker(time.Duration(g.C.ProxyAutoReloadSeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				proxyMap := generateProxyMap(g.C)
				g.ProxyMap = proxyMap
			case <-quit:
				ticker.Stop()
				log.Println("WatchProxyMap Ticker stopped")
				return
			}
		}
	}()
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

	// Set Expires header for 600 seconds
	expires := time.Now().UTC().Add(time.Second * time.Duration(g.C.ExpireSeconds)).Format(http.TimeFormat)
	gctx.Header("Expires", expires)

	gctx.String(http.StatusOK, pacString)
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

	d := &data{
		OuterProxy:      g.ProxyMap["outer"].Address,
		OuterTargets:    g.ProxyMap["outer"].TargertStr,
		InternalProxy:   g.ProxyMap["internal"].Address,
		InternalTargets: g.ProxyMap["internal"].TargertStr,
	}

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
