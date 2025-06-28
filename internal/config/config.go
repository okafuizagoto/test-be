package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	config      *Config
	cred        map[string]string
	watchConfig WatchCfg
	fileToWatch []string
)

const (
	envDevelopment   = "development"
	envStaging       = "staging"
	envProduction    = "production"
	envProductionCHC = "chc-production"
)

type (
	option struct {
		configFile      string
		credentialsFile string
	}

	WatchCfg struct {
		Path string
		Name string
	}
)

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile:      getDefaultConfigFile(),
		credentialsFile: getCredentialsFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	// configSplit := strings.Split(opt.configFile, "/")
	// lengthSplit := len(configSplit)
	// watchConfig.Path = strings.Join(configSplit[:lengthSplit-1], "/")
	// watchConfig.Name = configSplit[lengthSplit-1]

	err = yaml.Unmarshal(out, &config)
	if err != nil {
		return err
	} else {
		// ADD CONFIG FILE TO WATCH
		fileToWatch = append(fileToWatch, opt.configFile)
	}

	credentials, err := ioutil.ReadFile(opt.credentialsFile)
	if err != nil {
		return err
	}

	// fmt.Println("CREDENTIALS", credentials)

	err = json.Unmarshal(credentials, &cred)
	if err != nil {
		return err
	}

	// fmt.Println("CRED", cred)

	return yaml.Unmarshal(out, &config)
}

func PrepareWatchPath() {
	for _, path := range fileToWatch {
		viper.SetConfigFile(path)
		viper.SetConfigType("yaml")
		viper.WatchConfig()
	}
	// viper.SetConfigName(watchConfig.Name)
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(watchConfig.Path)
}

// Option ...

type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {
	configPath := "./files/etc/gold-gym-be/gold-gym-be.development.yaml"
	namespace, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	env := string(namespace)
	if os.Getenv("GOPATH") == "" {
		configPath = "files/etc/gold-gym-be/gold-gym-be.development.yaml"
	}

	if env != "" {
		switch env {
		case envStaging:
			time.Sleep(30 * time.Second)
			configPath = "/vault/secrets/database.yaml"
		case envProduction:
			time.Sleep(30 * time.Second)
			configPath = "/vault/secrets/database.yaml"
		default:
			configPath = "./gold-gym-be.development.yaml"
		}
	}

	if os.Getenv("chc") == "sementara" {
		configPath = "./gold-gym-be.chc.production.yaml"
	}

	return configPath
}

func getCredentialsFile() string {
	configPath := "./files/etc/gold-gym-be/credentials.development.json"
	namespace, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	env := string(namespace)

	if env != "" {
		switch {
		case strings.Contains(env, envDevelopment):
			configPath = "./credentials.development.json"
		case strings.Contains(env, envStaging):
			configPath = "/vault/secrets/database.json"
		case strings.Contains(env, envProduction):
			configPath = "/vault/secrets/database.json"
		default:
			if os.Getenv("GOPATH") == "" {
				configPath = "./credentials.development.json"
			}
		}
	}
	// fmt.Println("CONFIGPATH", configPath)
	return configPath
}

// Get ...
func Get() (*Config, map[string]string) {
	return config, cred
}
