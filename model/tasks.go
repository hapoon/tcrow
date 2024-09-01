package model

type Task struct {
	Id       int    `json:"id"`
	Key      string `json:"key"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	SafeUrl  string `json:"safe_url"`
	Label    string `json:"label"`
	TeamId   int    `json:"team_id"`
	ParentId int    `json:"parent_id"`
	RootId   int    `json:"root_id"`
}
