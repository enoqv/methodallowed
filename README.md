# HTTP Request Method Allowed

config example

```yml
# Static configuration

experimental:
  plugins:
    method_allowed:
        moduleName: github.com/enoqv/method_allowed
        version: v0.1.4

```

```yml
# Dynamic configuration

http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - my-plugin

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000
  
  middlewares:
    my-plugin:
      plugin:
        method_allowed:
          Message: "Method Allowed"
          Methods:
            - GET
            - POST
```