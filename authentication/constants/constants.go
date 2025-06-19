package constants

const (
	//Routers
	AddUserRoute        = "/authentication/createUser"
	RemoveUserRoute     = "/authentication/removeUser"

	// Validation Error Messages
	ErrNameRequired         = "Name is required"
	ErrPANInvalid           = "PAN must be exactly 10 alphanumeric characters"
	ErrMobileInvalid        = "Mobile number must be exactly 10 digits"
	ErrAgeInvalid           = "Age must be greater than 0"
	ErrDOBRequired          = "DOB is required and must be in YYYY-MM-DD format"
	ErrInvalidValue         = "Invalid value"

	// Generic Error Messages
	MsgInvalidInput         = "Invalid input: "
	MsgUserRegisterSuccess  = "User registered successfully"
	MsgUserRegisterremove  = "User removed successfully"
)