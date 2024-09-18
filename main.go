package main; 
type RegisteredUsers struct {
	Users []User
	Count int
}

type User struct {
	Name string
	Age int
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int {
	return u.Age
}

