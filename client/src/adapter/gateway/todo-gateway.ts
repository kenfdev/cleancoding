import { Todo } from '@/domain/todo';

export interface TodoGateway {
  findAllTodos(): Promise<Todo[]>;
  updateTodo(todo: Todo): Promise<Todo>;
  createTodo(todo: { title: string; description: string }): Promise<Todo>;
}
