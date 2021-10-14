package gener

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/rmrf/pacgen/config"
)

type Gener struct {
	ProMap      map[string]Pro
	ListenAddr  string
	PacTemplate string
}

type Pro struct {
	Address    string
	Targert    []string
	TargertStr string
}

func NewGener(confFile string) *Gener {
	var pMap = make(map[string]Pro)
	var conf config.C
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatalln(err)
	}
	for name, p := range conf.Proxies {
		var pro Pro
		pro.Targert = getTargetDomain(p.TargetFile)
		pro.Address = p.Address
		pro.TargertStr = genTargetStr(pro.Targert)
		pMap[name] = pro
	}
	return &Gener{pMap, conf.Listen, conf.PacTemplate}

}

func getTargetDomain(fn string) []string {
	var allDomains []string
	dat, err := os.ReadFile(fn)
	if err != nil {
		log.Printf("err = %+v\n", err)
	}
	for _, d := range strings.Split(string(dat), ",") {
		td := strings.Trim(d, "\n ")
		if td != "" {
			allDomains = append(allDomains, td)
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
		DirectTargets   string
	}

	var d = &data{OuterProxy: g.ProMap["outer"].Address,
		OuterTargets:    g.ProMap["outer"].TargertStr,
		InternalProxy:   g.ProMap["internal"].Address,
		InternalTargets: g.ProMap["internal"].TargertStr,
		DirectTargets:   g.ProMap["direct"].TargertStr}

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
