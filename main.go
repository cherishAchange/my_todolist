package main

import (
	"fmt"
	"my_todolist/config"
	"my_todolist/router"
)

func main() {

	r := router.InitRouter()

	err := r.Run(fmt.Sprintf(":%d", config.Info.Server.PORT))

	if err != nil {
		return
	}

}
