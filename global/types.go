package global

type AppConfigure struct {
	Jwt      *JwtConfig      `yaml:"Jwt"`
	Database *DatabaseConfig `yaml:"Database"`
}

type JwtConfig struct {
	JwtSecret           string `yaml:"jwtSecret"`
	TokenExpireDuration int    `yaml:"tokenExpireDuration"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
