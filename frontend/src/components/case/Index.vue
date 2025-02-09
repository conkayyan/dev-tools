<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Text">
      <el-input v-model="form.txt" type="textarea"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmitToUpper">ToUpper</el-button>
      <el-button type="primary" @click="onSubmitToLower">ToLower</el-button>
    </el-form-item>
    <el-form-item label="Result">
      <el-input v-model="result" type="textarea"/>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'

const form = reactive({
  txt: '',
})

const result = ref('')

const onSubmitToUpper = () => {
  result.value = form.txt.toUpperCase()
}

const onSubmitToLower = () => {
  result.value = form.txt.toLowerCase()
}

</script>

<style scoped>

</style>
