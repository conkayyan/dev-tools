<template>
  <el-form :model="form" label-width="80px">
    <el-form-item label="Text or Url">
      <el-input v-model="form.txt" type="textarea"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Generate QRCode</el-button>
    </el-form-item>
    <el-form-item label="QRCode">
      <el-image :src="img">
        <template #error>
          <div class="image-slot">
            <el-icon>
              <icon-picture/>
            </el-icon>
          </div>
        </template>
      </el-image>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive, ref} from 'vue'
import {GenerateQRCode} from '../../wailsjs/go/main/App'
import {Picture as IconPicture} from '@element-plus/icons-vue'

const form = reactive({
  txt: '',
})

const img = ref('')

const onSubmit = () => {
  console.log('submit!')
  console.log(form.txt)
  GenerateQRCode(form.txt).then(result => {
    console.log(result)
    img.value = result
  })
}

</script>

<style scoped>
.el-form-item__content .el-image {
  width: 200px;
  height: 200px;
  border: solid 1px var(--el-border-color);
}

.el-form-item__content .image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 30px;
}

.el-form-item__content .image-slot .el-icon {
  font-size: 30px;
}
</style>
