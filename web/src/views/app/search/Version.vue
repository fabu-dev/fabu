<template>
  <div>
    <a-table :columns="columns" :data-source="data" rowKey="id" bordered>
      <a slot="name" slot-scope="text">{{ text }}</a>
      <span slot="customTitle"><a-icon type="smile-o" /> Name</span>
      <span slot="size" slot-scope="size">
        {{ (size/1024/1024).toFixed(2) }}
      </span>
      <span slot="action">
        <a>下载</a>
        <a-divider type="vertical" />
        <a>删除</a>
      </span>
    </a-table>
  </div>
</template>

<script>
import moment from 'moment'
import { mapActions } from 'vuex'

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    slots: { title: 'customTitle' }
  },
  {
    title: '版本号',
    dataIndex: 'code',
    key: 'code',
    slots: { title: 'customTitle' }
  },
  {
    title: '更新说明',
    dataIndex: 'description',
    key: 'description'
  },
  {
    title: '大小M',
    dataIndex: 'size',
    key: 'size',
    scopedSlots: { customRender: 'size' }
  },
  {
    title: 'hash',
    key: 'hash',
    dataIndex: 'hash',
    scopedSlots: { customRender: 'hash' }
  },
  {
    title: '上传人',
    key: 'created_by',
    dataIndex: 'created_by',
    scopedSlots: { customRender: 'created_by' }
  },
  {
    title: '上传时间',
    key: 'created_at',
    dataIndex: 'created_at',
    scopedSlots: { customRender: 'created_at' }
  },
  {
    title: '操作',
    key: 'action',
    dataIndex: 'action',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  data () {
    return {
      data: [],
      columns,
      loading: true
    }
  },
  filters: {
    fromNow (date) {
      return moment(date).fromNow()
    }
  },
  mounted () {
    this.getList()
  },
  methods: {
    ...mapActions(['GetVersionList']),
    handleChange (value) {
      console.log(`selected ${value}`)
    },
    getList () {
      console.log()
      const { GetVersionList } = this
      const params = {
        'app_id': this.$route.query.id
      }
      GetVersionList(params).then(res => {
        this.data = res.result.app_version
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
