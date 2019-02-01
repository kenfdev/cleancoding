<template>
  <v-list-tile>
    <v-list-tile-action>
      <v-checkbox v-model="checked"
                  @change="onCheckboxChange"></v-checkbox>
    </v-list-tile-action>

    <v-list-tile-content @click.prevent="onContentClick">
      <v-list-tile-title>
        {{ title }}
      </v-list-tile-title>
    </v-list-tile-content>
  </v-list-tile>
</template>

<script>
import Vue from 'vue';
export default Vue.extend({
  props: {
    isChecked: Boolean,
    title: String,
  },
  data() {
    return {
      checked: this.isChecked,
    };
  },
  watch: {
    isChecked(newVal, oldVal) {
      if (newVal !== this.checked) {
        this.checked = newVal;
      }
    },
  },
  methods: {
    onCheckboxChange(isChecked) {
      this.$emit('clickCheckbox', isChecked);
    },
    onContentClick() {
      this.$emit('clickContent');
    },
  },
});
</script>
