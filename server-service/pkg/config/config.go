package config_server_service

import "github.com/spf13/viper"

type DataBasePostgres struct {
	Port               string `mapstructure:"PORT_SERVER_SERVICE"`
	DBConeectionString string `mapstructure:"POSTGRES_DB"`
	DBName             string `mapstructure:"DATABASENAME"`
	User               string `mapstructure:"USER"`
	UserPassword       string `mapstructure:"PASSWORD"`
	Host               string `mapstructure:"HOST"`
}

type S3Bucket struct {
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Region          string `mapstructure:"Region"`
	BucketName      string `mapstructure:"BucketName"`
}

type MongodDb struct {
	MongodbURL           string `mapstructure:"MONGODBURL"`
	DataBase             string `mapstructure:"MONOGODBDATABASE"`
	ServerChatCollection string `mapstructure:"SERVERCHATMONGODBCOLLECTION"`
	ServerForunPost      string `mapstructure:"SERVERFORUMMONGODBCOLLECTION"`
	ServerForumCommand   string `mapstructure:"SERVERFORUMCOMMENTMONGODBCOLLECTION"`
}

type Auth_service struct {
	Auth_Service_port string `mapstructure:"AUTH_SERVICE_PORT"`
}

type ServerCredential struct {
	ServerPort string `mapstructure:"PORT_SERVER_SERVICE"`
}

type Kafka struct {
	KafkaPort  string `mapstructure:"KAFKAPORT"`
	KafkaTopic string `mapstructure:"TOPIC"`
}

type Config struct {
	DB               DataBasePostgres
	ServerCredential ServerCredential
	S3               S3Bucket
	KafkaConsumer    Kafka
	MongoDB          MongodDb
	Auth             Auth_service
}

func ConfigInit() (*Config, error) {
	var config = Config{}
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.DB)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.ServerCredential)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.S3)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.KafkaConsumer)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.MongoDB)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config.Auth)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
