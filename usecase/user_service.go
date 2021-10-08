package usecase

import (
	"encoding/json"
	"fmt"
	"log"

	"apiclient/model"
	"apiclient/repository"

	"github.com/go-resty/resty/v2"
)

// GetAllUsers Get all user from server API
func GetAllUsers(client *resty.Client) {
	resp, err := client.R().Get("http://localhost:8080/users")

	if err != nil {
		log.Println(err)
		fmt.Println("Error getting users")
		return
	}

	// Unmarshal JSON data
	var jsonData []model.User
	textBytes := []byte(resp.Body())
	err = json.Unmarshal(textBytes, &jsonData)

	if err != nil {
		log.Println(err)
		fmt.Println("Error unmarchalling users")
		return
	}

	err = repository.SaveUsersToCSV("data/localusers.csv" , jsonData )

	if err != nil {
		log.Println(err)
		fmt.Println("Error saving users")
		return
	}

	fmt.Println("Data was sucesfully saved")


}
