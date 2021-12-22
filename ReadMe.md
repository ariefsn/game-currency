# Game Currency

> Game Currency.

## Dependencies

- [Go Chi](https://github.com/go-chi/chi)
- [Unrolled Secure](https://github.com/unrolled/secure)
- [Unrolled Render](https://github.com/unrolled/render)
- [Gorm](https://gorm.io/)
- [Go Resty](https://github.com/go-resty/resty)
- [HTTP Swagger](https://github.com/swaggo/http-swagger)

## Requirements

- Go 1.16
- Docker

## Tested on

- Ubuntu 20.04
- MacOS 12.1

## How To

1. Run `go mod tidy`
2. Copy `.env.example` to `.env`
3. Setup configuration `.env`
4. Run `sudo chmod +x *.sh`
5. Execute `./up-build.sh` or run `docker-compose up --build -d`
6. Go to `{{url}}/docs/index.html` for the Swagger Documentation

## Testing

- Run `go test ./services/currency-service/controllers ./services/currency-service/repositories -v`

## Tools

- `./docker-rmi-none.sh`
  : will clear docker images that having `none` tags or name
- `./logs.sh <service-name>`
  : see log for specific service, `service-name` is required
- `./swag-init.sh <service-name>`
  : initiate docs for swagger file, fill the `service-name` if want to specific. By default will generate all if not define. this will require `swag` binary, install it with `go get github.com/swaggo/swag/cmd/swag`. For more info please visit [HTTP Swagger](https://github.com/swaggo/http-swagger)
- `./up-build.sh`
  : rebuild and start all services with `docker-compose`
- `./stop.sh`
  : stop all services
