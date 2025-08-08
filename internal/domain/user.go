package domain

const (
	RoleTechnician = "technician"
	RoleManager    = "manager"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	Roles    []string
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
}
