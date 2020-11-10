<template>
  <page-header-wrapper
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    :tab-change="handleTabChange"
  >
    <template v-slot:content>
      <div class="ant-pro-page-header-search">
        <a-input-search size="large" style="width: 80%; max-width: 522px;">
          <template v-slot:enterButton>
            搜索
          </template>
        </a-input-search>
      </div>
    </template>
    <router-view />
  </page-header-wrapper>
</template>

<script>
const getActiveKey = (path) => {
  switch (path) {
    case '/app/search/article':
      return 'detail'
    case '/app/search/project':
      return 'version'
    case '/app/search/application':
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
        { key: 'detail', tab: '详细信息' },
        { key: 'version', tab: '版本列表' },
        { key: 'team', tab: '团队' }
      ],
      tabActiveKey: 'detail',
      search: true
    }
  },
  created () {
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
          this.$router.push('/app/search/article')
          break
        case 'version':
          this.$router.push('/app/search/project')
          break
        case 'team':
          this.$router.push('/app/search/application')
          break
        default:
          this.$router.push('/workplace')
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
