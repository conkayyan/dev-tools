<template>
  <el-row class="tac">
    <el-col :lg="4" :span="6" class="border-right">
      <el-menu
          :default-active="activeIndex"
          class="no-border-right"
          @open="handleOpen"
          @close="handleClose"
          @select="handleSelect"
      >
        <el-menu-item index="qrcode">
          <el-icon>
            <SvgIcon icon-class="qrcode"/>
          </el-icon>
          <span>QRCode</span>
        </el-menu-item>
        <el-menu-item index="base64">
          <el-icon>
            <SvgIcon icon-class="base64"/>
          </el-icon>
          <span>Base64</span>
        </el-menu-item>
        <el-menu-item index="urlCode">
          <el-icon>
            <SvgIcon icon-class="urlCode"/>
          </el-icon>
          <span>UrlCode</span>
        </el-menu-item>
        <el-menu-item index="case">
          <el-icon>
            <SvgIcon icon-class="case"/>
          </el-icon>
          <span>ToUpper/Lower</span>
        </el-menu-item>
        <el-menu-item index="md5">
          <el-icon>
            <SvgIcon icon-class="md5"/>
          </el-icon>
          <span>Md5</span>
        </el-menu-item>
        <el-menu-item index="strLength">
          <el-icon>
            <SvgIcon icon-class="strLength"/>
          </el-icon>
          <span>Text Length</span>
        </el-menu-item>
        <el-menu-item index="time">
          <el-icon>
            <Clock/>
          </el-icon>
          <span>Time</span>
        </el-menu-item>
        <el-menu-item index="uuid">
          <el-icon>
            <CollectionTag />
          </el-icon>
          <span>UUID</span>
        </el-menu-item>
        <el-menu-item index="json">
          <el-icon>
            <SvgIcon icon-class="json"/>
          </el-icon>
          <span>Json</span>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="18" class="show-content">
      <Qrcode v-if="activeIndex=='qrcode'"/>
      <Base64 v-else-if="activeIndex=='base64'"/>
      <UrlCode v-else-if="activeIndex=='urlCode'"/>
      <Case v-else-if="activeIndex=='case'"/>
      <Md5 v-else-if="activeIndex=='md5'"/>
      <StrLength v-else-if="activeIndex=='strLength'"/>
      <Time v-else-if="activeIndex=='time'"/>
      <Uuid v-else-if="activeIndex=='uuid'"/>
      <Json v-else-if="activeIndex=='json'"/>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {ref} from "vue"
import {Clock, CollectionTag} from '@element-plus/icons-vue'
import Qrcode from './components/qrcode/Index.vue'
import SvgIcon from './components/svgIcon/Index.vue'
import Base64 from './components/base64/Index.vue'
import UrlCode from './components/urlCode/Index.vue'
import Case from './components/case/Index.vue'
import Md5 from './components/md5/Index.vue'
import StrLength from './components/strLength/Index.vue'
import Time from './components/time/Index.vue'
import Uuid from './components/uuid/Index.vue'
import Json from './components/json/Index.vue'

const activeIndex = ref('qrcode')
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  activeIndex.value = key
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
