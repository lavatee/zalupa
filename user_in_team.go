package zalupa

type UserInTeam struct {
	Id int `json:"id" db:"id"`
	UserId int `json:"userId" db:"user_id"`
	TeamId int `json:"teamId" db:"team_id"`
	UserName string `json:"userName" db:"user_name"`
	TeamName string `json:"teamName" db:"team_name"`
	UserZalupasAmount int `json:"userZalupasAmount" db:"user_zalupas_amount"`
	UserProductivity int `json"userProductivity" db:"user_productivity"`
}