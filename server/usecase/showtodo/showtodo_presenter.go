package showtodo

type showTodoPresenter struct {
	viewModel *showTodoViewModel
}

func NewShowTodoPresenter() ShowTodoOutputPort {
	return &showTodoPresenter{
		viewModel: &showTodoViewModel{},
	}
}

func (p *showTodoPresenter) GetViewModel() *showTodoViewModel {
	return p.viewModel
}

func (p *showTodoPresenter) Present(model *ShowTodoResponseModel) {
	p.viewModel.ID = model.Todo.ID
	p.viewModel.Title = model.Todo.Title
	p.viewModel.Description = model.Todo.Description
}
