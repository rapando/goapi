package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rapando/goapi/pkg/books"
	"github.com/rapando/goapi/pkg/common/db"
	"github.com/rapando/goapi/pkg/common/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.InitLogger()
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	var proxies = viper.Get("PROXIES").(string)

	gin.SetMode(viper.Get("GIN_MODE").(string))

	var dbConn = db.Init(viper.Get("DB_URI").(string))

	var r = gin.Default()
	r.SetTrustedProxies(strings.Split(proxies, ","))

	books.RegisterRoutes(r, dbConn)

	r.Run(viper.Get("PORT").(string))

}
