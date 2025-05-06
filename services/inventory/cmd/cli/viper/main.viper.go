package main

import (
	"fmt"
	viper "github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Println("Server Port: ", viper.GetInt("server.port"))

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	fmt.Println("Config Port: ", config.Server.Port)
	fmt.Println("Database User: ", config.Databases[0].User)

	for _, database := range config.Databases {
		fmt.Println("Database User: ", database.User)
	}
}
