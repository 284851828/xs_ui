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
        <div class="flex">
          <div class="chart-tab">
            <el-radio-group v-model="mChartType" size="mini" @change="onChartTypeChange">
              <el-radio-button label="daily">日日线</el-radio-button>
              <el-radio-button label="day">分时线</el-radio-button>
            </el-radio-group>
            <el-date-picker
              v-if="mChartType === 'day'"
              v-model="mCurrentPredDate"
              type="date"
              size="mini"
              :clearable="false"
              @change="onCurrentPredDateChange"
              style="margin-left: 10px;"
            />
            <div v-if="mChartType === 'day'" class="ml-10">
              {{ mMoc900Data }}
            </div>
          </div>
          <div class="chart-title-extra">
            <el-button size="mini" :icon="FullScreen" circle @click="requestFullScreen" />
          </div>
        </div>
      </div>
      <div id="chart-content" class="chart-content">
        <VCharts
          :option="mEchartsOptions"
          :style="{ width: '100%', height: '100%' }"
        />
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
import { onMounted, onUnmounted, ref } from 'vue'
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
import { requestFutureClsList, requestPredList, requestPred0000ListList } from '@/api/xstock'
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

const mCurrentPredDate = ref(new Date())

const onCurrentPredDateChange = async() => {
  await getPredList()
}

const mMoc900Data = ref('');

const getPredList = async() => {
  if (mChartType.value === 'day') {
    const res = await requestPredList({
      cls: mActiveFuture.value.class,
      date: dayjs(mCurrentPredDate.value).format('YYYY-MM-DD'),
    });

    const moc900Data = res.data?.moc900 ?? {};
    mMoc900Data.value = Object.entries(moc900Data)
    .map(([key, value]) => `${key} : ${value}`)
    .join(' | ');

    const originalData = (res.data?.records ?? []).map(v => ({ time: v.time, y: v.y, model: v.model }));
    const data = groupBy(originalData, 'model');
    setEchartsOptions(data);
    initTimer();
  } else if (mChartType.value === 'daily') {
    const res = await requestPred0000ListList({
      cls: mActiveFuture.value.class,
      date: dayjs(mCurrentPredDate.value).format('YYYY-MM-DD'),
    });
    const originalData = (res.data?.records ?? []).map(v => ({ date: v.date, y: v.y, model: v.model }));
    const data = groupBy(originalData, 'model');
    setEchartsOptionsByDaily(data)
    initTimer()
    console.log('onChartTypeChange')
  }
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

const completeTimeSeries = (startTimeArray, endTimeStr) => {
  const endTime = new Date(`1970-01-01 ${endTimeStr}:00`)
  let result = startTimeArray.map(timeStr => {
    const [hours, minutes, seconds] = timeStr.split(':')
    return new Date(`1970-01-01 ${hours}:${minutes}:${seconds}`)
  })
  // 获取最后一个起始时间，并初始化为第一个时间的整分钟版本
  const lastTime = new Date(result.length ? result[result.length - 1].setSeconds(0, 0) : '1970-01-01 09:00:00') // 确保从整分钟开始
  while (lastTime < endTime) {
    lastTime.setMinutes(lastTime.getMinutes() + 1)
    lastTime.setSeconds(0) // 重置秒数为0
    result.push(new Date(lastTime))
  }
  result = result.map(date => {
    let hours = date.getHours()
    let minutes = date.getMinutes()
    let seconds = date.getSeconds()
    hours = hours < 10 ? '0' + hours : hours
    minutes = minutes < 10 ? '0' + minutes : minutes
    seconds = seconds < 10 ? '0' + seconds : seconds
    return `${hours}:${minutes}:${seconds}`
  })

  return result
}

/**
 *
 * @param data {time: string; y: number; model: string}
 */
const setEchartsOptions = (data) => {
  const xAxisTmpData = []
  Object.keys(data).forEach(v => {
    data[v].forEach(v2 => {
      xAxisTmpData.push(v2.time)
    })
  })

  const showSymbol = Object.keys(data).map(v => data[v].length).filter(v => v < 10).length > 0
  let xAxisData = uniq(xAxisTmpData).sort()
  const lastTime = xAxisData[xAxisData.length - 1]

  xAxisData = completeTimeSeries(xAxisData, '15:00')

  const legend = Object.keys(data)

  const series = []
  Object.keys(data).forEach((v) => {
    const seriesData = []
    let latest = 0
    // 在当前时间内绘制，超出后暂不绘制
    for (let i = 0; i < xAxisData.length; i++) {
      const target = data[v].find(v2 => v2.time === xAxisData[i])
      if (target) {
        latest = target.y
      }
      if (xAxisData[i] > lastTime) {
        seriesData.push(undefined)
      } else {
        seriesData.push(target ? target.y : latest)
      }
    }
    series.push({
      name: v,
      type: 'line',
      data: seriesData,
      showSymbol,
      lineStyle: {
        width: 2,
      },
      symbolSize: 4,
    })
  })

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
    series,
  }
}

const setEchartsOptionsByDaily = (data) => {
  const xAxisTmpData = []
  Object.keys(data).forEach(v => {
    data[v].forEach(v2 => {
      xAxisTmpData.push(v2.date)
    })
  });

  // console.log('setEchartsOptionsByDaily..', xAxisTmpData);

  let xAxisData = uniq(xAxisTmpData).sort()

  const lastTime = xAxisData[xAxisData.length - 1]

  const legend = Object.keys(data)

  const series = []
  Object.keys(data).forEach((v) => {
    const seriesData = []
    let latest = 0
    // 在当前时间内绘制，超出后暂不绘制
    for (let i = 0; i < xAxisData.length; i++) {
      const target = data[v].find(v2 => v2.date === xAxisData[i])
      if (target) {
        latest = target.y
      }
      if (xAxisData[i] > lastTime) {
        seriesData.push(undefined)
      } else {
        seriesData.push(target ? target.y : latest)
      }
    }
    series.push({
      name: v,
      type: 'line',
      data: seriesData,
      lineStyle: {
        width: 2,
      },
      symbolSize: 4,
    })
  })

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
    series,
  }
}

const timer = ref(null)
// 轮询数据
const initTimer = () => {
  if (timer.value) {
    clearInterval(timer.value)
  }
  timer.value = setInterval(() => {
    getPredList()
  }, 1000 * 10)
}

onMounted(() => {
  initTimer()
})

onUnmounted(() => {
  clearInterval(timer)
})

const mChartType = ref('daily') // daily day
const onChartTypeChange = (e) => {
  getPredList();
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
  height: 100%;
  width: 100%;
  background: white;
}
.variety-list {
  width: 200px;
  overflow-y: auto;
  flex-shrink: 0;
  height: 100%;
  box-shadow: 0px 0px 5px 3px rgb(0 0 0 / 10%);
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
    width: calc(100vw - 500px);
    height: calc(100% - 140px);
    //width: 100%;
    //height: calc(100vh - 60px);
  }
}

.ml-10 {
  margin-left: 10px;
}

</style>
