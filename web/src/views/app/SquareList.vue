<template>
  <page-header-wrapper>
    <section class="app-container" style="">
      <div class="page-apps">
        <div class="middle-wrapper">
        </div>
        <div class="middle-wrapper container-fluid">
          <div class="apps row">
            <div :key="index" v-for="(item, index) in data" class="col-xs-4 col-sm-4 col-md-4 app-animator">
              <div :class="'card app card-' + item.platform_name.toLowerCase()">
                <i :class="'type-icon icon-' + item.platform_name.toLowerCase()"></i>
                <div class="type-mark"></div>
                <a class="appicon">
                  <router-link target="_blank" :to="{ name: 'Preview', params:{ short_url: item.short_url } }"> <img class="icon" width="100" height="100" :src="item.icon" /> </router-link>
                </a>
                <br>
                <p class="appname">
                  <i class="icon-owner"></i>
                  <span><router-link target="_blank" :to="{ name: 'Preview', params:{ short_url: item.short_url } }"> {{ item.name }} </router-link></span> <span style="font-size: 14px;">{{ item.env | getEnvName }}</span>
                </p>
                <table>
                  <tbody>
                    <tr>
                      <td>包名:</td>
                      <td>
                        <span>{{ item.bundle_id }}</span>
                      </td>
                    </tr>
                    <tr>
                      <td>版本:</td>
                      <td>
                        <span>{{ item.current_version }}（{{ item.current_build }}）</span>
                      </td>
                    </tr>
                    <tr>
                      <td>更新:</td>
                      <td><span>{{ item.updated_at | formatDate }}</span></td>
                    </tr>
                  </tbody>
                </table>
                <div class="action">
                  <router-link target="_blank" :to="{ name: 'Preview', params:{ short_url: item.short_url } }">
                    <i class="icon-eye"></i> 预览
                  </router-link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
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
      data: []
    }
  },
  created () {
    this.getApp()
  },
  methods: {
    ...mapActions(['GetSquare']),
    getApp () {
      const { GetSquare } = this
      GetSquare().then(res => {
        this.data = res.result.app
        for (const key in this.data) {
          this.data[key].icon = process.env.VUE_APP_API_BASE_URL + '/' + this.data[key].icon
        }

        console.log('list data', this.data)
      }).catch((err) => {
        console.log('team list', err)
      })
    }
  }
}
</script>

<style lang="less" scoped>
.app-container {
  width: 1024px;
  margin: 0 auto;
  color: #9B9B9B;
  -webkit-font-smoothing: antialiased;
}
.col-md-4 {
  width: 33.3333333333%;
  float: left;
  position: relative;
  min-height: 1px;
  padding-left: 15px;
  padding-right: 15px;
}
.page-apps .card.app .action a,
.page-apps .card.app .appname,
.page-apps .card.app table tr td>span {
  white-space: nowrap;
  max-width: 100%;
  word-wrap: normal;
  overflow: hidden;
  text-overflow: ellipsis
}

.page-apps {
  padding-top: 20px;
}

.page-apps .card.app .action a,
.page-apps .card.app .appname,
.page-apps .card.app table tr td {
  font-family: open sans, sans-serif
}

.page-apps .card.app .appicon img {
  border-radius: 17.54%
}

.page-apps .container-fluid {
  padding: 15px 0 0
}

.page-apps .container-fluid .col-md-4 {
  margin-bottom: 30px
}

.page-apps .card.app {
  position: relative;
  height: 460px;
  padding: 44px;
  background-color: #FFF;
  transition: all .25s
}

.page-apps .card.app:hover {
  box-shadow: 0 15px 30px rgba(0, 0, 0, .1);
  transform: translateY(-4px)
}

.page-apps .card.app .type-icon {
  position: absolute;
  top: 9px;
  right: 7px;
  color: #FFF;
  z-index: 2
}

