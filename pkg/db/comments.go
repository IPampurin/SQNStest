package db

type Remark struct {
	ID        string `json:"id"`
	ComDate   string `json:"regdate"`
	Title     string `json:"login"`
	Comment   string `json:"comment"`
	Agreement bool   `json:"agreement"`
	UserID    string `json:"userid"`
}

/*
func AddRemark(user *Remark) (int64, error) {

}

func UpdateRemark(user *Remark) error {

}

func GetRemark(id string) (*Remark, error) {

}

func Remarks(id string, limit int, search string) ([]*Remark, error) {

}
*/
