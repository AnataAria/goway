package config

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host         string `env:"DB_HOST,required"`
	Port         string `env:"DB_PORT" envDefault:"5432"`
	User         string `env:"DB_USER,required"`
	Password     string `env:"DB_PASSWORD,required"`
	DBName       string `env:"DB_NAME,required"`
	SSLMode      string `env:"DB_SSLMODE" envDefault:"disable"`
	MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS" envDefault:"10"`
	MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS" envDefault:"5"`
}

type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`
}
