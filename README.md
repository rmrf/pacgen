# Pacgen

- Proxy [Pac](https://en.wikipedia.org/wiki/Proxy_auto-config) file Generator, serve pac file as web server locally.
- Config Browser proxy with PAC url: `http://loalhost:8001`
- Different domains using different proxies. For example:
  - use Outer Special proxy to access Gmail.
  - use Corp proxy to access corperations' internal sites.

![image](https://github.com/rmrf/pacgen/assets/42414/c15f7290-5937-4811-8ea7-c31eb0e00f6c)


## Why

- I love to use proxy for accessing sites, it's always clean and easy to troubleshooting.
- With fine-tuned domains inside [Pac](https://en.wikipedia.org/wiki/Proxy_auto-config) file, I can reach any site with best speed.
- That's why I start this `pacgen` project, to mapping `[domains] => [proxy]` easily.
- Those domains which not matching any rules will goest directly access.

## How to Run it

- Copy proxy domains folder

```bash
  cp -r proxy-domains-example proxy-domains
```

- Modify outer/internal domains txt file inside `proxy-domains`
- Modify `config.toml` in root of project:
  - define different proxies for each domain txt file inside `proxy-domains`
- Run it

```bash
go build && ./pacgen -config ./config.toml
```

- Browser proxy settings with PAC url: `http://loalhost:8001/`

## TODO

- [] Auto reload when proxy-domains file changed.
