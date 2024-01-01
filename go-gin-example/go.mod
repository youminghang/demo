module github.com/youminghang/go-gin-example

go 1.21.5

replace (
	github.com/youminghang/go-gin-example/conf => /workspaces/demo/go-gin-example/pkg/conf
	github.com/youminghang/go-gin-example/middleware => /workspaces/demo/go-gin-example/middleware
	github.com/youminghang/go-gin-example/models => /workspaces/demo/go-gin-example/models
	github.com/youminghang/go-gin-example/pkg/setting => /workspaces/demo/go-gin-example/pkg/setting
	github.com/youminghang/go-gin-example/routers => /workspaces/demo/go-gin-example/routers
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/gorm v1.25.5 // indirect
)
