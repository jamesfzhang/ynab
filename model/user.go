package model

// https://api.youneedabudget.com/v1#/User
type UserResponse struct {
	Data UserWrapper `json:"data"`
}

type UserWrapper struct {
	User User `json:"user"`
}

type User struct {
	Id string `json:"id"`
}
