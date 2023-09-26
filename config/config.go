package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"trail_backend/pkg/constants"
)

var (
	ConfigDefaultFile = "config/config.yml"
	ConfigReleaseFile = "config/config.release.yml"
	ConfigDevFile     = "config/config.yml"
	configType        = "yml"
)

var (
	configEnv     = "./config/app.env"
	configTypeEnv = "env"
	configEnvName = "app"
)

type (
	Config struct {
		Env            Env        `mapstructure:"env"`
		Cloudinary     Cloudinary `mapstructure:"cloudinary"`
		Debug          bool       `mapstructure:"debug"`
		ContextTimeout int        `mapstructure:"contextTimeout"`
		Server         Server     `mapstructure:"server"`
		Services       Services   `mapstructure:"services"`
		Database       Database   `mapstructure:"database"`
		Logger         Logger     `mapstructure:"logger"`
		Jwt            Jwt        `mapstructure:"jwt"`
		//Cloudinary     Cloudinary  `mapstructure:"cloudinary"`
		GoogleOAuth GoogleOAuth `mapstructure:"googleOAuth"`
	}

	Server struct {
		Host       string `mapstructure:"host"`
		Env        string `mapstructure:"env"`
		UseRedis   bool   `mapstructure:"useRedis"`
		Port       int    `mapstructure:"port"`
		UploadPath string `mapstructure:"uploadPath"`
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

	Env struct {
		Env           string `mapstructure:"ENV"`
		CloudinaryURL string `mapstructure:"CLOUDINARY_URL"`
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

func SetEnv(config *Config) {
	v := reflect.ValueOf(config.Env)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != "" {
			os.Setenv(v.Type().Field(i).Tag.Get("mapstructure"), v.Field(i).Interface().(string))
		}
	}
}

func initEnv(conf *Config) {
	if err := LoadConfigEnv(configEnv, configTypeEnv); err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	if err := UnmarsharConfig(&conf.Env); err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	SetEnv(conf)
}

func NewConfig() *Config {
	conf := &Config{}
	initEnv(conf)
	initConfig()
	if err := UnmarsharConfig(conf); err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return conf
}

func LoadConfigEnv(configFile, configType string) (err error) {
	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func UnmarsharConfig[E any](config *E) error {
	return viper.Unmarshal(config)

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

	if err := LoadConfigEnv(configFile, configType); err != nil {
		fmt.Println(err.Error())
	}
}
