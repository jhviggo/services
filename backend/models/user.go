package models

type User struct {
	DefaultModel
	Name     string    `json:"name"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password,omitempty"`
	Role     string    `json:"role,omitempty"`
	Token    string    `json:"token,omitempty" gorm:"-"`
	Vehicles []Vehicle `json:"-" gorm:"foreignKey:UserId;references:ID"`
}
