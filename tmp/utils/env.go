package utils

import (
	"fmt"
	"report-backend-golang/global"
	"report-backend-golang/structs"
	"strings"
	// "report-backend-golang/log"
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
			fmt.Println("沒有發現 config.yml,改抓取環境變數")
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

	config.Env.WaitSecond = viper.GetInt("env.wait_second")

	config.Cors.Allow.Headers = viper.GetStringSlice("cors.allow.headers")

	config.Email.User = viper.GetString("email.user")
	config.Email.Password = viper.GetString("email.password")
	config.Email.SMTP = viper.GetStringSlice("email.smtp")
	config.Email.Host = viper.GetString("email.host")
	config.Email.Port = viper.GetString("email.port")
	config.Email.Sender = viper.GetString("email.sender")
	config.Email.Auth = viper.GetBool("email.auth")
	config.Email.AuthType = viper.GetString("email.auth_type")
	config.Email.DisableTLS = viper.GetBool("email.disable_tls")

	config.Other.Backend = viper.GetString("other.backend")

	config.Files.FontFile = viper.GetString("files.font_file")
	config.Files.ScreenshotFile = viper.GetString("files.screenshot_file")
	config.Files.ReportFile = viper.GetString("files.report_file")
	config.Files.HtmlFile = viper.GetString("files.html_file")
	config.Files.LogPath = viper.GetString("files.log_path")
	config.Files.TemplateFile = viper.GetString("files.template_file")
	config.Files.MigrationsFile = viper.GetString("files.migrations_file")
	config.Files.ChromePath = viper.GetString("files.chrome_path")
	config.Files.LogoFile = viper.GetString("files.logo_file")

	config.SSO.Url = viper.GetString("sso.url")
	config.SSO.Realm = viper.GetString("sso.realm")
	config.SSO.User = viper.GetString("sso.user")
	config.SSO.Password = viper.GetString("sso.password")
	config.SSO.LicenseKey = viper.GetString("sso.license_key")
	config.SSO.AdminRole = viper.GetString("sso.admin_role")
	config.SSO.UserRole = viper.GetString("sso.user_role")
	config.SSO.ClientID = viper.GetString("sso.client_id")


	global.EnvConfig = &config
}
