<template>
  <page-header-wrapper
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    :tab-change="handleTabChange"
  >
    <template v-slot:content>
      <div class="detail-header">
        <div class="detail-header-top">
          <img class="icon" :src="data.icon">
          <p class="name">{{ data.name }}</p>
          <div class="appType-platform-wrapper">
            <div class="appType" v-show="2"></div>
            <div class="platform">
              <i :class="data.platform === '1' ? 'icon-ic_ios':'icon-ic_andr'"></i><span>适用于{{ data.platform_name }}系统</span>
            </div>
          </div>
        </div>
      </div>
    </template>
    <template v-slot:extra>
      <a-button-group style="margin-right: 4px;">
        <a-button>删除</a-button>
        <a-button><router-link target="_blank" :to="{ name: 'Preview', query:{ id: $route.query.id } }">预览</router-link></a-button>
        <a-button><a-icon type="ellipsis"/></a-button>
      </a-button-group>
      <a-button type="primary" ><router-link target="_blank" :to="{ name: 'AppUpload', query:{ teamId: $route.query.team_id } }">上传APP</router-link></a-button>
    </template>
    <router-view />
  </page-header-wrapper>
</template>

<script>
import { mapActions } from 'vuex'

const getActiveKey = (path) => {
  switch (path) {
    case '/app/search/detail':
      return 'detail'
    case '/app/search/version':
      return 'version'
    case '/app/search/team':
      return 'team'
    default:
      return 'detail'
  }
}
export default {
  name: 'SearchLayout',
  data () {
    return {
      data: {},
      tabList: [
        { key: 'detail', tab: '应用概述' },
        { key: 'version', tab: '版本列表' },
        { key: 'team', tab: '团队' }
      ],
      tabActiveKey: 'detail',
      search: true
    }
  },
  mounted () {
    this.getInfo()
    this.tabActiveKey = getActiveKey(this.$route.path)

    this.$watch('$route', (val) => {
      this.tabActiveKey = getActiveKey(val.path)
    })
  },
  methods: {
    ...mapActions(['GetAppInfo']),
    handleTabChange (key) {
      this.tabActiveKey = key
      switch (key) {
        case 'detail':
          this.$router.push({ path: '/app/search/detail', query: { id: this.$route.query.id, team_id: this.$route.query.team_id } })
          break
        case 'version':
          this.$router.push({ path: '/app/search/version', query: { id: this.$route.query.id, team_id: this.$route.query.team_id } })
          break
        case 'team':
          this.$router.push({ path: '/app/search/team', query: { id: this.$route.query.id, team_id: this.$route.query.team_id } })
          break
        default:
          this.$router.push({ path: '/app/search/detail', query: { id: this.$route.query.id, team_id: this.$route.query.team_id } })
      }
    },
    getInfo () {
      const { GetAppInfo } = this

      GetAppInfo(this.$route.query.id).then(res => {
        this.data = res.result
        this.data.shortUrl = process.env.VUE_APP_API_BASE_URL + '/' + this.data.short_url
        this.data.icon = process.env.VUE_APP_API_BASE_URL + '/' + this.data.icon
        console.log('search index data', this.data)
      }).catch((err) => {
        console.log('team list', err)
      })
    }
  }
}
</script>

<style lang="less" scoped>
.ant-pro-page-header-search {
  text-align: center;
  margin-bottom: 16px;
}
.detail-header {
  width: 100%;
  margin-top: 24px;
}
.detail-header-top {
  width: 100%;
  height: 148px;
  margin-bottom: 1px;
  background-color: white;
}
.detail-header-top {
  position: relative;
}
.detail-header-top .icon {
  position: absolute;
  top: 24px;
  left: 24px;
  width: 120px;
  height: 120px;
  background-size: cover;
  border-radius: 8px;
}
.detail-header-top .name {
  position: absolute;
  top: 54px;
  left: 150px;
  line-height: 24px;
  font-size: 24px;
  font-family: "PingFang SC";
}
.detail-header-top .appType-platform-wrapper {
  position: absolute;
  top: 95px;
  left: 120px;
  font-size: 0px;
}
.appType-platform-wrapper .appType {
  display: inline-block;
  line-height: 12px;
  font-size: 12px;
  padding: 3px 0px;
  border-radius: 2px;
  color: white;
  margin-right: 24px;
}
.appType-platform-wrapper .platform {
  display: inline-block;
  font-size: 14px;
  top: 84px;
  left: 130px;
}
</style>
