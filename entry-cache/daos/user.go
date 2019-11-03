package daos

import "encoding/json"

// User : struct
type User struct {
}

// NewUser :
func NewUser() *User {
	return &User{}
}

// GetUser : id
func (u *User) GetUser(ID uint32) (interface{}, error) {
	return redisClient.Do("GET", string(ID)).Result()
}

// SetUser : key, user
func (u *User) SetUser(key uint32, user interface{}) error {
	json, _ := json.Marshal(user)
	return redisClient.Do("SET", string(key), json).Err()
}
