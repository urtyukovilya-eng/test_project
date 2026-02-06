package configs

import "os"

const (
	portTag           = "PORT"
	loggerEnvironment = "ENVIRONMENT"
	sqlHost           = "SQL_HOST"
	sqlPOrt           = "SQL_PORT"
	sqlDbName         = "SQL_DB_NAME"
	sqlUser           = "SQL_USER"
	sqlPass           = "SQL_PASSWORD"
	sqlSsl            = "SQL_SSL_MODE"
	sqlTimezone       = "SQL_TIMEZONE"
	sqlTtl            = "SQL_TTL_STATEMENT"
)

type Config struct {
	Port         string
	DataBase     *SqlConfig
	LoggerConfig *LoggerConfig
}

type SqlConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	SSLMode  string
	Timezone string
	TTL      string
}

type LoggerConfig struct {
	Environment string
}

func NewConfig() *Config {
	var port string
	var sqlConfig SqlConfig
	var loggerConfig LoggerConfig
	port = os.Getenv(portTag)
	loggerConfig.Environment = os.Getenv(loggerEnvironment)
	sqlConfig.Host = os.Getenv(sqlHost)
	sqlConfig.Port = os.Getenv(sqlPOrt)
	sqlConfig.DbName = os.Getenv(sqlDbName)
	sqlConfig.User = os.Getenv(sqlUser)
	sqlConfig.Password = os.Getenv(sqlPass)
	sqlConfig.SSLMode = os.Getenv(sqlSsl)
	sqlConfig.Timezone = os.Getenv(sqlTimezone)
	sqlConfig.TTL = os.Getenv(sqlTtl)
	glCfg := &Config{
		Port:         port,
		DataBase:     &sqlConfig,
		LoggerConfig: &loggerConfig,
	}

	return glCfg
}
