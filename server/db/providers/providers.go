package providers

type Provider interface {
	// User provider interface
	AddUser()
	UpdateUser()
	DeleteUser()
	ListUser()
	GetUserByEmail()
	GetUserByID()
	UpdateUsers()
}
