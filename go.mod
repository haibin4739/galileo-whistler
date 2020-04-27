module guthub.com/haibin4739/galileo-whistler

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.55.0
	gopkg.in/yaml.v2 v2.2.8
)

replace (
	guthub.com/haibin4739/galileo-whistler/conf => C:/Code/Go/galileo-whistler/conf
	guthub.com/haibin4739/galileo-whistler/conf/setting => C:/Code/Go/galileo-whistler/conf/setting

	guthub.com/haibin4739/galileo-whistler/middleware => C:/Code/Go/galileo-whistler/middleware
	guthub.com/haibin4739/galileo-whistler/model => C:/Code/Go/galileo-whistler/model

	guthub.com/haibin4739/galileo-whistler/db => C:/Code/Go/galileo-whistler/db

	guthub.com/haibin4739/galileo-whistler/pkg => C:/Code/Go/galileo-whistler/pkg
	guthub.com/haibin4739/galileo-whistler/pkg/e => C:/Code/Go/galileo-whistler/pkg
	guthub.com/haibin4739/galileo-whistler/pkg/util => C:/Code/Go/galileo-whistler/pkg/util

	guthub.com/haibin4739/galileo-whistler/routers => C:/Code/Go/galileo-whistler/routers
	guthub.com/haibin4739/galileo-whistler/routers/api => C:/Code/Go/galileo-whistler/routers/api
)
