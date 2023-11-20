# GenshinAcademyCore

## Local build

```bash
make rebuild
```

## Deploy

[Read about deploying here.](deploy/README.md)

## Swagger Documentation

Get swaggo pre-compiled binary from the official [release page](https://github.com/swaggo/swag/releases).

### Generate documentation

```bash
./swag init --parseInternal -d ./internal -g ../cmd/web/main.go --ot yaml 
```

### Format annotations

```bash
./swag fmt
```

### Preview swagger documentation

```bash
docker run -p 8585:8080 -e SWAGGER_JSON=/docs/swagger.yaml -v /docs/:/docs swaggerapi/swagger-ui
```