package main

import (
	"testing"

	"apiclient/client"
)

func TestMain(m *testing.M) {
	client.StartClient()
}
