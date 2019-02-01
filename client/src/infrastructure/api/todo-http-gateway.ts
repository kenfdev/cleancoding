import { TodoGateway } from '@/adapter/gateway/todo-gateway';
import { Todo } from '@/domain/todo';
import { AxiosInstance } from 'axios';

export class TodoHttpGateway implements TodoGateway {
  private client: AxiosInstance;

  constructor(client: AxiosInstance) {
    this.client = client;
  }

  async findAllTodos(): Promise<Todo[]> {
    const result = await this.client.get('/todos');
    return result.data.todos;
  }

  async createTodo(todo: {
    title: string;
    description: string;
  }): Promise<Todo> {
    const result = await this.client.post(`/todos`, todo);
    return result.data;
  }
  async updateTodo(todo: Todo): Promise<Todo> {
    const result = await this.client.put(`/todos/${todo.id}/complete`, {
      isCompleted: todo.isCompleted,
    });
    return result.data;
  }
}
