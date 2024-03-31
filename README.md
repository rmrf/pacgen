# pacgen

- Proxy Pac file Generator, serve pac file as web server locally.
- Config Browser proxy with PAC url: `http://loalhost:8001`
- Support different domains using different proxies. For example:
  - Outer proxy to access Gmail.
  - Internal proxy to access  corperations' sites.

## How to Run it

- Copy proxy domains folder

```bash
  cp -r proxy-domains-example proxy-domains
```

- Modify outer/internal domains txt file inside `proxy-domains`
- Modify `config.toml` in root of project:
  - define diffent proxy for each domain txt file inside `proxy-domains`
- Run it

```bash
go build && ./pacgen -config ./config.toml
```

- Browser proxy settings with PAC url: `http://loalhost:8001/`
