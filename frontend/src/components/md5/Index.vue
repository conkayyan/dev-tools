<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Text">
      <el-input v-model="form.txt" type="textarea"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmitEncode">Md5</el-button>
    </el-form-item>
    <el-form-item label="Result">
      <el-input v-model="result" type="textarea"/>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {Md5Encode} from '../../../wailsjs/go/main/App'

const form = reactive({
  txt: '',
})

const result = ref('')

const onSubmitEncode = () => {
  Md5Encode(form.txt).then(ret => {
    result.value = ret
  })
}

</script>

<style scoped>

</style>
