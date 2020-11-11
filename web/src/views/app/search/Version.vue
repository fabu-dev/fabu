<template>
  <div>
    <a-table :columns="columns" :data-source="data">
      <a slot="name" slot-scope="text">{{ text }}</a>
      <span slot="customTitle"><a-icon type="smile-o" /> Name</span>
      <span slot="tags" slot-scope="tags">
        <a-tag
          v-for="tag in tags"
          :key="tag"
          :color="tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'"
        >
          {{ tag.toUpperCase() }}
        </a-tag>
      </span>
      <span slot="action" slot-scope="text, record">
        <a>Invite 一 {{ record.name }}</a>
        <a-divider type="vertical" />
        <a>Delete</a>
        <a-divider type="vertical" />
        <a class="ant-dropdown-link"> More actions <a-icon type="down" /> </a>
      </span>
    </a-table>
  </div>
</template>

<script>
import moment from 'moment'
import { mapActions } from 'vuex'

const columns = [
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
    title: '大小',
    dataIndex: 'size',
    key: 'size'
  },
  {
    title: 'hash',
    key: 'hash',
    dataIndex: 'hash',
    scopedSlots: { customRender: 'hash' }
  },
  {
    title: '操作',
    key: 'action',
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
        this.data = res.result
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
