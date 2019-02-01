import { GetterTree, ActionTree, MutationTree } from 'vuex';
import { RootState } from '@/store';
import { FetchTodosInteractor } from './usecase/fetch-todos/fetch-todos-interactor';
import { Todo } from '@/domain/todo';
import { MyStore } from '@/store';

export const name = 'todos';

export interface State {
  todos: Todo[];
}

export const state = (): State => ({
  todos: [],
});

export const getters: GetterTree<State, RootState> = {
  todos(state) {
    return state.todos;
  },
};

export const actionTypes = {
  FETCH_TODOS: 'FETCH_TODOS',
};

export const mutationTypes = {
  SET_TODOS: 'SET_TODOS',
};

export const actions: ActionTree<State, RootState> = {
  [actionTypes.FETCH_TODOS](context) {
    const { todoGateway } = (this as MyStore).$adapters;
    const interactor = new FetchTodosInteractor(todoGateway);
    return interactor.fetchTodos().then(res => {
      context.commit(mutationTypes.SET_TODOS, res.todos);
    });
  },
};

export const mutations: MutationTree<State> = {
  [mutationTypes.SET_TODOS](state, todos: Todo[]) {
    state.todos = todos;
  },
};
