package model

//User is struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserTable users table type
type UserTable struct {
	User
	CommonColumn
}
