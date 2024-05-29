<template>
  <div class="container">
    <div class="variety-list">
      <div
        v-for="future in mFutureList"
        :key="future.id"
        class="variety-item"
        :class="{active: future.ID === mActiveFuture.ID}"
        @click="onFutureClick(future)"
      >
        <div class="variety-item-title">{{ future.mainsymbol }}</div>
        <div class="variety-item-content">{{ future.exchange }} / {{ future.class }}</div>
      </div>
    </div>
    <div class="chart-container">
      <div class="chart-header">
<!--        <div class="chart-title">-->
<!--          {{ mActiveFuture }}-->
<!--        </div>-->
        <div class="flex">
          <div class="chart-tab">
            <el-date-picker
              v-model="mCurrentPredDate"
              type="date"
              size="small"
              :clearable="false"
              @change="onCurrentPredDateChange"
            />
            <!--            <el-radio-group v-model="mCurrentDuration" size="small" @change="onChartDurationChange">-->
            <!--              <el-radio-button v-for="duration in durations" :key="duration.value" :label="duration.label" :value="duration.value" />-->
            <!--            </el-radio-group>-->
          </div>
          <div class="chart-title-extra">
            <el-button size="small" :icon="FullScreen" circle @click="requestFullScreen" />
          </div>
        </div>
      </div>
      <div id="chart-content" class="chart-content">
        <VCharts
          :option="mEchartsOptions"
          :style="{ width: '100%', height: '100%' }"
        />
<!--        <VCharts-->
<!--          :option="mEchartsOptions"-->
<!--          :style="{ width: '500px', height: '300px' }"-->
<!--        />-->
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DashboardSimple'
}
</script>

<script setup>
import {  onMounted, ref } from 'vue'
import {
  FullScreen
} from '@element-plus/icons-vue'
import VCharts from 'vue-echarts'

import {
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DataZoomComponent,
  GraphicComponent,
  TitleComponent,
  DatasetComponent,
  TransformComponent,
} from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import { BarChart, LineChart, PieChart, RadarChart } from 'echarts/charts'
import { use } from 'echarts/core'
import { requestFutureClsList, requestPredList } from '@/api/xstock'
import dayjs from 'dayjs'
import { groupBy, uniq } from 'lodash-es'

use([
  CanvasRenderer,
  BarChart,
  LineChart,
  PieChart,
  RadarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DataZoomComponent,
  GraphicComponent,
  TitleComponent,
  DatasetComponent,
  TransformComponent,
])

// const durations = ref([
//   {
//     label: '1分钟',
//     value: 'New York'
//   },
//   {
//     label: '5分钟',
//     value: 'Washington',
//     disabled: true
//   },
//   {
//     label: '15分钟',
//     value: 'Los Angeles'
//   },
//   {
//     label: '30分钟',
//     value: 'Chicago'
//   },
//   {
//     label: '1小时',
//     value: 'Chicago'
//   },
//   {
//     label: '4小时',
//     value: 'Chicago'
//   },
//   {
//     label: '1天',
//     value: 'Chicago'
//   },
//   {
//     label: '7天',
//     value: 'Chicago'
//   }
// ])

const mCurrentPredDate = ref(new Date())

const onCurrentPredDateChange = async() => {
  await getPredList()
}

const getPredList = async() => {
  const res = await requestPredList({
    cls: mActiveFuture.value.class,
    date: dayjs(mCurrentPredDate.value).format('YYYY-MM-DD'),
  })
  const originalData = res.data?.records ?? []
  // const originalData = [
  //   { time: '08:00:01', y: 0.51, model: '008' },
  //   { time: '08:00:02', y: 0.52, model: '008' },
  //   { time: '08:00:03', y: 0.53, model: '008' },
  //   { time: '08:00:04', y: 0.54, model: '008' },
  //   { time: '08:00:05', y: 0.55, model: '008' },
  //   { time: '08:00:01', y: 0.61, model: '009' },
  //   { time: '08:00:02', y: 0.42, model: '009' },
  //   { time: '08:00:03', y: 0.33, model: '009' },
  //   { time: '08:00:04', y: 0.44, model: '009' },
  //   { time: '08:00:05', y: 0.65, model: '019' },
  //   { time: '08:00:06', y: 0.65, model: '019' },
  //   { time: '08:00:07', y: 0.65, model: '019' },
  //   { time: '08:00:08', y: 0.65, model: '019' },
  //   { time: '08:10:08', y: 0.65, model: '019' },
  //   { time: '22:10:08', y: 1.5, model: '029' },
  //   { time: '22:10:09', y: 1.9, model: '029' },
  // ]
  const data = groupBy(originalData, 'model')
  console.log('data..', data)
  setEchartsOptions(data)
}
// #region 左侧
const mFutureList = ref([])
const mActiveFuture = ref({})
const getFutureList = async() => {
  const res = await requestFutureClsList()
  mFutureList.value = (res.data?.records ?? []).sort((a, b) => a.showindex - b.showindex)
  console.log('res.', res)
  if (mFutureList.value.length) {
    mActiveFuture.value = mFutureList.value[0]
  }
}

