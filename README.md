# Toqcer.id
## golang-api-boilerplate
>golang-api-boilerplate is api for store management system web app by ToqCer.id, this api built with golang programming language to aim best performance and speed of development.
## How To Run

### copy .env.example to .env and edit your .env config
    cp .env.example .env
### install dependencies
    go mod tidy
### run database migration (if needed) more details at [commands](cmd/USAGE.md)
    go run cmd/migration.go up
### run server
    go run main.go
### go to api docs (open in browser)
    http://localhost:6991/swagger