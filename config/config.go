package config

type C struct {
	Listen                 string           `toml:"listen"`
	PacTemplate            string           `toml:"pac_template"`
	ProxyAutoReloadSeconds int              `toml:"proxy_auto_reload_seconds"`
	Proxies                map[string]Proxy `toml:"proxy"`
}

type Proxy struct {
	Address    string `toml:"address"`
	TargetFile string `toml:"target_file"`
}
