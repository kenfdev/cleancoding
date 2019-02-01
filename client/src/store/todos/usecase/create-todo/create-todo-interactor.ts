import { CreateTodoResponseModel } from './create-todo-responsemodel';
import { TodoGateway } from '@/adapter/gateway/todo-gateway';
import { CreateTodoRequestModel } from './create-todo-requestmodel';

export class CreateTodoInteractor {
  private todoGateway: TodoGateway;
  constructor(todoGateway: TodoGateway) {
    this.todoGateway = todoGateway;
  }
  async createTodo(
    requestModel: CreateTodoRequestModel
  ): Promise<CreateTodoResponseModel> {
    const { todo, currentTodos } = requestModel;

    const newTodo = await this.todoGateway.createTodo(todo);
    const newTodos = { ...currentTodos, newTodo };

    return { todos: newTodos };
  }
}
