package utils

import (
	"fmt"
	"report-backend-golang/global"
	"report-backend-golang/structs"
	"strings"

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
	config.Database.Migration = viper.GetBool("database.migration")

	config.Server.Mode = viper.GetString("server.mode")
	config.Server.Port = viper.GetString("server.port")

	config.Redis.Url = viper.GetString("redis.url")
	config.Redis.Password = viper.GetString("redis.password")
	config.Redis.Database = viper.GetInt("redis.database")
	config.Redis.Idle = viper.GetInt("redis.idle")
	config.Redis.Active = viper.GetInt("redis.active")
	config.Redis.Protocol = viper.GetString("redis.protocol")

	config.Cors.Allow.Headers = viper.GetStringSlice("cors.allow.headers")

	config.Email.User = viper.GetString("email.user")
	config.Email.Password = viper.GetString("email.password")
	config.Email.SMTP = viper.GetStringSlice("email.smtp")
	config.Email.Host = viper.GetString("email.host")
	config.Email.Port = viper.GetString("email.port")
	config.Email.Sender = viper.GetString("email.sender")
	config.Email.Auth = viper.GetBool("email.auth")

	config.Other.Backend = viper.GetString("other.backend")

	config.Files.FontFile = viper.GetString("files.font_file")
	config.Files.ScreenshotFile = viper.GetString("files.screenshot_file")
	config.Files.ReportFile = viper.GetString("files.report_file")
	config.Files.HtmlFile = viper.GetString("files.html_file")
	config.Files.LogFile = viper.GetString("files.log_file")


	// config.Email.User = viper.GetString("email.user")
	// config.Email.Password = viper.GetString("email.password")
	// config.Email.Port = viper.GetString("email.port")
	// config.Email.Host = viper.GetString("email.host")

	global.EnvConfig = &config
}
