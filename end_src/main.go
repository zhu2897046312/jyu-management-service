package main

import (
	"fmt"
	"jyu-service/routers"
)
func main() {
    fmt.Println("Hello, World!")
	r := routers.Router()
	r.Run(":8081")
}