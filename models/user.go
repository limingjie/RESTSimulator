package models

type (
	// User represents the structure of our resource
	User struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
		Age    int    `json:"age"`
	}
)
