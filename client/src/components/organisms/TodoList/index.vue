<template>
  <div>
    <v-list>
      <template v-for="(todo, index) in todos">
        <checkbox-list-item :key="index"
                            :is-checked="todo.isCompleted"
                            :title="todo.title"
                            @clickCheckbox="onClickCheckbox(todo, $event)"
                            @clickContent="onClickContent(todo)" />
      </template>
      <slot></slot>
    </v-list>
  </div>
</template>

<script>
import Vue from 'vue';
import CheckboxListItem from '@/components/molecules/CheckboxListItem';

export default Vue.extend({
  props: {
    todos: Array,
  },
  components: {
    CheckboxListItem,
  },
  methods: {
    onClickCheckbox(todo, isCompleted) {
      this.$emit('toggleComplete', todo, isCompleted);
    },
    onClickContent(todo) {
      console.log('content clicked', todo);
    },
  },
});
</script>

