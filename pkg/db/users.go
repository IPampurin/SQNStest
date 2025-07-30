package db

type User struct {
	ID       string `json:"id"`
	RegDate  string `json:"regdate"`
	Username string `json:"username"`
	Login    string `json:"login"`
	EMail    string `json:"email"`
	Password string `json:"password"`
}

/*
func AddUser(user *User) (int64, error) {

}

func UpdateUser(user *User) error {

}

func GetUser(id string) (*User, error) {

}
*/
