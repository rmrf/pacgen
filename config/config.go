package config

type C struct {
	Listen  string           `toml:"listen"`
	Proxies map[string]Proxy `toml:"proxy"`
}

type Proxy struct {
	Address    string `toml:"address"`
	TargetFile string `toml:"target_file"`
}
