// eslint-disable-next-line
import { UserLayout, BasicLayout, BlankLayout } from '@/layouts'
import { bxAnaalyse } from '@/core/icons'

const RouteView = {
  name: 'RouteView',
  render: (h) => h('router-view')
}

export const asyncRouterMap = [

  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/app',
    children: [
      // dashboard
      {
        path: '/dashboard',
        name: 'dashboard',
        redirect: '/dashboard/workplace',
        component: RouteView,
        meta: { title: '工作台', keepAlive: true, icon: bxAnaalyse, permission: [ 'dashboard' ] },
        hidden: true,
        children: [
          {
            path: '/dashboard/analysis/:pageNo([1-9]\\d*)?',
            name: 'Analysis',
            component: () => import('@/views/dashboard/Analysis'),
            meta: { title: 'menu.dashboard.analysis', keepAlive: false, permission: [ 'dashboard' ] }
          },
          {
            path: '/dashboard/workplace',
            name: 'Workplace',
            component: () => import('@/views/dashboard/Workplace'),
            meta: { title: 'menu.dashboard.workplace', keepAlive: true, permission: [ 'dashboard' ] }
          }
        ]
      },
      // list
      {
        path: '/app',
        name: 'app',
        component: RouteView,
        meta: { title: '应用管理', icon: 'table', keepAlive: true, permission: [ 'table' ] },
        redirect: '/app/index',
        hideChildrenInMenu: true, // 强制显示 MenuItem 而不是 SubMenu
        children: [
          {
            path: '/app/index',
            name: 'AppIndex',
            hidden: true,
            component: () => import('@/views/app/BasicList'),
            meta: { title: '应用列表', keepAlive: true, hidden: true, permission: [ 'table' ] }
          },
          {
            path: '/app/upload',
            name: 'AppUpload',
            hidden: true,
            component: () => import('@/views/app/uploadForm/StepForm'),
            meta: { title: '上传APP', keepAlive: true, hidden: true, permission: [ 'table' ] }
          },
          {
            path: '/app/info',
            name: 'AppInfo',
            hidden: true,
            component: () => import('@/views/app/search/SearchLayout'),
            redirect: '/app/search/detail',
            meta: { title: '应用详情', keepAlive: true, hidden: true, permission: [ 'table' ] },
            hideChildrenInMenu: true,
            children: [
              {
                path: '/app/search/detail',
                name: 'AppDetail',
                component: () => import('../views/app/search/Detail'),
                meta: { title: '应用概述', permission: [ 'table' ] }
              },
              {
                path: '/app/search/version',
                name: 'AppVersion',
                component: () => import('../views/app/search/Version'),
                meta: { title: '版本列表', permission: [ 'table' ] }
              },
              {
                path: '/app/search/team',
                name: 'AppTeam',
                component: () => import('../views/app/search/Team'),
                meta: { title: '团队', permission: [ 'table' ] }
              }
            ]
          }
        ]
      },
      // profile
      {
        path: '/team',
        name: 'team',
        component: () => import('@/views/team/BasicList'),
        meta: { title: '我的团队', icon: 'profile', permission: [ 'team' ] }
      },

      // account
      {
        path: '/account',
        component: () => import('@/views/account/settings/Index'),
        redirect: '/account/settings/base',
        name: 'account',
        meta: { title: '个人设置', icon: 'user', keepAlive: true, permission: [ 'user' ] },
        hideChildrenInMenu: true,
        hidden: true,
        children: [
          {
            path: '/account/settings/base',
            name: 'BaseSettings',
            component: () => import('@/views/account/settings/BaseSetting'),
            meta: { title: '基本设置', hidden: true, permission: [ 'user' ] }
          },
          {
            path: '/account/settings/security',
            name: 'SecuritySettings',
            component: () => import('@/views/account/settings/Security'),
            meta: { title: '安全设置', hidden: true, keepAlive: true, permission: [ 'user' ] }
          },
          {
            path: '/account/settings/custom',
            name: 'CustomSettings',
            component: () => import('@/views/account/settings/Custom'),
            meta: { title: '个性化设置', hidden: true, keepAlive: true, permission: [ 'user' ] }
          },
          {
            path: '/account/settings/binding',
            name: 'BindingSettings',
            component: () => import('@/views/account/settings/Binding'),
            meta: { title: '账户绑定', hidden: true, keepAlive: true, permission: [ 'user' ] }
          },
          {
            path: '/account/settings/notification',
            name: 'NotificationSettings',
            component: () => import('@/views/account/settings/Notification'),
            meta: { title: '新消息通知', hidden: true, keepAlive: true, permission: [ 'user' ] }
          }
        ]
      },

      // other
      {
        path: '/other',
        name: 'otherPage',
        component: RouteView,
        meta: { title: '其他组件', icon: 'slack', permission: [ 'dashboard' ] },
        redirect: '/other/icon-selector',
        hidden: true,
        children: [
          {
            path: '/other/icon-selector',
            name: 'TestIconSelect',
            component: () => import('@/views/other/IconSelectorView'),
            meta: { title: '图标', icon: 'tool', keepAlive: true, permission: [ 'dashboard' ] }
          },
          {
            path: '/other/list/tree-list',
            name: 'TreeList',
            component: () => import('@/views/other/TreeList'),
            meta: { title: '树目录表格', icon: 'tool', keepAlive: true, permission: [ 'dashboard' ] }
          },
          {
            path: '/other/list/edit-table',
            name: 'EditList',
            component: () => import('@/views/other/TableInnerEditList'),
            meta: { title: '内联编辑表格', keepAlive: true }
          },
          {
            path: '/other/list/user-list',
            name: 'UserList',
            component: () => import('@/views/other/UserList'),
            meta: { title: '用户列表', keepAlive: true }
          },
          {
            path: '/other/list/role-list',
            name: 'RoleList',
            component: () => import('@/views/other/RoleList'),
            meta: { title: '角色列表', keepAlive: true }
          },
          {
            path: '/other/list/system-role',
            name: 'SystemRole',
            component: () => import('@/views/role/RoleList'),
            meta: { title: '角色列表2', keepAlive: true }
          },
          {
            path: '/other/list/permission-list',
            name: 'PermissionList',
            component: () => import('@/views/other/PermissionList'),
            meta: { title: '权限列表', keepAlive: true }
          }
        ]
      }
    ]
  },
  {
    path: '/s/:short_url',
    name: 'Preview',
    hidden: true,
    component: () => import('@/views/app/Preview'),
    meta: { title: '安装APP', keepAlive: true, hidden: true }
  },
  {
    path: '*', redirect: '/404', hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Register')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/RegisterResult')
      },
      {
        path: 'recover',
        name: 'recover',
        component: undefined
      }
    ]
  },

  {
    path: '/404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  }

]
