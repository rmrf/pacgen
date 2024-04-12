# Pacgen

- Proxy [Pac](https://en.wikipedia.org/wiki/Proxy_auto-config) file Generator, local web server serve pac file.
- Config Browser proxy with PAC url: `http://loalhost:8001`
- Different domains using different proxies. For example:
  - with Outer Special proxy to access Gmail.
  - with Corp proxy to access corperations' internal sites.

![image](https://github.com/rmrf/pacgen/assets/42414/f83f0667-fcbd-4d5d-802f-394f871c172c)

## Why

- I love to use proxy for accessing sites, it's always clean and easy to troubleshooting.
- With fine-tuned domains inside [Pac](https://en.wikipedia.org/wiki/Proxy_auto-config) file, I can reach any site with best speed.
- That's why I start this `pacgen` project, to mapping `[domains] => [proxy]` easily.
- Those domains which not matching any rules will  access directly.

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

- [x] Auto reload proxy-domains files.
