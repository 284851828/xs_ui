import service from '@/utils/request'

// import axios from 'axios'
// const instant = axios.create({
//   baseURL: 'http://121.199.78.191:8888',
// })

/**
 * 获取商品品种列表
 * @returns {*}
 */
export const requestFutureClsList = () => {
  return service({
    url: '/xstock/getFutureClsList',
    method: 'POST',
    data: {}
  })
}

export const requestPredList = (data) => {
  return service({
    url: '/xstock/getPredList',
    method: 'POST',
    data: data
  })
}

