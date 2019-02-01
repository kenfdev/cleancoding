import { TodoGateway } from '@/adapter/gateway/todo-gateway';
import { Todo } from '@/domain/todo';

export class TodoInMemoryGateway implements TodoGateway {
  todos: Todo[] = [{ id: '1', title: 'todo1', description: 'description1' }];

  findAllTodos(): Promise<Todo[]> {
    return new Promise(resolve => {
      resolve(this.todos);
    });
  }
}
