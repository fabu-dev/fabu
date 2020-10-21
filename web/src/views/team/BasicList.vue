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
        <a-select v-model="teamSelect" style="width: 120px" @change="getTeamMember">
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
              <span>Owner</span>
              <p>{{ item.role_name }}</p>
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
      teamSelect: '', // 靠，这里不能给0
      data: [],
      status: 'all'
    }
  },
  created () {
    this.getTeamData()
    console.log('teamSelect3', this.teamSelect)
  },
  mounted () {

  },
  methods: {
    ...mapActions(['TeamIndex', 'TeamMember']),
    add () {
      this.$dialog(TeamForm,
        // component props
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
      console.log('record', record)
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
    getTeamData () {
      const { TeamIndex } = this
      TeamIndex().then(res => {
        if (res.result.length > 0) {
          this.teamSelect = res.result[0].id
          this.teamData = res.result

          this.getTeamMember(this.teamSelect)
        }

        console.log('teamSelect2', this.teamSelect)
      }).catch((err) => {
        console.log('team list', err)
      })
    },
    getTeamMember (id) {
      console.log('get team by id :', id)
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
