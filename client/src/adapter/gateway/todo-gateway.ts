import { Todo } from '@/domain/todo';

export interface TodoGateway {
  findAllTodos(): Promise<Todo[]>;
}
