<template>
    <el-form :inline="true" class="demo-form-inline">
        <el-form-item label="Gen Nums">
            <el-input-number v-model="num" :min="1" />
        </el-form-item>
        <el-form-item>
            <el-checkbox v-model="isHyphen" label="Is Hyphen" border />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSubmit">Generate</el-button>
        </el-form-item>
    </el-form>
    <el-input v-model="result" type="textarea" :rows="20"/>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { GetUUID } from '../../../wailsjs/go/main/App'

const onSubmit = () => {
    result.value = ''
    for (let i = 0; i < num.value; i++) {
        GetUUID(isHyphen.value).then(ret => {
            result.value += ret + '\n'
        })
    }
}

const result = ref('')
const num = ref(1)
const isHyphen = ref(true)

</script>

<style>
.demo-form-inline .el-input {
    --el-input-width: 220px;
}

.demo-form-inline .el-select {
    --el-select-width: 220px;
}
</style>