const initPageData = async() => {
  await getFutureList()
  await getPredList()
}

const onFutureClick = async(future) => {
  console.log('onFutureClick', future)
  mActiveFuture.value = future
  await getPredList()
}

// #endregion
const requestFullScreen = () => {
  const element = document.getElementById('chart-content')
  if (element.requestFullscreen) {
    element.requestFullscreen()
  } else if (element.mozRequestFullScreen) {
    element.mozRequestFullScreen()
  } else if (element.webkitRequestFullScreen) {
    element.webkitRequestFullScreen()
  } else if (element.msRequestFullscreen) {
    element.msRequestFullscreen()
  }
}

onMounted(() => {
  initPageData()
})

const mEchartsOptions = ref({
  grid: {
    left: 20,
    right: 20,
    top: 20,
    bottom: 20,
    containLabel: true,
  },
  xAxis: {
    type: 'category',
    show: true,
    data: []
  },
  yAxis: {
    show: true,
    type: 'value',
    minInterval: 0.01,
  },
  tooltip: {
    show: true,
    trigger: 'axis',
  },
  legend: {
    data: []
  },
  series: [],
})

// const randomColor = ['#165DFF', '#FFA800', '#FF3A3A', '#00C2FF', '#00C2FF', '#00C2FF', '#00C2FF', '#00C2FF', '#00C2FF', '#00C2FF', '#00C2FF', '#00C2FF']

/**
 *
 * @param data {time: string; y: number; model: string}
 */
const setEchartsOptions = (data) => {
  console.log(data)
  const xAxisTmpData = []
  Object.keys(data).forEach(v => {
    data[v].forEach(v2 => {
      xAxisTmpData.push(v2.time)
    })
  })
  const xAxisData = uniq(xAxisTmpData).sort()

  const legend = Object.keys(data)

  mEchartsOptions.value = {
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    xAxis: {
      type: 'category',
      show: true,
      data: xAxisData
    },
    yAxis: {
      show: true,
      type: 'value',
      minInterval: 0.01,
    },
    tooltip: {
      show: true,
      trigger: 'axis',
    },
    legend: {
      data: legend,
    },
    series: legend.map((v, index) => {
      return {
        name: v,
        data: data[v].map(v2 => v2.y),
        type: 'line',
        showSymbol: false,
        // smooth: true,
        lineStyle: {
          // color: randomColor[index],
          width: 2,
        },
      }
    })
  }
}

</script>

<style lang="scss" scoped>
* {
  box-sizing: border-box;
}
.flex {
  display: flex;
}
.container {
  display: flex;
  height: 100vh;
  width: 100vw;
  background: white;
  padding: 4px;
}
.variety-list {
  width: 200px;
  overflow-y: auto;
  flex-shrink: 0;
  height: 100vh;
  .variety-item {
    height: 40px;
    border-radius: 4px;
    padding: 4px;
    margin-bottom: 4px;
    &:hover {
      background: var(--el-color-primary);
      color: white;
    }
    &.active {
      background: var(--el-color-primary);
      color: white;
    }
    cursor: pointer;
  }
}
.chart-container {
  flex: 1;
  padding: 20px;
  .chart-title {
    height: 40px;
  }
  .chart-tab {
    height: 40px;
    display: flex;
    align-items: center;
  }
  .chart-title-extra {
    align-items: center;
    margin-left: auto;
  }
  .chart-content {
    background: #fff;
    width: calc(100vw - 240px);
    height: calc(100vh - 140px);
    //width: 100%;
    //height: calc(100vh - 60px);
  }
}

</style>
