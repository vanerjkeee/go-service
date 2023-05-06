package main

import (
	"fmt"
	"github.com/vanerjkeee/go-service/api"
	"github.com/vanerjkeee/go-service/service"
)

func main() {
	config, err := service.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	api.Start(config)
}
