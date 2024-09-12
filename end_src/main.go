package main

import (
	"fmt"
	"jyu-service/routers"
	"jyu-service/models"
)
func main() {
    fmt.Println("Hello, World!")

	go models.UpdateCourseEnrollmentsFromRedis()

	r := routers.Router()
	r.Run(":8081")
}