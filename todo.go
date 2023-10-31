package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}