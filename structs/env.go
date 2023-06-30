package structs

type EnviromentModel struct {
	Database database
	Server   server
	Redis    redisModel
	Cors     corsModel
	Email    email
	Other    other
	Files    files
	SSO      sso
}

type files struct {
	FontFile       string
	ScreenshotFile string
	ReportFile     string
	HtmlFile       string
	LogFile        string
	TemplateFile   string
	MigrationsFile string
	ChromePath     string
}

type other struct {
	Backend string
}

type email struct {
	User     string
	Password string
	Host     string
	Port     string
	Sender   string
	Auth     bool
	SMTP     []string
}

type sso struct {
	SsoUrl         string
	SsoRealm       string
	SsoUser        string
	SsoPassword    string
	SsoLicense_key string
}

type server struct {
	Port string
	Mode string
}

type corsModel struct {
	Allow corsAllowModel
}

type corsAllowModel struct {
	Headers []string
}

type database struct {
	Client      string
	MaxIdle     uint
	MaxLifeTime string
	MaxOpenConn uint
	User        string
	Password    string
	Host        string
	Db          string
	Params      string
	Port        string
	LogEnable   int
	Migration   bool
}

type redisModel struct {
	Url      string
	Password string
	Database int
	Idle     int
	Active   int
	Protocol string
}
