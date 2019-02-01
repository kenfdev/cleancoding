import { Todo } from '@/domain/todo';

export interface CreateTodoRequestModel {
  todo: {
    title: string;
    description: string;
  };
  currentTodos: Todo[];
}
