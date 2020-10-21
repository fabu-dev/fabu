<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="8" :xs="24">
          <info title="我的团队" value="8个团队" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="本周团队平均处理时间" value="32分钟" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="本周完成团队数" value="24个" />
        </a-col>
      </a-row>
    </a-card>

    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="标准列表">

      <div slot="extra">
        团队：
        <a-select v-model="selectTeam" style="width: 120px" @change="getTeamMember" v-decorator="[ 'team', {rules: []} ]">
          <a-select-option v-for="item in teamData" :key="item.id" >
            {{ item.name }}
          </a-select-option>
        </a-select>
      </div>

      <div class="operate">
        <a-button type="dashed" style="width: 100%" icon="plus" @click="add">创建团队</a-button>
      </div>

      <a-list size="large" :pagination="{showSizeChanger: true, showQuickJumper: true, pageSize: 5, total: 50}">
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.member_account">
            <a-avatar slot="avatar" size="large" shape="square" :src="item.avatar"/>
            <a slot="title">{{ item.member_name }}</a>
          </a-list-item-meta>
          <div slot="actions">
            <a @click="exit(item)">退出</a>
          </div>
          <div slot="actions">
            <a-dropdown>
              <a-menu slot="overlay">
                <a-menu-item v-if="item.role == 3"><a>解散</a></a-menu-item>
                <a-menu-item v-if="item.role == 2 || item.role == 3" @click="edit(item)"><a>编辑团队名称</a></a-menu-item>
                <a-menu-item v-if="item.role == 2 || item.role == 3"><a>邀请团队成员</a></a-menu-item>
              </a-menu>
              <a>更多<a-icon type="down"/></a>
            </a-dropdown>
          </div>
          <div class="list-content">
            <div class="list-content-item">
              <span>角色</span>
              <p>{{ item.role_name }}</p>
            </div>
          </div>
          <div class="list-content">
            <div class="list-content-item">
              <span>加入时间</span>
              <p>{{ item.created_at }}</p>
            </div>
          </div>
        </a-list-item>
      </a-list>
    </a-card>
  </page-header-wrapper>
</template>

<script>
// 演示如何使用 this.$dialog 封装 modal 组件
import TeamForm from './modules/TeamForm'
import Info from './components/Info'
import { mapActions } from 'vuex'

export default {
  name: 'StandardList',
  components: {
    TeamForm,
    Info
  },
  data () {
    return {
      teamData: [],
      selectTeam: '', // 靠，这里不能给0
      data: [],
      status: 'all'
    }
  },
  created () {
    this.getTeamData()
  },
  mounted () {

  },
  methods: {
    ...mapActions(['TeamIndex', 'TeamMember', 'TeamMemberExit']),
    add () {
      this.$dialog(TeamForm,
        {
          record: {},
          on: {
            ok () {
              console.log('ok 回调')
            },
            cancel () {
              console.log('cancel 回调')
            },
            close () {
              console.log('modal close 回调')
            }
          }
        },
        // modal props
        {
          title: '创建团队',
          width: 700,
          centered: true,
          maskClosable: false
        })
    },
    edit (record) {
      let name = ''
      for (const key in this.teamData) {
        if (this.teamData[key].id === this.selectTeam) {
          name = this.teamData[key].name
        }
      }
      record = {
        'id': this.selectTeam,
        'name': name
      }

      this.$dialog(TeamForm,
        // component props
        {
          record,
          on: {
            ok () {
              console.log('ok 回调')
            },
            cancel () {
              console.log('cancel 回调')
            },
            close () {
              console.log('modal close 回调')
            }
          }
        },
        // modal props
        {
          title: '操作',
          width: 700,
          centered: true,
          maskClosable: false
        })
    },
    exit (record) {
      record = {
        'id': record.id
      }
      const { TeamMemberExit } = this
      this.$confirm({
        title: '确定要退出团队么?',
        content: '退出团队后，您将不能在维护APP。',
        onOk () {
          return TeamMemberExit(record).then(res => {
            if (res.result.length > 0) {
              this.selectTeam = res.result[0].id
              this.teamData = res.result

              this.getTeamMember(this.selectTeam)
            }
          }).catch((err) => {
            console.log('team list', err)
          })
        },
        onCancel () {}
      })
    },
    getTeamData () {
      const { TeamIndex } = this
      TeamIndex().then(res => {
        if (res.result.length > 0) {
          this.selectTeam = res.result[0].id
          this.teamData = res.result

          this.getTeamMember(this.selectTeam)
        }
      }).catch((err) => {
        console.log('team list', err)
      })
    },
    getTeamMember (id) {
      const { TeamMember } = this
      TeamMember(id).then(res => {
        this.data = res.result
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