.page-apps .card.app .type-mark {
  position: absolute;
  top: 0;
  right: 0;
  z-index: 1
}

.page-apps .card.app .appicon {
  display: inline-block;
  width: 100px;
  height: 100px;
  cursor: pointer
}

.page-apps .card.app .appname {
  display: inline-block;
  margin-top: 36px;
  font-size: 18px;
  cursor: pointer
}

.page-apps .card.app .appname i {
  margin-right: 6px;
  color: #F8BA0B;
  font-size: 16px
}

.page-apps .card.app:hover .appname {
  color: #4A4A4A
}

.page-apps .card.app table {
  width: 100%;
  table-layout: fixed
}

.page-apps .card.app table tr td {
  padding: 2px 0;
  font-size: 12px
}

.page-apps .card.app table tr td>span {
  display: inline-block
}

.page-apps .card.app table tr td:last-child {
  width: 80%;
  color: #1A1A1A
}

.page-apps .card.app .action {
  position: absolute;
  left: 0;
  bottom: 0;
  width: 100%;
  padding: 40px 0 40px 40px
}

.page-apps .card.app .action a {
  display: inline-block;
  margin-right: 4px;
  padding: 8px 20px;
  border: 1px solid;
  color: #9B9B9B;
  text-decoration: none;
  text-align: center;
  border-radius: 40px;
  transition: all .25s
}

.page-apps .card.app .action a:hover {
  color: #F8BA0B
}

.page-apps .card.app .action a i {
  display: inline-block;
  margin-right: 4px;
  font-size: 18px;
  vertical-align: middle
}

.page-apps .card.app .action a .icon-pen,
.page-apps .card.app .action .btn-remove i {
  font-size: 16px
}

.page-apps .card.app .action .btn-remove {
  padding: 10px;
  border: 1px solid;
  background: 0 0;
  font-size: 0;
  vertical-align: top;
  border-radius: 50%
}

.page-apps .card.app .action .btn-remove:hover {
  background-color: #EC4242;
  color: #FFF
}

.page-apps .card-android .type-mark {
  width: 0;
  height: 0;
  border-top: 48px solid #A4C639;
  border-left: 48px solid transparent
}

.page-apps .card-ios .type-mark {
  width: 0;
  height: 0;
  border-top: 48px solid #C6C7C9;
  border-left: 48px solid transparent
}

.app-animator {
  transition: all .25s;
  opacity: 1
}

@font-face {
  font-family: icomoon;
  font-weight: 400;
  font-style: normal;
  src: url(../../assets/icomoon.ttf?wcusdg) format('truetype')
}

[class*=" icon-"],
[class^=icon-] {
  line-height: 1;
  -webkit-font-smoothing: antialiased;
  font-family: icomoon!important;
  font-style: normal;
  font-weight: 400;
  font-variant: normal;
  -moz-osx-font-smoothing: grayscale;
  speak: none;
  text-transform: none
}

.icon-angle-right:before {
  content: '\e606'
}

.icon-logo:before {
  content: '\e620'
}

.icon-menu:before {
  content: '\e61e'
}

.icon-pen:before {
  content: '\e626'
}

.icon-eye:before {
  content: '\e62b'
}

.icon-owner:before {
  content: '\e62f'
}

.icon-upload-cloud2:before {
  content: '\e632'
}

.icon-android:before {
  content: '\e633'
}

.icon-trash:before {
  content: '\e634'
}

.icon-ios:before {
  content: '\e63f'
}

.icon-cloud-download:before {
  content: '\e607'
}

.icon-device:before {
  content: '\e622'
}

.icon-file:before {
  content: '\e624'
}

.icon-calendar:before {
  content: '\e637'
}

.icon-plus:before {
  content: "\e62c";
}

.icon-box:before {
  content: "\e620";
}

button:active,
button:focus,
button:focus:active {
  outline: 0
}

.appicon {
  position: relative;
  border-radius: 17.544%;
  overflow: hidden
}
</style>
