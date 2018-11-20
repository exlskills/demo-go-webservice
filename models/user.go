package models

type User struct {
	ID       uint64 `json:"-"`
	FullName string `json:"-"`
}

func (user *User) FindById() error {
	user.ID = 1
	user.FullName = "Demo User"
	return nil
}
