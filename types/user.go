package types

type User struct {
	ID    int64  `json:"id,omitempty"`
	Login string `json:"login"`
	Name  string `json:"name"`
}
