package config

type (
	// Config ...
	Config struct {
		Server     ServerConfig   `yaml:"server"`
		Database   DatabaseConfig `yaml:"database"`
		API        APIConfig      `yaml:"api"`
		Credential Credential     `yaml:"credential"`
		Firebase   FirebaseConfig `yaml:"firebase"`
		Swagger    SwaggerConfig  `yaml:"swagger"`
		Redis      Redis          `yaml:"redis"`
	}

	// ServerConfig ...
	ServerConfig struct {
		Port string `yaml:"port"`
	}

	// DatabaseConfig ...
	DatabaseConfig struct {
		Master string `yaml:"master"`
	}

	// APIConfig ...
	APIConfig struct {
		Auth string `yaml:"auth"`
	}

	SwaggerConfig struct {
		Host    string   `yaml:"host"`
		Schemes []string `yaml:"schemes"`
	}

	Credential struct {
		Id string `yaml:"id"`
		Pw string `yaml:"pw"`
		Ip string `yaml:"ip"`
	}

	// FirebaseConfig ...
	FirebaseConfig struct {
		ProjectID     string `yaml:"projectID"`
		DatabaseURL   string `yaml:"databaseURL"`
		StorageBucket string `yaml:"storageBucket"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
	}
)
