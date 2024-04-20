package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Url  string
	Port uint
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	Drive    string
}

func init() {
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.drive", "postgres")
	viper.SetDefault("api.url", "http://localhost")
	viper.SetDefault("api.port", "8080")
}

/*
Função para carregar as configurações.

Retorno:

	error: retorno caso um erro ocorra

*/
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.API = APIConfig{
		Url:  viper.GetString("api.url"),
		Port: viper.GetUint("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
		Drive:    viper.GetString("database.drive"),
	}

	return nil
}

/*
Função para pegar a configuração da API.

Retorno:

	APIConfig: struct da configuração da API

*/
func GetApiConfig() APIConfig {
	return cfg.API
}

/*
Função para pegar a configuração do banco de dados.

Retorno:

	DBConfig: struct da configuração do banco de dados

*/
func GetDbConfig() DBConfig {
	return cfg.DB
}
