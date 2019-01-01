package structs

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Children []User `json:"children"`
}

type Task struct {
	Title string `json:"title"`
	Done bool `json:"done"`
	Important bool `json:"important"`
}

type Json struct {
	User User `json:"user"`
	Auth struct{
		SessionID string `json:"session_id"`
		CsrfToken string `json:"csrf_token"`
	} `json:"auth"`
	Tasks []Task
}
