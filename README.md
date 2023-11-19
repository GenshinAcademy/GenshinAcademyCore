# GenshinAcademyCore
Core for GenshinAcademy infrastructure

### Swagger Documentation
TODO
```bash
./swag init --parseInternal -d ./internal -g ../cmd/web/main.go --ot yaml 
./swag fmt
```

```bash
docker run -p 8585:8080 -e SWAGGER_JSON=/docs/swagger.yaml -v /home/jagerente/Documents/GitHub/GenshinAcademyCore/docs/:/docs swaggerapi/swagger-ui
```