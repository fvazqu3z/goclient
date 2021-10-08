package client

import (
	"github.com/go-resty/resty/v2"

	"apiclient/usecase"
)

// StartClient Initialize the client
func StartClient(){
	client := resty.New()
	usecase.GetAllUsers(client)
}