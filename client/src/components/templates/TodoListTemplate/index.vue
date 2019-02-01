<template>
  <v-container>
    <v-layout text-xs-center
              wrap>
      <v-flex mb-4>
        <todo-list class="todo-list"
                   :todos="todos"
                   @toggleComplete="onToggleComplete">

          <v-fab-transition>
            <v-btn color="pink"
                   dark
                   absolute
                   bottom
                   right
                   fab
                   @click="onClickAdd">
              <v-icon>add</v-icon>
            </v-btn>
          </v-fab-transition>
        </todo-list>
      </v-flex>
    </v-layout>
    <v-dialog v-model="showTodoForm"
              persistent
              max-width="600px">
      <todo-form :form-model="formModel"
                 @clickClose="onClickClose"
                 @clickAdd="onClickSave"
                 @change="onChangeFormModel" />
    </v-dialog>
  </v-container>
</template>

<script>
import Vue from 'vue';
import TodoList from '@/components/organisms/TodoList';
import TodoForm from '@/components/organisms/TodoForm';
import { name as todoNS, actionTypes as todoActions } from '@/store/todos';

export default Vue.extend({
  props: {
    formModel: Object,
    todos: Array,
    showTodoForm: Boolean,
  },
  components: {
    TodoList,
    TodoForm,
  },
  methods: {
    onToggleComplete(todo, isCompleted) {
      this.$store.dispatch(`${todoNS}/${todoActions.TOGGLE_COMPLETE_TODO}`, {
        targetTodo: todo,
        isCompleted,
      });
    },
    onClickClose() {
      this.$store.dispatch(`${todoNS}/${todoActions.CLOSE_FORM}`);
    },
    onClickAdd() {
      this.$store.dispatch(`${todoNS}/${todoActions.OPEN_FORM}`);
    },
    onChangeFormModel(model) {
      this.$store.dispatch(`${todoNS}/${todoActions.CHANGE_FORM_MODEL}`, model);
    },
    onClickSave() {
      this.$store.dispatch(`${todoNS}/${todoActions.CREATE_TODO}`);
    },
  },
});
</script>

<style lang="scss" scoped>
.todo-list {
  position: relative;
}
</style>