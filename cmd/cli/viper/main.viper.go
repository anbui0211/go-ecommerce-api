package main

import (
	"fmt"

	"github.com/spf13/viper"
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
	viper.AddConfigPath("./config") // path name
	viper.SetConfigName("local")    // name file
	viper.SetConfigType("yaml")

	// read configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	// read server configuration
	fmt.Println("Server Port:: ", viper.GetInt("server.port"))
	fmt.Println("Security key:: ", viper.GetString("security.jwt.key"))

	// configuration struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config: %v", err)
	}

	fmt.Println("Config Port:: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database User: %s, password: %s, host: %s", db.User, db.Password, db.Host)
		
	}
}
