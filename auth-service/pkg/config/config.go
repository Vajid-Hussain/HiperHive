package configl_auth_server

import (
	"github.com/spf13/viper"
)

type DataBase struct {
	Port               string `mapstructure:"PORT_AUTH_SVC"`
	DBConeectionString string `mapstructure:"DBCONNECTION"`
	// DBConeectionString string `mapstructure:"DEPLOYDBCONNECTION"`
	DBName             string `mapstructure:"DATABASENAME"`
	User               string `mapstructure:"USER"`
	UserPassword       string `mapstructure:"USERPASSWORD"`
	Host               string `mapstructure:"HOST"`
}

type Token struct {
	AdminSecurityKey string `mapstructure:"ADMIN_TOKENKEY"`
	UserSecurityKey  string `mapstructure:"USER_TOKENKEY"`
	TemperveryKey    string `mapstructure:"TEMPERVERY_TOKENKEY"`
}

type S3Bucket struct {
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Region          string `mapstructure:"Region"`
	BucketName      string `mapstructure:"BucketName"`
}

type Mail struct {
	SecretKey string `mapstructure:"SECRETKEY"`
	From      string `mapstructure:"FROM"`
}

type Config struct {
	DB    DataBase
	S3    S3Bucket
	Token Token
	Mail  Mail
}

func InitServer() (*Config, error) {

	var config = Config{}

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.Unmarshal(&config.DB)
	viper.Unmarshal(&config.S3)
	viper.Unmarshal(&config.Token)
	viper.Unmarshal(&config.Mail)

	return &config, nil
}