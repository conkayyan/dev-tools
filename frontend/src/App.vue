<template>
  <el-row class="tac">
    <el-col :lg="4" :span="6" class="border-right">
      <el-menu
          :default-active="activeIndex"
          class="no-border-right"
          @open="handleOpen"
          @close="handleClose"
      >
        <el-menu-item index="qrcode">
          <el-icon>
            <SvgIcon icon-class="qrcode"/>
          </el-icon>
          <span>QRCode</span>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="18" class="show-content">
      <Qrcode v-if="activeIndex=='qrcode'"/>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import Qrcode from './components/qrcode/Index.vue'
import SvgIcon from './components/svgIcon/Index.vue'

const activeIndex = ref('qrcode')
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
.show-content {
  height: 100vh;
  padding: 20px;
}

.border-right {
  border-right: 1px solid var(--el-border-color);
}

.no-border-right {
  border: none;
}
</style>
