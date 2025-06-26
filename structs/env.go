package structs

type EnviromentModel struct {
	DevMode  devMode
	Database database
	Server   server
	Cors     corsModel
	Email    email
	Other    other
	Files    files
	SSO      sso
	Auth     auth
	Env      env
}

type devMode struct {
	TestDocker     bool   `mapstructure:"test_docker"`
	TestRadiusRole string `mapstructure:"test_radius_role"`
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
	User       string
	Password   string
	Host       string
	Port       string
	Sender     string
	Auth       bool
	SMTP       []string
	AuthType   string
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

type auth struct {
	Type     string       `mapstructure:"type"`
	Keycloak authKeycloak `mapstructure:"keycloak"`
	Radius   authRadius   `mapstructure:"radius"`
}

type authKeycloak struct {
	Url    string `mapstructure:"url"`
	Realm  string `mapstructure:"realm"`
	User   string `mapstructure:"user"`
	Secret string `mapstructure:"secret"`
}

type authRadius struct {
	Server         string   `mapstructure:"server"`
	Secret         string   `mapstructure:"secret"`
	NasPort        string   `mapstructure:"nas_port"`
	TimeoutSeconds int      `mapstructure:"timeout_seconds"`
	TokenLifespan  string   `mapstructure:"token_lifespan"`
	AdminVcRole    []string `mapstructure:"admin_vc_role"`
	GroupKey       string   `mapstructure:"group_key"`
	AllowedGroupValues []string `mapstructure:"allowed_group_values"`
}
