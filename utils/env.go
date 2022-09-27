package utils

import (
	"fmt"
	"strings"
	"report-backend-golang/global"
	"report-backend-golang/structs"

	"github.com/spf13/viper"
)

func LoadEnvironment() {
	loadConfigFile()
	viperConfigToModel()
}

func loadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("沒有發現 config.yml，改抓取環境變數")
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		} else {
			// 有找到 config.yml 但是發生了其他未知的錯誤
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func viperConfigToModel() {
	var config structs.EnviromentModel
	config.Database.Client = viper.GetString("database.client")
	config.Database.Host = viper.GetString("database.host")
	config.Database.User = viper.GetString("database.user")
	config.Database.Password = viper.GetString("database.password")
	config.Database.Db = viper.GetString("database.name")
	config.Database.MaxIdle = uint(viper.GetInt("database.max_idle"))
	config.Database.MaxOpenConn = uint(viper.GetInt("database.max_open_conn"))
	config.Database.MaxLifeTime = viper.GetString("database.max_life_time")
	config.Database.Params = viper.GetString("database.params")
	config.Database.Port = viper.GetString("database.port")
	config.Database.LogEnable = viper.GetInt("database.log_enable")

	config.Server.Mode = viper.GetString("server.mode")
	config.Server.Port = viper.GetString("server.port")

	config.Cors.Allow.Headers = viper.GetStringSlice("cors.allow.headers")

	config.Other.Backend = viper.GetString("other.backend")
	config.Other.Migration = viper.GetBool("other.migration")
	config.Reportengine.HtmlPath = viper.GetString("reportengine.htmlPath")
	config.Reportengine.PicturePath = viper.GetString("reportengine.picturePath")
	config.Reportengine.PdfPath = viper.GetString("reportengine.pdfPath")
	config.Reportengine.LogPath = viper.GetString("reportengine.logpath")
	config.Email.User = viper.GetString("email.user")
	config.Email.Password = viper.GetString("email.password")
	config.Email.Port = viper.GetString("email.port")
	config.Email.Host = viper.GetString("email.host")
	config.Email.Subject = viper.GetString("email.subject")

	global.EnvConfig = &config
}
