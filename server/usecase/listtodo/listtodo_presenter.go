package listtodo

type listTodoPresenter struct {
	viewModel *listTodoViewModel
}

func NewListTodoPresenter() ListTodoOutputPort {
	return &listTodoPresenter{
		viewModel: &listTodoViewModel{
			Todos: []*ViewableTodoSummary{},
		},
	}
}

func (p *listTodoPresenter) GetViewModel() *listTodoViewModel {
	return p.viewModel
}

func (p *listTodoPresenter) Present(model *ListTodoResponseModel) {
	var summaries []*ViewableTodoSummary
	for _, todo := range model.Todos {
		summary := &ViewableTodoSummary{
			ID:          todo.ID,
			Title:       todo.Title,
			IsCompleted: todo.IsCompleted,
		}
		summaries = append(summaries, summary)
	}

	p.viewModel.Todos = summaries
}
