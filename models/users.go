package models

// Project struct contains data about specific project
type User struct {
	id        int      `json:"userId"`       // user id
	UserName  string   `json:"userName"`     // username
	Email     string   `json:"userEmail"`    // email
	Password  string   `json:"userPassword"` // password (SHA256 || Bcrypt)
	Type      string   `json:"userType"`     // type => admin, org, user
	CreatedAt datetime `json:"createdAt"`    // ...
	UpdatedAt datetime `json:"updatedAt"`    // ...
	Status    int      `json:"status"`       // 1 = active, 0 = deleted, 66 = banned
}
