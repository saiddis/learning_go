package main

type User struct {
	username string
	email    string
	age      int
}

type UserManager struct {
	Users []User
}

func (um *UserManager) AddUser(username, email string, age int) {
	user := User{
		username: username,
		email:    email,
		age:      age,
	}

	um.Users = append(um.Users, user)
}

func (um UserManager) UsersOlderThan(age int) []User {
	var olderUsers []User

	for _, v := range um.Users {
		if v.age > age {
			olderUsers = append(olderUsers, v)
		}
	}

	return olderUsers
}
