import Vue from 'vue';
import Vuex, { Store } from 'vuex';
import axios from 'axios';
import { TodoGateway } from '@/adapter/gateway/todo-gateway';
// import { TodoInMemoryGateway } from '@/infrastructure/api/todo-inmemory-gateway';

import * as todosModule from './todos';
import { TodoHttpGateway } from '@/infrastructure/api/todo-http-gateway';

Vue.use(Vuex);

export interface RootState {}

export interface Adapters {
  todoGateway: TodoGateway;
}

export interface MyStore extends Store<RootState> {
  $adapters: Adapters;
}

const store = new Vuex.Store({
  state: {},
  actions: {},
  mutations: {},
  modules: {
    [todosModule.name]: todosModule,
  },
});
const axiosInstance = axios.create({
  baseURL: '/api',
});

(store as any).$adapters = {
  // todoGateway: new TodoInMemoryGateway(),
  todoGateway: new TodoHttpGateway(axiosInstance),
} as Adapters;

export default store;
