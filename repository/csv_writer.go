package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"apiclient/model"
)

// SaveUsersToCSV Save users to csv file
func SaveUsersToCSV(fileName string,users []model.User) error {
	// open file
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return errors.New("Error opening file " + fileName)
	}
	fmt.Println("File opened")

	//close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	var separator string
	var endLine string
	var header string
	separator = ","
	endLine = "\n"

	header = "UserID,Name,Email,Cellphone"
	header += endLine
	_ , err = f.WriteString(header)
	if err != nil {
		log.Println(err)
		return errors.New("Error writing headers to file " + fileName)
	}

	for _, user := range users {
		var row string
		row = strconv.Itoa(user.UserID)
		row += separator
		row += user.Name
		row += separator
		row += user.Email
		row += separator
		row += user.Phone
		row += endLine
		_ , err = f.WriteString(row)

		if err != nil {
			log.Println(err)
			return errors.New("Error writing " + user.Name + " to file " + fileName)
		}
	}

	err = f.Sync()
	if err != nil {
		log.Println(err)
		return errors.New("Error synchronizing the file " + fileName)
	}

	return nil


}

