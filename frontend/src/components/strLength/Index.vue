<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Text">
      <el-input v-model="form.txt" type="textarea" @change="onSubmit" @keyup="onSubmit"/>
    </el-form-item>
    <el-form-item label="Result">
      <el-input v-model="result" disabled placeholder="0"/>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {StrLength} from '../../../wailsjs/go/main/App'

const form = reactive({
  txt: '',
})

const result = ref(0)

const onSubmit = () => {
  StrLength(form.txt).then(ret => {
    result.value = ret
  })
}

</script>

<style scoped>

</style>
