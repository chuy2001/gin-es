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
          <div slot="header" >
            <span>卡片名称{{item.i}}</span>
            <el-button @click="changeType(index)" style="float: right; padding: 0  px 0" type="text" >切换图形</el-button>
          </div>
          <!-- <template v-if="item.i === '0'"> -->
          <div>
            <ve-chart :data="chartData" :width="item.chartWidth" :height="item.chartHight" :settings="item.chartSettings" ref="mychart"></ve-chart>
          </div>
          <!-- </template> -->
        </el-card>
      </d2-grid-item>
    </d2-grid-layout>
  </d2-container>
</template>

<script>
import { Masters } from '@/api'
export default {
  data () {
    this.typeArr = ['line', 'histogram', 'pie', 'bar', 'ring', 'waterfall', 'funnel', 'radar', 'map', 'heatmap']
    this.index = 0
    return {
      filename: __filename,
      layout: {
        layout: [
          { x: 0, y: 0, w: 4, h: 10, i: '0', chartWidth: '400px', chartHight: '300px', chartSettings: { type: 'line', labelMap: { 'v0': 'cpu' } } },
          { x: 4, y: 0, w: 4, h: 10, i: '1', chartWidth: '400px', chartHight: '300px', chartSettings: { type: 'histogram' } },
          { x: 8, y: 0, w: 4, h: 10, i: '2', chartWidth: '400px', chartHight: '300px', chartSettings: { type: 'pie' } }
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
      chartData: {
        columns: [],
        rows: []
      },
      loading: false,
      dataEmpty: false,
      chartSettings: { type: this.typeArr[this.index] },
      formData: '',
      timer: null
    }
  },
  created () {
    // this.formData = '{"index":"telegraf*"}' + '\n' + '{"query":{"match_all":{}}}' + '\n'
    this.formData = '{"search_type":"query_then_fetch","ignore_unavailable":true,"index":["tele*"]}' + '\n' + '{"size":0,"query":{"bool":{"filter":[{"range":{"@timestamp":{"gte":"1546004110960","lte":"1546004410960","format":"epoch_millis"}}},{"query_string":{"analyze_wildcard":true,"query":"*"}}]}},"aggs":{"2":{"date_histogram":{"interval":"30s","field":"@timestamp","min_doc_count":0,"extended_bounds":{"min":"1546004110960","max":"1546004410960"},"format":"epoch_millis"},"aggs":{"1":{"avg":{"field":"mem.available_percent"}}}}}}' + '\n'
    this.getData()
  },
  mounted () {
    // 加载完成后显示提示
    this.showInfo()
    window.vue = this
  },
  beforeDestroy () {
    if (this.timer) { // 如果定时器还在运行 或者直接关闭，不用判断
      clearInterval(this.timer) //  关闭
    }
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
      this.layout.layout[i].chartWidth = '100%'
      this.$refs.mychart[i].echarts.resize()
    },
    movedHandler (i, newX, newY) {
      this.log('movedHandler', `i: ${i}, newX: ${newX}, newY: ${newY}`)
    },
    changeType: function (index) {
      this.index++
      if (this.index >= this.typeArr.length) { this.index = 0 }
      this.layout.layout[index].chartSettings.type = this.typeArr[this.index]
      // this.getData()
    },
    getData () {
      this.loading = true
      // ajax get data ....
      this.getESData()
      this.timer = setInterval(() => {
        this.getESData()
      }, 10000)
    },
    getESData () {
      var d = new Date()
      var currentT = d.getTime()
      var req = this.formData
      var minT = /1546004110960/g
      req = req.replace(minT, currentT - 300000)
      var maxT = /1546004410960/g
      req = req.replace(maxT, currentT)
      Masters.post(`es/msearch`, req)
        .then(response => {
          var lengthRow = response.data.Message.length
          var lengthRows = response.data.Message[0].aggregations[2].buckets.length
          var tmpRows = []
          var tmpCol = ['date']
          d.getTime()
          if (lengthRow > 0 && lengthRows > 0) {
            for (var i = 0; i < lengthRows; i++) {
              var tmpRow = Object()
              tmpRow['date'] = new Date(response.data.Message[0].aggregations[2].buckets[i].key).toLocaleTimeString()
              for (var j = 0; j < lengthRow; j++) {
                tmpRow['v' + j] = response.data.Message[j].aggregations[2].buckets[i]['1'].value
              }
              tmpRows.push(tmpRow)
            }
            for (var k = 0; k < lengthRow; k++) {
              tmpCol.push('v' + k)
            }
          }
          this.chartData.rows = tmpRows
          this.chartData.columns = tmpCol
          // console.log('msg:' + this.chartData.rows)
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
