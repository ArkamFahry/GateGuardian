package providers

type Provider interface {
	// User provider interfaces
	AddUser()
	UpdateUser()
	DeleteUser()
	ListUser()
	GetUserByEmail()
	GetUserByID()
	UpdateUsers()
}
