package config

import (
	"fmt"
	"os"
	"trail_backend/utils/constants"

	"github.com/spf13/viper"
)

var (
	ConfigDefaultFile = "config/config.dev.yml"
	ConfigReleaseFile = "config/config.release.yml"
	ConfigDevFile     = "config/config.dev.yml"
	configType        = "yml"
)

type (
	Config struct {
		Debug          bool        `mapstructure:"debug"`
		ContextTimeout int         `mapstructure:"contextTimeout"`
		Server         Server      `mapstructure:"server"`
		Services       Services    `mapstructure:"services"`
		Database       Database    `mapstructure:"database"`
		Logger         Logger      `mapstructure:"logger"`
		Jwt            Jwt         `mapstructure:"jwt"`
		Cloudinary     Cloudinary  `mapstructure:"cloudinary"`
		GoogleOAuth    GoogleOAuth `mapstructure:"googleOAuth"`
	}

	Server struct {
		Host     string `mapstructure:"host"`
		Env      string `mapstructure:"env"`
		UseRedis bool   `mapstructure:"useRedis"`
		Port     int    `mapstructure:"port"`
	}

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
		TimeZone string `mapstructure:"timeZone"`
	}

	Jwt struct {
		Secret                string `mapstructure:"secret"`
		AccessTokenExpiresIn  int64  `mapstructure:"accessTokenExpiresIn"`
		RefreshTokenExpiresIn int64  `mapstructure:"refreshTokenExpiresIn"`
	}

	Logger struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		Prefix string `mapstructure:"prefix"`
	}

	Services struct {
	}

	Cloudinary struct {
		CloudName string `mapstructure:"cloudName"`
		ApiKey    string `mapstructure:"apiKey"`
		ApiSecret string `mapstructure:"apiSecret"`
		PublicId  string `mapstructure:"publicId"`
	}

	GoogleOAuth struct {
		RedirectURL  string   `mapstructure:"redirectURL"`
		ClientID     string   `mapstructure:"clientID"`
		ClientSecret string   `mapstructure:"clientSecret"`
		Scopes       []string `mapstructure:"scopes"`
	}

	FacebookOAuth struct {
		RedirectURL string `mapstructure:"redirectURL"`
		AppID       string `mapstructure:"appID"`
		AppSecret   string `mapstructure:"appSecret"`
		GraphAPIURL string `mapstructure:"graphAPIURL"`
	}
)

func NewConfig() *Config {
	initConfig()
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return conf
}

func initConfig() {
	var configFile string
	switch os.Getenv("ENV") {
	case constants.Prod:
		configFile = ConfigReleaseFile
		fmt.Printf("config%s\n", ConfigReleaseFile)
	case constants.Dev:
		configFile = ConfigDevFile
		fmt.Printf("config file: %s\n", ConfigDevFile)
	default:
		configFile = ConfigDefaultFile
		fmt.Printf("config%s\n", ConfigDefaultFile)
	}
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
}
