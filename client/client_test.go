package client

import (
	"os"
	"testing"

	"apiclient/usecase"

	"github.com/go-resty/resty/v2"
)

func TestClient(t *testing.T) {
	client := resty.New()
	usecase.GetAllUsers(client)

	localData, localError := os.ReadFile("data/localusers.csv")
	remoteData, remoteError := os.ReadFile("data/users.csv")

	if localError != nil{
		t.Errorf("Error reading local users")
	}

	if remoteError != nil{
		t.Errorf("Error reading remote users")
	}

	if string(localData) != string(remoteData){
		t.Errorf("The local data do not match with server data")
	}


}