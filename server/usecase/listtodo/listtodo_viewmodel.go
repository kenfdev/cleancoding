package listtodo

type listTodoViewModel struct {
	Todos []*ViewableTodoSummary `json:"todos"`
}

type ViewableTodoSummary struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}
