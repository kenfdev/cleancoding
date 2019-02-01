import { TodoGateway } from '@/adapter/gateway/todo-gateway';
import { Todo } from '@/domain/todo';

export class TodoInMemoryGateway implements TodoGateway {
  todos: Todo[] = [
    { id: '1', title: 'todo1', description: 'description1', isCompleted: true },
  ];

  createTodo(todo: { title: string; description: string }): Promise<Todo> {
    throw new Error('Method not implemented.');
  }

  findAllTodos(): Promise<Todo[]> {
    return new Promise(resolve => {
      resolve(this.todos);
    });
  }

  updateTodo(todo: Todo): Promise<Todo> {
    return new Promise(resolve => {
      const index = this.todos.findIndex(t => t.id === todo.id);
      this.todos[index] = todo;
      resolve(todo);
    });
  }
}
