package toggletodo

type toggleTodoPresenter struct {
	viewModel *toggleTodoViewModel
}

func NewToggleTodoPresenter() ToggleTodoOutputPort {
	return &toggleTodoPresenter{
		viewModel: &toggleTodoViewModel{},
	}
}

func (p *toggleTodoPresenter) GetViewModel() *toggleTodoViewModel {
	return p.viewModel
}

func (p *toggleTodoPresenter) Present(model *ToggleTodoResponseModel) {
	p.viewModel.ID = model.Todo.ID
	p.viewModel.Title = model.Todo.Title
	p.viewModel.Description = model.Todo.Description
	p.viewModel.IsCompleted = model.Todo.IsCompleted
}
