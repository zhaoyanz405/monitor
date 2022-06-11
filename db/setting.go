package db

type Config struct {
	pgUser string
	pgPass string
	pgHost string
}

var CurConfig *Config

func init() {
	CurConfig = &Config{
		pgUser: "",
		pgPass: "",
		pgHost: "",
	}
}
