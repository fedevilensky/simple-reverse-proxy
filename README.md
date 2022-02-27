Table of Contents:
- [Simple Reverse Proxy](#simple-reverse-proxy)
  - [Configuration](#configuration)


# Simple Reverse Proxy
This reverse proxy works by proxying a subdomain to whatever address you want

To run it, just call the compiled program, it will look for a file called `config.yaml` in `$pwd`. If you want to load another config file, you need to pass it as an argument `simpleproxy /path/to/config.yml`. The extension does not really matter, as long as the file itself is written in yaml

## Configuration
As the name suggests, it's pretty simple, you just need to specify which port the program will run, and the routing. For example
```yaml
---
port: 80
routes:
  - subdomain: away
    to: http://myurl/
  - subdomain: home
    to: http://127.0.0.1/
  - subdomain: farfaraway
    to: http://anotherurl:port/from/here/
```
