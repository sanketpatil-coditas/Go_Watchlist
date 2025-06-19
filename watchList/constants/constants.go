package constants

const (
	AddWatchlistRoute     = "/watchlist/add"
	RemoveWatchlistRoute  = "/watchlist/remove"
	CreateWatchlistRoute  = "/watchlist/create"
	GetUserWatchlistRoute = "/watchlist/get"

	// Validation Error Messages
	ErrNameRequired  = "Name is required"
	ErrPANInvalid    = "PAN must be exactly 10 alphanumeric characters"
	ErrMobileInvalid = "Mobile number must be exactly 10 digits"
	ErrAgeInvalid    = "Age must be greater than 0"
	ErrDOBRequired   = "DOB is required and must be in YYYY-MM-DD format"
	ErrInvalidValue  = "Invalid value"

	// Generic Error Messages
	MsgInvalidInput     = "Invalid input: "
	MsgWatchlistSuccess = "Watchlist created successfully"
	MsgWatchlistremove  = "Watchlist removed successfully"
)
