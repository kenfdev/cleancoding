import Vue from 'vue';
import Router from 'vue-router';
import TodoListPage from './pages/TodoListPage.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      redirect: '/todos',
    },
    {
      path: '/todos',
      name: 'TodoListPage',
      component: TodoListPage,
    },
    // {
    //   path: '/todos/:id',
    //   name: 'TodoDetailPage',
    //   // route level code-splitting
    //   // this generates a separate chunk (about.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () =>
    //     import(/* webpackChunkName: "about" */ './pages/TodoDetailPage.vue'),
    // },
  ],
});
