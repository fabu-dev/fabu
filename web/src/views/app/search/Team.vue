<template>
  <div>
    <a-card :bordered="false" title="团队信息">
      <a-descriptions :column="4">
        <a-descriptions-item label="ID">{{ data.id }}</a-descriptions-item>
        <a-descriptions-item label="团队名称">{{ data.name }}</a-descriptions-item>
        <a-descriptions-item label="创建人">{{ data.created_by }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ data.created_at }}</a-descriptions-item>
      </a-descriptions>
    </a-card>

    <a-card style="margin-top: 24px" :bordered="false" title="团队成员">
      <a-table
        style="margin-bottom: 24px"
        row-key="id"
        :columns="columns"
        :data-source="memberData"
        :pagination="false"
        bordered
      >
      </a-table>
    </a-card>
  </div>
</template>

<script>
import moment from 'moment'
import { mapActions } from 'vuex'

const colorList = ['#90D9FF', '#f56a00', '#7265e6', '#ffbf00', '#00a2ae']

const columns = [
  {
    title: 'ID',
    dataIndex: 'member_id',
    key: 'member_id'
  },
  {
    title: '姓名',
    dataIndex: 'member_name',
    key: 'member_name'
  },
  {
    title: '账户',
    dataIndex: 'member_account',
    key: 'member_account',
    scopedSlots: { customRender: 'member_account' }
  },
  {
    title: '邮箱',
    dataIndex: 'member_email',
    key: 'member_email'
  },
  {
    title: '角色',
    dataIndex: 'role_name',
    key: 'role_name'
  },
  {
    title: '加入时间',
    dataIndex: 'created_at',
    key: 'created_at'
  }
]

export default {
  components: {
  },
  data () {
    return {
      columns,
      color: colorList[0],
      data: {},
      memberData: [],
      form: this.$form.createForm(this),
      loading: true
    }
  },
  filters: {
    fromNow (date) {
      return moment(date).fromNow()
    }
  },
  mounted () {
    this.getInfo(this.$route.query.team_id)
    this.getTeamMember(this.$route.query.team_id)
  },
  methods: {
    ...mapActions(['TeamInfo', 'TeamMember']),
    handleChange (value) {
      console.log(`selected ${value}`)
    },
    getInfo (id) {
      if (id) {
        const { TeamInfo } = this
        TeamInfo(id).then(res => {
          this.data = res.result
        }).catch((err) => {
          console.log('team list', err)
        })
      }
    },
    getTeamMember (id) {
      if (id) {
        const { TeamMember } = this
        TeamMember(id).then(res => {
          this.memberData = res.result.member_list
        }).catch((err) => {
          console.log('team list', err)
        })
      }
    }
  }
}
</script>

<style lang="less" scoped>
.ant-pro-components-tag-select {
  /deep/ .ant-pro-tag-select .ant-tag {
    margin-right: 24px;
    padding: 0 8px;
    font-size: 14px;
  }
}
.ant-pro-pages-list-projects-cardList {
  margin-top: 24px;

  /deep/ .ant-card-meta-title {
    margin-bottom: 4px;
  }

  /deep/ .ant-card-meta-description {
    height: 44px;
    overflow: hidden;
    line-height: 22px;
  }

  .cardItemContent {
    display: flex;
    height: 20px;
    margin-top: 16px;
    margin-bottom: -4px;
    line-height: 20px;

    > span {
      flex: 1 1;
      color: rgba(0,0,0,.45);
      font-size: 12px;
    }

    /deep/ .ant-pro-avatar-list {
      flex: 0 1 auto;
    }
  }
}
</style>
