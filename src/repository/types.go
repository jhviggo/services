package repository

/* The types reside in here because they are closely connected to the database tables */
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}
