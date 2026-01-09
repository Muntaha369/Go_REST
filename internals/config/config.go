package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

//!This mentions where to start reading code
//?Tips

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"` //This is same as ENV but no second field
}

//This is the Config struct defines what is necessary
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"` //!The 1st K:V pairs defines map ENV to env field in yaml file 2nd K:V pairs defines env Can be overiden by ENV variable explicitly mentioned in terminal 3rd K:V pair defines that it is required without it should not work
	StoragePath string `yaml:"storage_path" env-required:"true"` //This is same as ENV but no second field
	HttpServer  `yaml:"http_server"`//This is same as ENV but no second field
}

func Mustload() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH") //This gets the ENV variable from the Terminal
	//? TIPS instead of mentioning path every time while running the project mention ENV variable lite this $env:CONFIG_PATH=./config/local.yaml

	if configPath == "" {
		//IF Config path is not mentioned before running the program Check if its mentioned in the flags
		flags := flag.String("config", "", "path to config") // first field mendtions which flag to read second sets the defaoult value if no path mentioned third random text
		//The above flags variable store the value of the flag mentioned in the Terminal while running
		flag.Parse() //This Parses the flag (Parsing differs based on the usecases it can be done to parse the struct or object to json for making api calls for this case it is diffrent) so the operation can be done

		configPath = *flags 

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	
	if err != nil {
		log.Fatal("Can't read config file:", err.Error())
	}

	return &cfg
}
