<template>
  <svg
      aria-hidden="true"
      class="svg-icon"
      :style="`width:${size};height:${size}`">
    <use :xlink:href="symbolId" :fill="color"/>
  </svg>
</template>

<script setup lang="ts">
import {computed} from 'vue'

const props = defineProps({
  prefix: {
    type: String,
    default: 'icon'
  },
  iconClass: {
    type: String,
    default: ''
  },
  color: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: '18px'
  }
})

const symbolId = computed(() => `#${props.prefix}-${props.iconClass}`)

</script>

<style lang="scss" scoped>
.svg-icon {
  fill: currentColor;
//currentColor变量表示当前图标color值
}
</style>