package settings

type Settings struct {
	JwtSecret []byte
	Port      string
}

var DefaultSettings Settings

func InitSettings() {
	DefaultSettings = Settings{
		JwtSecret: []byte("THIS_IS_VERY_BIG_SECRET"),
		Port:      ":3000",
	}
}
