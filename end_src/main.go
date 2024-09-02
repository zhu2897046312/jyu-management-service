package main

import (
	"fmt"
	"jyu-service/models"
	"jyu-service/utils"
	//"jyu-service/routers"
)
func main() {
    fmt.Println("Hello, World!")
	utils.InitConfig("")
	utils.InitMySQL()
	utils.InitRedis()
	// r := routers.Router()
	// r.Run()

	uer := models.UserBasicInformation{
		Account: "221110136",
		Password: "fdlkjhakf",
		EthnicGroup: "fdkakkhjkgjhjd",
	}
	uer.Update(uer)
}