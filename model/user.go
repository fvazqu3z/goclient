package model

import "encoding/json"

// User Structure to store User
type User struct {
	UserID int    `json:"userID"`
	Name   string `json:"name"`
	Email 	string `json:"email"`
	Phone 	string `json:"phone"`
}

// IsValid Validate User data
func (u *User) IsValid() bool {
	if u.UserID == 0 {
		return false
	}
	if u.Email == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.Phone == "" {
		return false
	}
	return true
}

// ToJson Convert User to JSon
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
