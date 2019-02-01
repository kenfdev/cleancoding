import { Todo } from '@/domain/todo';

export interface ToggleCompleteRequestModel {
  originalTodos: Todo[];
  targetTodo: Todo;
  newCompletetionValue: boolean;
}
