<template>
  <page-header-wrapper
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    :tab-change="handleTabChange"
  >
    <template v-slot:content>
      <div class="ant-pro-page-header-search">

      </div>
    </template>
    <template v-slot:extra>
      <a-button-group style="margin-right: 4px;">
        <a-button>操作</a-button>
        <a-button>操作</a-button>
        <a-button><a-icon type="ellipsis"/></a-button>
      </a-button-group>
      <a-button type="primary" ><router-link :to="{ name: 'AppUpload', query:{ teamId: $route.query.team_id } }">上传APP</router-link></a-button>
    </template>
    <router-view />
  </page-header-wrapper>
</template>

<script>
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
    this.tabActiveKey = getActiveKey(this.$route.path)

    this.$watch('$route', (val) => {
      this.tabActiveKey = getActiveKey(val.path)
    })
  },
  methods: {
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
    }
  }
}
</script>

<style lang="less" scoped>
.ant-pro-page-header-search {
  text-align: center;
  margin-bottom: 16px;
}
</style>
