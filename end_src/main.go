package main

import (
	"fmt"
	"jyu-service/utils"
	"jyu-service/routers"
)
func main() {
    fmt.Println("Hello, World!")
	utils.InitConfig("")
	r := routers.Router()
	r.Run(":8081")
}