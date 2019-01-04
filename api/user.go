package api

import "github.com/jamesfzhang/ynab/model"

// Get current user info.
// https://api.youneedabudget.com/v1#/User/getUser
func (service *UserService) Get() (user model.User, err error) {

	var result model.UserResponse
	err = service.Client.get("/user", &result)
	if err != nil {
		return
	}

	user = result.Data.User
	return
}
