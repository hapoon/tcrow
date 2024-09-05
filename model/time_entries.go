package model

type ResponsePostTimeEntry struct {
	Id              int    `json:"id"`
	StartedAt       string `json:"started_at"`
	StoppedAt       string `json:"stopped_at"`
	TimeTrackableId int    `json:"time_trackable_id"`
}
