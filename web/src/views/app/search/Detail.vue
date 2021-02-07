<template>
  <div>
    <a-card style="margin-top: 24px" :bordered="false" :title="data.name">
      <a-descriptions title="基本信息" :column="4">
        <a-descriptions-item label="Bundle Id">{{ data.bundle_id }}</a-descriptions-item>
        <a-descriptions-item label="最新版本">{{ data.current_version }}</a-descriptions-item>
        <a-descriptions-item label="平台">{{ data.platform_name }}</a-descriptions-item>
        <a-descriptions-item label="下载地址">{{ data.shortUrl }}</a-descriptions-item>
      </a-descriptions>
    </a-card>

    <a-card style="margin-top: 24px" :bordered="false" title="操作信息">
      <a-descriptions :column="4">
        <a-descriptions-item label="创建人">{{ data.created_by }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ data.created_at }}</a-descriptions-item>
        <a-descriptions-item label="最后操作人">{{ data.updated_by }}</a-descriptions-item>
        <a-descriptions-item label="最后更新时间">{{ data.updated_at | formatDate }}</a-descriptions-item>
        <a-descriptions-item ></a-descriptions-item>
        <a-descriptions-item ></a-descriptions-item>
      </a-descriptions>
    </a-card>
  </div>
</template>

<script>
import { StandardFormRow, ArticleListContent } from '@/components'
import IconText from './components/IconText'
import { mapActions } from 'vuex'

export default {
  components: {
    StandardFormRow,
    ArticleListContent,
    IconText
  },
  data () {
    return {
      id: 0,
      loading: true,
      loadingMore: false,
      data: {},
      form: this.$form.createForm(this)
    }
  },
  mounted () {
    this.id = this.$route.query.id
    this.getInfo()
  },
  methods: {
    ...mapActions(['GetAppInfo']),
    handleChange (value) {
      console.log(`selected ${value}`)
    },
    getInfo () {
      const { GetAppInfo } = this

      GetAppInfo(this.id).then(res => {
        this.data = res.result
        this.data.shortUrl = process.env.VUE_APP_BASE_URL + '/s/' + this.data.short_url
        console.log('data', this.data)
      }).catch((err) => {
        console.log('team list', err)
      })
    }
  }
}
</script>

<style lang="less" scoped>

</style>
