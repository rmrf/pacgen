package config

type C struct {
	Listen                 string           `toml:"listen"`
	ExpireSeconds          int              `toml:"expire_seconds"`
	PacTemplate            string           `toml:"pac_template"`
	ProxyAutoReloadSeconds int              `toml:"proxy_auto_reload_seconds"`
	Proxies                map[string]Proxy `toml:"proxy"`
}

type Proxy struct {
	Address    string `toml:"address"`
	TargetFile string `toml:"target_file"`
}
