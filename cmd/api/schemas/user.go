package schemas

import "encoding/json"

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (u User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
