package zalupa

type Team struct {
	Id int `json:"id" db:"id"`
	ZalupasAmount int `json:"zalupasAmount" db:"zalupas_amount"`
	Name string `json:"name" db:"name"`
}