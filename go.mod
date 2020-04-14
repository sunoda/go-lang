module Golang-Learn
go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.12
	api/apis v0.0.0
	api/database v0.0.0
	api/migrate v0.0.0
	api/models v0.0.0
	api/router v0.0.0
)

replace (
    api/apis => ./api/apis
    api/database => ./api/database
	api/migrate => ./api/migrate
    api/models => ./api/models
	api/router => ./api/router
)