package types

type User struct {
	ID    *int64  `bson:"id,omitempty" json:"id,omitempty"`
	Login *string `bson:"login" json:"login"`
	Name  *string `bson:"name" json:"name"`
}
