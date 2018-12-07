import request from '@/plugin/axios'

export function BusinessTable1List (data) {
  return request({
    url: '/business/table',
    method: 'post',
    data
  })
}
