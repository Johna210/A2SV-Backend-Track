package bootstrap

type Env struct {
	AppEnv                string `mapstructure:"APP_ENV"`
	ServerAddress         string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout        int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPass                string `mapstructure:"DB_PASS"`
	DBName                string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret     string `mapstructure:"ACCESS_TOKEN"`
}

func NewEnv() *Env {
	env := Env{
		AppEnv:                "development",
		ServerAddress:         "localhost:4000",
		ContextTimeout:        100,
		DBHost:                "localhost",
		DBPort:                "27017",
		DBUser:                "",
		DBPass:                "",
		DBName:                "taskManager",
		AccessTokenExpiryHour: 24,
		AccessTokenSecret:     "asfgherhwegsdcvds",
	}

	return &env
}
