package initialize

import (
	"fmt"

	"ecommerce/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config: %v", err)
	}
}
