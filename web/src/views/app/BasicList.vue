<template>
  <page-header-wrapper>
<!--    <a-card :bordered="false">-->
<!--      <a-row>-->
<!--        <a-col :sm="8" :xs="24">-->
<!--          <info title="我的待办" value="8个任务" :bordered="true" />-->
<!--        </a-col>-->
<!--        <a-col :sm="8" :xs="24">-->
<!--          <info title="本周任务平均处理时间" value="32分钟" :bordered="true" />-->
<!--        </a-col>-->
<!--        <a-col :sm="8" :xs="24">-->
<!--          <info title="本周完成任务数" value="24个" />-->
<!--        </a-col>-->
<!--      </a-row>-->
<!--    </a-card>-->

    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="应用列表">
      <div slot="extra">
        团队：
        <a-select v-model="selectTeam" style="width: 120px;margin-right: 16px;" @change="getTeamApp" v-decorator="[ 'team', {rules: []}]">
          <a-select-option v-for="item in teamData" :key="item.id" >
            {{ item.name }}
          </a-select-option>
        </a-select>

        <a-button type="primary" ><router-link :to="{ name: 'AppUpload', query:{ teamId: this.selectTeam} }">上传APP</router-link></a-button>
      </div>

      <div class="operate">
      </div>

      <a-list size="large">
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.env | getEnvName" style="flex: 0.25">
            <a-avatar slot="avatar" size="large" shape="square" :src="item.icon"/>{{ item.icon }}
            <a slot="title"><router-link :to="{ name: 'AppInfo', query:{ id: item.id, team_id: item.team_id } }"> {{ item.name }} </router-link></a>
          </a-list-item-meta>
          <div slot="actions">
            <router-link target="_blank" :to="{ name: 'Preview', params:{ short_url: item.short_url } }"><a>预览</a></router-link>
          </div>
          <div slot="actions">
            <a-dropdown>
              <a-menu slot="overlay">
                <a-menu-item><a @click="edit(item)">编辑</a></a-menu-item>
                <a-menu-item><a @click="del(item.id)">删除</a></a-menu-item>
              </a-menu>
              <a>更多<a-icon type="down"/></a>
            </a-dropdown>
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
              <p>{{ item.current_version }}</p>
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
      role: '',
      data: [],
      teamData: [],
      selectTeam: '', // 靠，这里不能给0
      selectTeamName: '',
      CountTeam: '',
      CountApp: '',
      CountAppDownload: ''
    }
  },
  created () {
    this.getTeamData()
  },
  methods: {
    ...mapActions(['TeamIndex', 'GetList', 'DeleteApp']),
    edit (record) {
      console.log('record', record)
      this.$dialog(TaskForm,
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

          this.getTeamApp(this.selectTeam)
        }
      }).catch((err) => {
        console.log('team list', err)
      })
    },
    getTeamApp (teamId) {
      if (teamId) {
        const { GetList } = this
        const params = {
          'team_id': teamId
        }
        GetList(params).then(res => {
          this.data = res.result.app
          for (const key in this.data) {
            this.data[key].icon = process.env.VUE_APP_API_BASE_URL + '/' + this.data[key].icon
          }

          console.log('list data', this.data)
        }).catch((err) => {
          console.log('team list', err)
        })
      }
    },
    del (id) {
      const { DeleteApp } = this
      const params = {
        'id': id
      }
      this.$confirm({
        title: '确定要删除该应用么?',
        content: '',
        onOk () {
          return DeleteApp(params).then(res => {
            if (res.result.length > 0) {
              this.getTeamApp(this.selectTeam)
            }
          }).catch((err) => {
            console.log('delete app err:', err)
          })
        },
        onCancel () {}
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
