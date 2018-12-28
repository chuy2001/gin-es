<template>
  <d2-container :filename="filename" type="full" class="page">
    <d2-grid-layout v-bind="layout" @layout-updated="layoutUpdatedHandler">
      <d2-grid-item
        v-for="(item, index) in layout.layout"
        :key="index"
        v-bind="item"
        :x="item.x"
        :y="item.y"
        :w="item.w"
        :h="item.h"
        @resize="resizeHandler"
        @move="moveHandler"
        @resized="resizedHandler"
        @moved="movedHandler"
      >
        <el-card shadow="hover" class="page_card">
          <div slot="header" class="clearfix">
            <span>卡片名称{{item.i}}</span>
            <el-button @click="changeType(index)" style="float: right; padding: 3  px 0" type="text" >切换图形</el-button>
          </div>
          <!-- <template v-if="item.i === '0'"> -->
            <ve-chart :data="chartData" :height="item.chartHight" :settings="item.chartSettings" ref="mychart"></ve-chart>
          <!-- </template> -->
        </el-card>
      </d2-grid-item>
    </d2-grid-layout>
  </d2-container>
</template>

<script>
import { Masters } from '@/api'
const DATA_FROM_BACKEND = {
  columns: ['date', 'PV', 'Order'],
  rows: [
    { 'date': '2018-05-22', 'PV': 32371, 'Order': 19810 },
    { 'date': '2018-05-23', 'PV': 12328, 'Order': 4398 },
    { 'date': '2018-05-24', 'PV': 92381, 'Order': 52910 }
  ]
}
const EMPTY_DATA = {
  columns: [],
  rows: []
}
export default {
  data () {
    this.typeArr = ['line', 'histogram', 'pie']
    this.index = 0
    return {
      filename: __filename,
      layout: {
        layout: [
          { x: 0, y: 0, w: 4, h: 10, i: '0', chartHight: '300px', chartSettings: { type: 'line', metrics: ['value'], dimension: ['key'] } },
          { x: 4, y: 0, w: 4, h: 10, i: '1', chartHight: '300px', chartSettings: { type: 'histogram' } },
          { x: 8, y: 0, w: 4, h: 10, i: '2', chartHight: '300px', chartSettings: { type: 'pie' } }
        ],
        colNum: 12,
        rowHeight: 30,
        isDraggable: true,
        isResizable: true,
        isMirrored: false,
        verticalCompact: true,
        margin: [5, 5],
        useCssTransforms: true
      },
      // chartData: {
      //   columns: ['日期', '访问用户', '下单用户', '下单率'],
      //   rows: [
      //     { 日期: '1/1', 访问用户: 1393, 下单用户: 1093, 下单率: 0.32 },
      //     { 日期: '1/2', 访问用户: 3530, 下单用户: 3230, 下单率: 0.26 },
      //     { 日期: '1/3', 访问用户: 2923, 下单用户: 2623, 下单率: 0.76 },
      //     { 日期: '1/4', 访问用户: 1723, 下单用户: 1423, 下单率: 0.49 },
      //     { 日期: '1/5', 访问用户: 3792, 下单用户: 3492, 下单率: 0.323 },
      //     { 日期: '1/6', 访问用户: 4593, 下单用户: 4293, 下单率: 0.78 }
      //   ]
      // },
      chartData: {
        columns: ['value', 'key_as_string', 'key', 'doc_count'],
        rows: []
      },
      loading: false,
      dataEmpty: false,
      chartSettings: { type: this.typeArr[this.index] },
      formData: ''
    }
  },
  created () {
    // this.formData = '{"index":"telegraf*"}' + '\n' + '{"query":{"match_all":{}}}' + '\n'
    this.formData = '{"search_type":"query_then_fetch","ignore_unavailable":true,"index":["tele*2018.12.28"]}' + '\n' + '{"size":0,"query":{"bool":{"filter":[{"range":{"@timestamp":{"gte":"1546004110960","lte":"1546004410960","format":"epoch_millis"}}},{"query_string":{"analyze_wildcard":true,"query":"*"}}]}},"aggs":{"2":{"date_histogram":{"interval":"1m","field":"@timestamp","min_doc_count":0,"extended_bounds":{"min":"1546004110960","max":"1546004410960"},"format":"epoch_millis"},"aggs":{"1":{"avg":{"field":"mem.available_percent"}}}}}}' + '\n'
    // this.getData()
  },
  mounted () {
    // 加载完成后显示提示
    this.showInfo()
    window.vue = this
  },
  methods: {
    log (arg1 = 'log', ...logs) {
      if (logs.length === 0) {
        console.log(arg1)
      } else {
        console.group(arg1)
        logs.forEach(e => {
          console.log(e)
        })
        console.groupEnd()
      }
    },
    // 显示提示
    showInfo () {
      this.$notify({
        title: '提示',
        message:
          '你可以按住卡片拖拽改变位置；或者在每个卡片的右下角拖动，调整卡片大小'
      })
    },
    // 测试代码
    layoutUpdatedHandler (newLayout) {
      console.group('layoutUpdatedHandler')
      newLayout.forEach(e => {
        console.log(
          `{'x': ${e.x}, 'y': ${e.y}, 'w': ${e.w}, 'h': ${e.h}, 'i': '${e.i}'},`
        )
      })
      console.groupEnd()
    },
    resizeHandler (i, newH, newW) {
      this.log('resizeHandler', `i: ${i}, newH: ${newH}, newW: ${newW}`)
    },
    moveHandler (i, newX, newY) {
      this.log('moveHandler', `i: ${i}, newX: ${newX}, newY: ${newY}`)
    },
    resizedHandler (i, newH, newW, newHPx, newWPx) {
      this.log(
        'resizedHandler',
        `i: ${i}, newH: ${newH}, newW: ${newW}, newHPx: ${newHPx}, newWPx: ${newWPx}`
      )
      this.layout.layout[i].chartHight = (newHPx - 40) + 'px'
      this.$refs.mychart[i].echarts.resize()
    },
    movedHandler (i, newX, newY) {
      this.log('movedHandler', `i: ${i}, newX: ${newX}, newY: ${newY}`)
    },
    changeType: function (index) {
      this.index++
      if (this.index >= this.typeArr.length) { this.index = 0 }
      this.layout.layout[index].chartSettings.type = this.typeArr[this.index]
      this.getESData()
    },
    getData () {
      this.loading = true
      // ajax get data ....
      setInterval(() => {
        // this.chartData = this.chartData.rows.length ? EMPTY_DATA : DATA_FROM_BACKEND
        this.chartData = DATA_FROM_BACKEND
        this.dataEmpty = !this.chartData.rows.length
        this.loading = false
      }, 3000)
    },
    getESData () {
      Masters.post(`es/msearch`, this.formData)
        .then(response => {
          // this.tables = response[`data`]
          //  this.tables = response.data.cmlist
          this.chartData.rows = response.data.Message[0].aggregations[2].buckets
          console.log('msg:' + this.chartData)
        })
        .catch(error => {
          console.log('err:' + error)
        })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '~@/assets/style/public.scss';
.page {
  .vue-grid-layout {
    background-color: $color-bg;
    border-radius: 4px;
    margin: -5px;
    .page_card {
      height: 100%;
      width: 100%;
      @extend %unable-select;
    }
    .vue-resizable-handle {
      opacity: 0.3;
      &:hover {
        opacity: 1;
      }
    }

    .clearfix:before,
    .clearfix:after {
      display: table;
      content: "";
    }
    .clearfix:after {
      clear: both;
    }
  }
}
</style>
