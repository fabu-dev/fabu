<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="应用列表">

      <div class="operate">
      </div>

      <a-list size="large">
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.env | getEnvName" style="flex: 0.25">
            <a-avatar slot="avatar" size="large" shape="square" :src="item.icon"/>{{ item.icon }}
            <a slot="title"><router-link target="_blank" :to="{ name: 'Preview', params:{ short_url: item.short_url } }"> {{ item.name }} </router-link></a>
          </a-list-item-meta>
          <div slot="actions">
            <router-link target="_blank" :to="{ name: 'Preview', params:{ short_url: item.short_url } }"><a>预览</a></router-link>
          </div>
          <div class="list-content">
            <div class="list-content-item" style="width:200px;">
              <span>BundleID</span>
              <p>{{ item.bundle_id }}</p>
            </div>
            <div class="list-content-item" style="width:50px;">
              <span>平台</span>
              <p>{{ item.platform_name }}</p>
            </div>
            <div class="list-content-item" style="width:100px;">
              <span>当前版本</span>
              <p>{{ item.current_version }}（{{ item.current_build }}）</p>
            </div>
            <div class="list-content-item" style="width:150px;">
              <span>更新时间</span>
              <p>{{ item.updated_at | formatDate }}</p>
            </div>
          </div>
        </a-list-item>
      </a-list>
    </a-card>
  </page-header-wrapper>
</template>

<script>
// 演示如何使用 this.$dialog 封装 modal 组件
import TaskForm from './modules/TaskForm'
import Info from './components/Info'
import { mapActions } from 'vuex'
import { RouteView } from '@/layouts'

export default {
  name: 'StandardList',
  components: {
    TaskForm,
    Info,
    RouteView
  },
  data () {
    return {
      data: []
    }
  },
  created () {
    this.getApp()
  },
  methods: {
    ...mapActions(['GetSquare']),
    getApp () {
      const { GetSquare } = this
      GetSquare().then(res => {
        this.data = res.result.app
        for (const key in this.data) {
          this.data[key].icon = process.env.VUE_APP_API_BASE_URL + '/' + this.data[key].icon
        }

        console.log('list data', this.data)
      }).catch((err) => {
        console.log('team list', err)
      })
    }
  }
}
</script>

<style lang="less" scoped>
.ant-avatar-lg {
    width: 48px;
    height: 48px;
    line-height: 48px;
}

.list-content-item {
    color: rgba(0, 0, 0, .45);
    display: inline-block;
    vertical-align: middle;
    font-size: 14px;
    margin-left: 40px;
    span {
        line-height: 20px;
    }
    p {
        margin-top: 4px;
        margin-bottom: 0;
        line-height: 22px;
    }
}
</style>
