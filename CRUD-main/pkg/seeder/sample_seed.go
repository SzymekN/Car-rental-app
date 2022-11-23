package seeder

import "github.com/SzymekN/CRUD/pkg/model"

var (
	Users = []*model.User{
		{Firstname: "Szymon", Lastname: "Nowak", Age: 22},
		{Firstname: "Jan", Lastname: "Kowalski", Age: 31},
		{Firstname: "Chuck", Lastname: "Norris", Age: 18},
		{Firstname: "Andrzej", Lastname: "Duda", Age: 41},
	}
	Operators = []*model.Operator{
		{Username: "admin1", Email: "admin1@admin", Password: "admin", Role: "admin"},
		{Username: "user1", Email: "user1@user", Password: "user", Role: "user"},
	}
)
