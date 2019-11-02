package daos

// User : struct
type User struct {
}

// NewUser :
func NewUser() *User {
	return &User{}
}

// GetUser : id
func (u *User) GetUser(ID uint) (interface{}, error) {
	return redisClient.Get(string(ID)).Result()
}

// SetUser : key, user
func (u *User) SetUser(key uint, user interface{}) error {
	return redisClient.Set(string(key), user, 0).Err()
}
