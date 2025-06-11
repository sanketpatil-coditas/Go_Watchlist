// package model

// type User struct {
// 	Name   string `json:"name"`
// 	PAN    string `json:"pan"`
// 	Mobile string `json:"mobile"`
// 	Age    int    `json:"age"`
// 	DOB    string `json:"dob"`
// }

package model

type User struct {
	ID     uint   `gorm:"primaryKey" json:"-"`
	Name   string `gorm:"not null" json:"name"`
	PAN    string `gorm:"unique;not null" json:"pan"`
	Mobile string `gorm:"not null" json:"mobile"`
	Age    int    `gorm:"not null" json:"age"`
	DOB    string `gorm:"not null" json:"dob"`
}
