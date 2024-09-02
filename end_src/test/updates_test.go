package test

import (
	"jyu-service/models"
	"testing"
)

func TestMain(t *testing.T) {
	user := models.UserBasicInformation{
		Account: "221110136",
		Password: "fdlkjaldf",
		Name: "zhuyi",
	}
	user.Update(user)
}
