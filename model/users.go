package model

type UserWorking struct {
	IsWorking bool `json:"is_working"`
}

type UserInfo struct {
	Id               int    `json:"id"`
	Email            string `json:"email"`
	Nickname         string `json:"nickname"`
	Image            string `json:"image"`
	AvatarUrl        string `json:"avatar_url"`
	Daily            bool   `json:"daily"`
	Weekly           bool   `json:"weekly"`
	NotifyExported   bool   `json:"notify_exported"`
	CalendarZoomRate int    `json:"calendar_zoom_rate"`
	DailyTime        string `json:"daily_time"`
	IsAnonymous      bool   `json:"is_anonymous"`
	HasPassword      bool   `json:"has_password"`
	HtmlUrl          string `json:"html_url"`
	IsDeactivated    bool   `json:"is_deactivated"`
	TimeZone         string `json:"time_zone"`
}
