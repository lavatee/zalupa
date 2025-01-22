package zalupa

type User struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	ZalupaId string `json:"zalupaId" db:"zalupa_id"`
}