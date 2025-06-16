package models

type RemoveUserRequest struct {
	Name   string `json:"name" binding:"required"`
	PAN    string `json:"pan" binding:"required,alphanum,len=10"`
	Mobile string `json:"mobile" binding:"required,len=10,numeric"`
	Age    int    `json:"age" binding:"required,gt=0"`
	DOB    string `json:"dob" binding:"required"`
}

type RemoveUserResponse struct {
	Name   string `json:"name"`
	PAN    string `json:"pan"`
	Mobile string `json:"mobile"`
	Age    int    `json:"age"`
	DOB    string `json:"dob"`
}
