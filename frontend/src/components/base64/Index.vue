<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Text">
      <el-input v-model="form.txt" type="textarea" :rows="10" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmitEncode">Encrypt Base64</el-button>
      <el-button type="primary" @click="onSubmitDecode">Decrypt Base64</el-button>
    </el-form-item>
    <el-form-item label="Result">
      <el-input v-model="result" type="textarea" :rows="10" />
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {Base64Decode, Base64Encode} from '../../../wailsjs/go/main/App'

const form = reactive({
  txt: '',
})

const result = ref('')

const onSubmitEncode = () => {
  Base64Encode(form.txt).then(ret => {
    result.value = ret
  })
}

const onSubmitDecode = () => {
  Base64Decode(form.txt).then(ret => {
    result.value = ret
  })
}

</script>

<style scoped>

</style>
