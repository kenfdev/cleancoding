<template>
  <todo-list-template :todos="todos"
                      :form-model="formModel"
                      :show-todo-form="isShowTodoForm" />
</template>

<script>
import Vue from 'vue';
import { mapGetters } from 'vuex';
import TodoListTemplate from '@/components/templates/TodoListTemplate';
import { name as todoNS, actionTypes as todoActions } from '@/store/todos';

export default Vue.extend({
  components: {
    TodoListTemplate,
  },
  computed: {
    ...mapGetters(todoNS, ['todos', 'isShowTodoForm', 'formModel']),
  },
  mounted() {
    this.$store.dispatch(`${todoNS}/${todoActions.FETCH_TODOS}`);
  },
});
</script>
