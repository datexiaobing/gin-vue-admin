package config

type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Videoport     int64
	AdminPort     int64
	WebPort       int64
	Name          string
	JwtSecret     string
	Mode          string
	DemoMsg       string
	EnableDP      bool
}

var ApplicationConfig = new(Application)
