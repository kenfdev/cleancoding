import { FetchTodosResponseModel } from './fetch-todos-responsemodel';
import { TodoGateway } from '@/adapter/gateway/todo-gateway';

export class FetchTodosInteractor {
  private todoGateway: TodoGateway;
  constructor(todoGateway: TodoGateway) {
    this.todoGateway = todoGateway;
  }
  async fetchTodos(): Promise<FetchTodosResponseModel> {
    const todos = await this.todoGateway.findAllTodos();
    return { todos };
  }
}
