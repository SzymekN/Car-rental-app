package controller

import (
	"fmt"

	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/storage"

	"github.com/gocql/gocql"
)

func SaveUserCassandra(u model.User) error {
	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Insert into userapi.users(id, firstname,lastname, age) values (?,?,?,?)`, u.Id, u.Firstname, u.Lastname, u.Age).Exec(); err != nil {
		fmt.Println(err)
	}
	return nil
}

func UpdateUserCassandra(id int, u model.User) error {

	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Update userapi.users set firstname=?, lastname=?, age=? where id=?`, u.Firstname, u.Lastname, u.Age, id).Exec(); err != nil {
		return err
	}
	return nil
}

func DeleteUserCassandra(id int) (model.User, error) {

	cas := storage.GetCassandraInstance()
	u, err := GetUserByIdCassandra(id)

	if err != nil {
		return u, err
	}

	if err := cas.Query(`Delete from userapi.users where id=?`, id).Exec(); err != nil {
		return u, err
	}

	return u, nil
}

func GetUserByIdCassandra(id int) (model.User, error) {
	cas := storage.GetCassandraInstance()
	u := model.User{}

	if err := cas.Query(`Select id, firstname, lastname, age from userapi.users where id=?`, id).Consistency(gocql.One).Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age); err != nil {
		return u, err
	}
	return u, nil
}

func GetUsersCassandra() ([]model.User, error) {
	users := []model.User{}
	u := model.User{}

	cas := storage.GetCassandraInstance()

	iter := cas.Query(`Select id, firstname, lastname, age from userapi.users`).Iter()
	for iter.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age) {
		users = append(users, u)
		u = model.User{}
	}
	if err := iter.Close(); err != nil {
		return users, err
		// log.Fatal(err)
	}
	return users, nil
}
