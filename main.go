package main

import "/Users/takeuchishougo/sampleapp/golang-OAuth/config"

func main() {
	config := config.NewConfig()
	server := server.NewServer()

	server.Run(config.Port)
}
