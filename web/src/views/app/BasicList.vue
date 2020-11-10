<template>
  <page-header-wrapper>
    <a-card :bordered="false">
      <a-row>
        <a-col :sm="8" :xs="24">
          <info title="我的待办" value="8个任务" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="本周任务平均处理时间" value="32分钟" :bordered="true" />
        </a-col>
        <a-col :sm="8" :xs="24">
          <info title="本周完成任务数" value="24个" />
        </a-col>
      </a-row>
    </a-card>

    <a-card
      style="margin-top: 24px"
      :bordered="false"
      title="应用列表">

      <div slot="extra">
        团队：
        <a-select v-model="selectTeam" style="width: 120px" @change="getTeamApp" v-decorator="[ 'team', {rules: []}]">
          <a-select-option v-for="item in teamData" :key="item.id" >
            {{ item.name }}
          </a-select-option>
        </a-select>

        <a-button type="primary" style="margin: 0 5px 0 5px" v-if="role == 3" @click="dissolve()">
          上传
        </a-button>
      </div>

      <div class="operate">
        <a-button type="dashed" style="width: 100%" icon="plus">
          <router-link :to="{ name: 'AppUpload', query:{ teamId: selectTeam, teamName: selectTeamName } }"> <a>上传APP</a> </router-link>
        </a-button>
      </div>

      <a-list size="large" :pagination="{showSizeChanger: true, showQuickJumper: true, pageSize: 5, total: 50}">
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.name" style="flex: 0.25">
            <a-avatar slot="avatar" size="large" shape="square" :src="item.icon"/>
            <a slot="title"><router-link :to="{ name: 'AppInfo', query:{ id: item.id } }"> {{ item.name }} </router-link></a>
          </a-list-item-meta>
          <div slot="actions">
            <a @click="edit(item)">编辑</a>
          </div>
          <div slot="actions">
            <a-dropdown>
              <a-menu slot="overlay">
                <a-menu-item><a>编辑</a></a-menu-item>
                <a-menu-item><a>删除</a></a-menu-item>
              </a-menu>
              <a>更多<a-icon type="down"/></a>
            </a-dropdown>
          </div>
          <div class="list-content">
            <div class="list-content-item">
              <span>BundleID</span>
              <p>{{ item.bundle_id }}</p>
            </div>
            <div class="list-content-item">
              <span>平台</span>
              <p>{{ item.platform }}</p>
            </div>
            <div class="list-content-item">
              <span>下载次数</span>
              <p>{{ 0 }}</p>
            </div>
            <div class="list-content-item">
              <span>浏览次数</span>
              <p>{{ 0 }}</p>
            </div>
            <div class="list-content-item">
              <span>identifier</span>
              <p>{{ item.identifier ? item.identifier : ' 1' }}</p>
            </div>
            <div class="list-content-item">
              <span>短连接</span>
              <p>{{ item.short_url ? item.short_url : ' 1' }}</p>
            </div>
            <div class="list-content-item">
              <span>当前版本</span>
              <p>{{ item.current_version }}</p>
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
    ...mapActions(['TeamIndex', 'GetList']),
    add () {
      this.$dialog(TaskForm,
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
          title: '新增',
          width: 700,
          centered: true,
          maskClosable: false
        })
    },
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

          for (const key in this.teamData) {
            if (this.teamData[key].id === this.selectTeam) {
              this.selectTeamName = this.teamData[key].name
            }
          }

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
          console.log('data', this.data)
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
