package structs

type EnviromentModel struct {
	Database     database
	Server       server
	Cors         corsModel
	Other        other
	Reportengine reportengine
	Email        email
	Migration    migration
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

type reportengine struct {
	PicturePath string
	HtmlPath    string
	PdfPath     string
	LogPath     string
}

type email struct {
	User     string
	Password string
	Host     string
	Port     string
	Subject  string
}

type migration struct {
	Controller string
}
