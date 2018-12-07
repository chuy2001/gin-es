<template>
  <d2-container>
  <div style="width: 100%; margin: 0px auto;">
    <el-button size="medium" type="primary" @click="on_add_table" icon="el-icon-plus">添加</el-button>
    <el-table :data="tables" stripe=ture style="width: 100%" :default-sort = "{prop: 'creation_time', order: 'descending'}">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="表名">
              <span>{{ props.row.name }}</span>
            </el-form-item>
            <el-form-item label="别名">
              <span>{{ props.row.alias }}</span>
            </el-form-item>
            <el-form-item label="描述">
              <span>{{ props.row.readme }}</span>
            </el-form-item>
            <el-form-item label="ID">
              <span>{{ props.row.id }}</span>
            </el-form-item>
            <el-form v-for="(item) in props.row.fields" :key="item.id" size="small" :inline="true" label-width="60px">
              <el-form-item label="字段">
                <span>{{ item.name}}</span>
              </el-form-item>
              <el-form-item label="别名">
                <span>{{ item.alias}}</span>
              </el-form-item>
              <el-form-item label="描述">
                <span>{{ item.readme}}</span>
              </el-form-item>
              <el-form-item label="多值">
                <span>{{ item.is_multi}}</span>
              </el-form-item>
              <el-form-item label="必选">
                <span>{{ item.required}}</span>
              </el-form-item>
            </el-form>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column type="index" width="50">
      </el-table-column>
      <el-table-column label="表名" prop="name" sortable>
      </el-table-column>
      <el-table-column label="别名" prop="alias" sortable>
      </el-table-column>
      <el-table-column label="创建人" prop="creator_username" sortable>
      </el-table-column>
      <el-table-column label="创建时间" prop="creation_time" sortable>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" type="info" style="padding: 4px" @click="on_show_detail(scope.row)">详情</el-button>
          <el-button size="mini" type="warning" style="padding: 4px" @click="on_change(scope.row)">修改</el-button>
          <el-button size="mini" type="danger" style="padding: 4px" @click="on_delete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  </d2-container>
</template>
<script>
import Vue from 'vue'
import ChangeTable from './components/ChangeTable'

import { master, Masters } from '@/api'

export default {
  name: 'table-mgmt',
  data () {
    return {
      tables: [],
      add_location: {
        right: (window.innerWidth - 1200) / 2 > 20 ? Math.round((window.innerWidth - 1200) / 2) : 20,
        top: window.innerHeight - 140
      }
    }
  },
  mounted () {
    master.get(`mgmt/table`).then((response) => {
      // this.tables = response[`data`]
      //  this.tables = response.data.cmlist
      this.tables = response.data.data
    }).catch((error) => {
      console.log('err:' + error)
    })
  },
  methods: {
    on_show_detail (table) {
      console.log(table)
      var ShowTable = Vue.component(`ShowTable`, ChangeTable)
      var vm = new ShowTable({ propsData: { data: table } })
      vm.$mount()
      this.$el.appendChild(vm.$el)
    },
    on_change (table) {
      console.log(table)
      var Comp = Vue.component(`ChangeTable`, ChangeTable)
      var vm = new Comp({ propsData: { is_create: false, data: table } })
      vm.$mount()
      this.$el.appendChild(vm.$el)
    },
    on_delete (table) {
      console.log(table.id)
      this.$confirm(`删除表将删除表中所有数据，包括修改记录和已删除记录，是否继续？`, `警告`, { type: `warning` }).then(() => {
        var formData = JSON.parse(JSON.stringify(table))
        Masters.delete(`mgmt/table/${formData.id}`).then((response) => {
          this.$message.success(`删除成功`)
          this.tables.splice(this.tables.indexOf(table), 1)
        }).catch((error) => {
          if (error.response.data && error.response.data.detail) {
            this.$message.error(String(error.response.data.detail))
          }
        })
      })
    },
    on_add_table () {
      var Comp = Vue.component(`ChangeTable`, ChangeTable)
      var vm = new Comp({ propsData: { is_create: true, id: this.tables.length } })
      console.log(this.tables.length)
      vm.$mount()
      vm.$on(`add_table`, (table) => {
        this.tables.push(table)
        console.log(table)
      })
      this.$el.appendChild(vm.$el)
    }
  }
}
</script>
<style scoped>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 100;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 20%;
  }
</style>
