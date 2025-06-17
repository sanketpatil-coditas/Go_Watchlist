package models

type RemoveUserRequest struct {
	Name   string `json:"name" validate:"required"`
	PAN    string `json:"pan" validate:"required,alphanum,len=10"`
	Mobile string `json:"mobile" validate:"required,len=10,numeric"`
	Age    int    `json:"age" validate:"required,gt=0"`
	DOB    string `json:"dob" validate:"required"`
}

type RemoveUserResponse struct {
	Name   string `json:"name"`
	PAN    string `json:"pan"`
	Mobile string `json:"mobile"`
	Age    int    `json:"age"`
	DOB    string `json:"dob"`
}
