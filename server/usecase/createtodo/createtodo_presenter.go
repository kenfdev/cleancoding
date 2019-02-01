package createtodo

type createTodoPresenter struct {
	viewModel *createTodoViewModel
}

func NewCreateTodoPresenter() CreateTodoOutputPort {
	return &createTodoPresenter{
		viewModel: &createTodoViewModel{},
	}
}

func (p *createTodoPresenter) GetViewModel() *createTodoViewModel {
	return p.viewModel
}

func (p *createTodoPresenter) Present(model *CreateTodoResponseModel) {
	p.viewModel.ID = model.Todo.ID
	p.viewModel.Title = model.Todo.Title
	p.viewModel.Description = model.Todo.Description
	p.viewModel.IsCompleted = model.Todo.IsCompleted
}
