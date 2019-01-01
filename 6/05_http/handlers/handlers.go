package handlers

type Todo struct {
	Name      string `json:"name"`
	Done      bool   `json:"done"`
}

type Handler struct {
	Todos *[]Todo
}
