<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Text">
      <el-input v-model="form.txt" type="textarea" :rows="10"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmitPretty">Pretty</el-button>
    </el-form-item>
    <el-form-item label="Result">
      <vue-json-pretty :data="result" :showLine="true" :showDoubleQuotes="true" theme="light"/>
    </el-form-item>
  </el-form>
</template>

<script lang="ts">
import VueJsonPretty from 'vue-json-pretty';
import 'vue-json-pretty/lib/styles.css';
import {reactive, ref} from "vue";

export default {
  name: 'MyComponent',
  components: {
    VueJsonPretty,
  },
  setup() {
    const form = reactive({
      txt: '{"key":"value"}',
    })

    const result = ref('')

    const onSubmitPretty = () => {
      result.value = JSON.parse(form.txt)
    }

    return {
      form,
      result,
      onSubmitPretty,
    }
  },
  mounted() {
    this.onSubmitPretty()
  }
};
</script>

<style scoped>
.vjs-tree {
  color: var(--el-text-color-regular);
}
</style>