package structs

type EnviromentModel struct {
	Database database
	Server   server
	Cors     corsModel
	Email    email
	Other    other
	Files    files
	SSO      sso
	Env      env
}

type env struct {
	WaitSecond int
}

type files struct {
	FontFile       string
	ScreenshotFile string
	ReportFile     string
	HtmlFile       string
	LogPath        string
	TemplateFile   string
	MigrationsFile string
	ChromePath     string
	LogoFile       string
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
	AuthType string
	DisableTLS bool
}

type sso struct {
	Url        string
	Realm      string
	User       string
	Password   string
	LicenseKey string
	ClientID   string
	AdminRole  string
	UserRole   string
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
