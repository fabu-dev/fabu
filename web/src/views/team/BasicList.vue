<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="8" :xs="24">
          <info title="我的团队" :value="this.CountTeam" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="我的APP" :value="this.CountApp" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="APP本周下载" :value="this.CountAppDownload" />
        </a-col>
      </a-row>
    </a-card>

    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="标准列表">
      <div>
      </div>
      <div slot="extra">
        团队：
        <a-select v-model="selectTeam" style="width: 120px" @change="getTeamMember" v-decorator="[ 'team', {rules: []}]">
          <a-select-option v-for="item in teamData" :key="item.id" >
            {{ item.name }}
          </a-select-option>
        </a-select>

        <a-button type="primary" style="margin: 0 5px 0 5px" v-if="role == 3" @click="dissolve()">
          解散
        </a-button>
        <a-button type="primary" style="margin: 0 5px 0 5px" v-if="role > 1" @click="edit()">
          编辑团队名称
        </a-button>
        <a-button type="primary" style="margin: 0 5px 0 5px" v-if="role > 1" @click="invite()">
          邀请团队成员
        </a-button>
      </div>

      <div class="operate">
        <a-button type="dashed" style="width: 100%" icon="plus" @click="add">创建团队</a-button>
      </div>

      <a-list size="large" :pagination="{showSizeChanger: true, showQuickJumper: true, pageSize: 50, total: 50}">
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.member_account">
            <a-avatar slot="avatar" shape="square" size="large" :style="{ backgroundColor: color, verticalAlign: 'middle' }">
              {{ item.member_name }}
            </a-avatar>
            <a slot="title">{{ item.member_name }}</a>
          </a-list-item-meta>
          <div slot="actions">
            <a @click="exit(item)">退出</a>
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
import InviteForm from './modules/InviteForm'
import Info from './components/Info'
import { mapActions } from 'vuex'

const colorList = ['#90D9FF', '#f56a00', '#7265e6', '#ffbf00', '#00a2ae']

export default {
  name: 'StandardList',
  components: {
    TeamForm,
    Info
  },
  data () {
    return {
      color: colorList[0],
      role: '',
      CountTeam: '',
      CountApp: '',
      CountAppDownload: '',
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
    ...mapActions(['TeamIndex', 'TeamMember', 'TeamMemberExit', 'TeamDelete']),
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
    edit () {
      let name = ''
      for (const key in this.teamData) {
        if (this.teamData[key].id === this.selectTeam) {
          name = this.teamData[key].name
        }
      }
      const record = {
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
          title: '编辑团队',
          width: 700,
          centered: true,
          maskClosable: false
        })
    },
    invite () {
      let name = ''
      for (const key in this.teamData) {
        if (this.teamData[key].id === this.selectTeam) {
          name = this.teamData[key].name
        }
      }
      const record = {
        'id': this.selectTeam,
        'name': name
      }
      this.$dialog(InviteForm,
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
          title: '邀请团队成员',
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
    dissolve () {
      const record = {
        'id': this.selectTeam
      }
      const { TeamDelete } = this
      this.$confirm({
        title: '确定要解散团队么?',
        content: '退出团队后，您将不能在维护APP。有维护app的团队不可以解散！',
        onOk () {
          return TeamDelete(record).then(res => {
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
        // TODO 要有全局的统一处理异常code。 否则这里在请求异常时， length会变成字符串长度
        if (res.result['team'].length > 0) {
          this.selectTeam = res.result['team'][0].id
          this.teamData = res.result['team']
          this.CountApp = res.result['count_app'] + '个'
          this.CountTeam = res.result['count_team'] + '个'
          this.CountAppDownload = res.result['count_app_download'] + '次'

          this.getTeamMember(this.selectTeam)
        }
      }).catch((err) => {
        console.log('team list', err)
      })
    },
    getTeamMember (id) {
      if (id) {
        const { TeamMember } = this
        TeamMember(id).then(res => {
          this.data = res.result.member_list
          for (const key in this.data) {
            if (this.data[key].avatar === '') {
              this.data[key].avatar = '/public/man.png'
            }
          }
          this.role = res.result.role
        }).catch((err) => {
          console.log('team list', err)
        })
      }
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
