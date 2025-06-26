package database

import (
	"fmt"
)

type User struct {
	ID int32
}

type UserList []User

func (u *UserList) FindUser(id int32) (User, error) {
	for i := range *u {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}

	return User{}, fmt.Errorf("User %d could not be found", id)
}

func (u *UserList) AddUser(newUser User) {
	*u = append(*u, newUser)
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (p *UserListProxy) FindUser(id int32) (User, error) {
	user, err := p.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		p.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = p.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	p.addUserToStack(user)

	fmt.Println("Returning user from database")
	p.DidLastSearchUsedCache = false
	return user, nil
}

func (p *UserListProxy) addUserToStack(user User) {
	if len(p.StackCache) >= p.StackCapacity {
		p.StackCache = append(p.StackCache[1:], user)
	} else {
		p.StackCache.AddUser(user)
	}
}
