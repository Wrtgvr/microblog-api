package config

type Config struct {
	Server *ServerConfig
	DB     *DatabaseConfig
}

func LoadConfig() (*Config, error) {
	dbCfg, err := loadDbConfig()
	if err != nil {
		return nil, err
	}
	serverCfg, err := loadServerConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: serverCfg,
		DB:     dbCfg,
	}, nil
}
