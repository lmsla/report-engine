package structs

type EnviromentModel struct {
	Database database
	Server   server
	Cors     corsModel
	Other    other
}

type other struct {
	Backend   string
	Migration bool
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
}
