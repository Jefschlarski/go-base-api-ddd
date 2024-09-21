package valueobjects

type UserType uint

const (
	Admin UserType = 1
	User  UserType = 2
)

func (ut *UserType) Admin() {
	*ut = Admin
}

func (ut *UserType) User() {
	*ut = User
}

func (ut *UserType) IsAdmin() bool {
	return *ut == Admin
}

func (ut *UserType) IsUser() bool {
	return *ut == User
}

func (ut *UserType) UserTypesList() map[UserType]string {
	return map[UserType]string{
		Admin: "admin",
		User:  "user",
	}
}

func (ut *UserType) String() string {
	switch *ut {
	case Admin:
		return "admin"
	case User:
		return "user"
	default:
		return "unknown"
	}
}
