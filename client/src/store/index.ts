import Vue from 'vue';
import Vuex, { Store } from 'vuex';
import { TodoGateway } from '@/adapter/gateway/todo-gateway';
import { TodoInMemoryGateway } from '@/infrastructure/api/todo-inmemory-gateway';

Vue.use(Vuex);

export interface RootState {
  msg: string;
}

export interface Adapters {
  todoGateway: TodoGateway;
}

export interface MyStore extends Store<RootState> {
  $adapters: Adapters;
}

const store = new Vuex.Store({
  state: {
    msg: '',
  },
  actions: {
    async msgAction({ commit }, msg) {
      const promise = new Promise(resolve => {
        setTimeout(() => {
          resolve(msg);
        }, 2500);
      });

      const result = await promise;
      commit('setMsg', result);
    },
  },
  mutations: {
    setMsg(state, msg) {
      state.msg = msg;
    },
  },
});

(store as any).$adapters = {
  todoGateway: new TodoInMemoryGateway(),
} as Adapters;

export default store;
