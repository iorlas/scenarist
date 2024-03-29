package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "5432")
	viper.SetDefault("user", "postgres")
	viper.SetDefault("schema", "public")
	viper.SetDefault("datapath", "data/scenarios")

	log.Println("Starting...")
	server()
}
