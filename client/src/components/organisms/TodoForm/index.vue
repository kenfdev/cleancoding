<template>
  <v-card>
    <v-card-title>
      <span class="headline">Add Todo</span>
    </v-card-title>
    <v-card-text>
      <v-container grid-list-md>
        <v-layout wrap>
          <v-flex xs12>
            <v-text-field label="Title*"
                          required
                          :value="formModel.title"
                          @input="onChangeTitle" />
          </v-flex>
          <v-flex xs12>
            <v-text-field label="Description*"
                          required
                          :value="formModel.description"
                          @input="onChangeDescription" />
          </v-flex>
        </v-layout>
      </v-container>
      <small>*indicates required field</small>
    </v-card-text>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="blue darken-1"
             flat
             @click="onClickClose">Close</v-btn>
      <v-btn color="blue darken-1"
             flat
             @click="onClickAdd">Add</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  props: {
    formModel: Object,
  },
  methods: {
    onClickClose() {
      this.$emit('clickClose');
    },
    onClickAdd() {
      this.$emit('clickAdd');
    },
    onChangeTitle(value) {
      this.$emit('change', { ...this.formModel, title: value });
    },
    onChangeDescription(value) {
      this.$emit('change', { ...this.formModel, description: value });
    },
  },
});
</script>
