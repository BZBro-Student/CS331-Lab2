package user

type User struct {
	Type string
}

func CurrUser(Type string) *User {
	return &User(Type)
}

