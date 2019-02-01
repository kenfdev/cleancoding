import { ToggleCompleteResponseModel } from './toggle-complete-responsemodel';
import { TodoGateway } from '@/adapter/gateway/todo-gateway';
import { ToggleCompleteRequestModel } from './toggle-complete-requestmodel';
import { Todo } from '@/domain/todo';

export class ToggleCompleteInteractor {
  private todoGateway: TodoGateway;
  constructor(todoGateway: TodoGateway) {
    this.todoGateway = todoGateway;
  }
  async toggleComplete(
    requestModel: ToggleCompleteRequestModel
  ): Promise<ToggleCompleteResponseModel> {
    const { originalTodos, newCompletetionValue, targetTodo } = requestModel;

    const updatedTodo: Todo = {
      ...targetTodo,
      isCompleted: newCompletetionValue,
    };
    const newTodo = await this.todoGateway.updateTodo(updatedTodo);

    const index = originalTodos.findIndex(t => t.id === newTodo.id);
    const newTodos = [
      ...originalTodos.slice(0, index),
      newTodo,
      ...originalTodos.slice(index + 1),
    ];
    return { todos: newTodos };
  }
}
