<template>
  <VCharts
    :option="option"
    :style="{ width, height }"
  />
</template>

<script>
export default {
  name: 'Chart'
}
</script>

<script setup>
import VCharts from 'vue-echarts'

defineProps({
  width: {
    type: String,
    default: '100%'
  },
  height: {
    type: String,
    default: '100%'
  },
  option: {
    type: Object,
    default: () => ({})
  }
})
</script>
