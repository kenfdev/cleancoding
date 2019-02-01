import { GetterTree, ActionTree, MutationTree } from 'vuex';
import { RootState } from '@/store';
import { FetchTodosInteractor } from './usecase/fetch-todos/fetch-todos-interactor';
import { Todo } from '@/domain/todo';
import { MyStore } from '@/store';
import { ToggleCompleteInteractor } from './usecase/toggle-complete-todo/toggle-complete-interactor';
import { ToggleCompleteRequestModel } from './usecase/toggle-complete-todo/toggle-complete-requestmodel';
import { CreateTodoInteractor } from './usecase/create-todo/create-todo-interactor';
import { CreateTodoRequestModel } from './usecase/create-todo/create-todo-requestmodel';

export const namespaced = true;
export const name = 'todos';

export interface State {
  isShowTodoForm: boolean;
  todos: Todo[];
  formModel: {
    title: string;
    description: string;
  };
}

export const state = (): State => ({
  isShowTodoForm: false,
  todos: [],
  formModel: {
    title: '',
    description: '',
  },
});

export const getters: GetterTree<State, RootState> = {
  isShowTodoForm(state) {
    return state.isShowTodoForm;
  },
  todos(state) {
    return state.todos;
  },
  formModel(state) {
    return state.formModel;
  },
};

export const actionTypes = {
  FETCH_TODOS: 'FETCH_TODOS',
  TOGGLE_COMPLETE_TODO: 'TOGGLE_COMPLETE_TODO',
  CLOSE_FORM: 'CLOSE_FORM',
  OPEN_FORM: 'OPEN_FORM',
  CHANGE_FORM_MODEL: 'CHANGE_FORM_MODEL',
  CREATE_TODO: 'CREATE_TODO',
};

export const mutationTypes = {
  SET_TODOS: 'SET_TODOS',
  SET_IS_SHOW_TODO_FORM: 'SET_IS_SHOW_TODO_FORM',
  SET_FORM_MODEL: 'SET_FORM_MODEL',
};

export const actions: ActionTree<State, RootState> = {
  async [actionTypes.OPEN_FORM](context) {
    context.commit(mutationTypes.SET_IS_SHOW_TODO_FORM, true);
  },
  async [actionTypes.CLOSE_FORM](context) {
    context.commit(mutationTypes.SET_IS_SHOW_TODO_FORM, false);
  },
  async [actionTypes.CHANGE_FORM_MODEL](context, formModel: any) {
    context.commit(mutationTypes.SET_FORM_MODEL, formModel);
  },
  async [actionTypes.FETCH_TODOS](context) {
    const { todoGateway } = (this as MyStore).$adapters;

    const interactor = new FetchTodosInteractor(todoGateway);
    const res = await interactor.fetchTodos();

    context.commit(mutationTypes.SET_TODOS, res.todos);
  },
  async [actionTypes.CREATE_TODO](context) {
    const { todoGateway } = (this as MyStore).$adapters;

    const interactor = new CreateTodoInteractor(todoGateway);
    const req: CreateTodoRequestModel = {
      todo: context.state.formModel,
      currentTodos: context.state.todos,
    };
    const res = await interactor.createTodo(req);

    context.commit(mutationTypes.SET_TODOS, res.todos);
    context.commit(mutationTypes.SET_IS_SHOW_TODO_FORM, false);
    context.commit(mutationTypes.SET_FORM_MODEL, {
      title: '',
      description: '',
    });
  },
  async [actionTypes.TOGGLE_COMPLETE_TODO](
    context,
    { targetTodo, isCompleted }
  ) {
    const { todoGateway } = (this as MyStore).$adapters;

    const interactor = new ToggleCompleteInteractor(todoGateway);
    const req: ToggleCompleteRequestModel = {
      originalTodos: context.state.todos,
      targetTodo,
      newCompletetionValue: isCompleted,
    };
    const res = await interactor.toggleComplete(req);

    context.commit(mutationTypes.SET_TODOS, res.todos);
  },
};

export const mutations: MutationTree<State> = {
  [mutationTypes.SET_TODOS](state, todos: Todo[]) {
    state.todos = todos;
  },
  [mutationTypes.SET_IS_SHOW_TODO_FORM](state, isShow: boolean) {
    state.isShowTodoForm = isShow;
  },
  [mutationTypes.SET_FORM_MODEL](state, formModel: any) {
    state.formModel = formModel;
  },
};
