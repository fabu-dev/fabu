<template>
  <div>
    <a-form :form="form" style="max-width: 500px; margin: 40px auto 0;">
      <a-alert
        :closable="true"
        message="确认转账后，资金将直接打入对方账户，无法退回。"
        style="margin-bottom: 24px;"
      />
      <a-form-item
        label="APP名称"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        {{ sendData.name }}
      </a-form-item>
      <a-form-item
        label="团队"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        {{ teamData.name + ' ( ' + sendData.team_id + ' ) ' }}
      </a-form-item>
      <a-form-item
        label="BundleID"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        {{ sendData.bundle_id }}
      </a-form-item>
      <a-form-item
        label="版本"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        {{ sendData.version }}
      </a-form-item>
      <a-form-item
        label="ICON"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        {{ sendData.icon }}
      </a-form-item>
      <a-form-item
        label="大小"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        {{ (sendData.size/1024/1024).toFixed(2) }} m
        <a-input type="hidden" v-decorator="['team_id', { initialValue: sendData.team_id, rules: [{required: true, message: '参数错误'}] }]"/>
        <a-input type="hidden" v-decorator="['identifier', { initialValue: sendData.identifier, rules: [{required: true, message: '参数错误'}] }]"/>
      </a-form-item>
      <a-form-item
        label="更新说明"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
        class="stepFormText"
      >
        <a-textarea placeholder="请填写更新说明" allow-clear :rows="4" v-decorator="['description', { initialValue: '', rules: [{required: true, message: '请填写更新说明'}] }]"/>
      </a-form-item>
      <a-form-item :wrapperCol="{span: 19, offset: 5}">
        <a-button :loading="loading" type="primary" @click="nextStep">提交</a-button>
        <a-button style="margin-left: 8px" @click="prevStep">上一步</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'Step2',
  props: ['sendData'], // 用来接收父组件传给子组件的数据
  data () {
    return {
      teamData: {},
      labelCol: { lg: { span: 5 }, sm: { span: 5 } },
      wrapperCol: { lg: { span: 19 }, sm: { span: 19 } },
      form: this.$form.createForm(this),
      loading: false,
      timer: 0
    }
  },
  mounted () {
    console.log('team id', this.sendData)
    this.getTeamInfo(this.sendData.team_id)
  },
  methods: {
    ...mapActions(['SaveApp', 'TeamInfo']),
    nextStep () {
      const that = this
      const { form: { validateFields }, SaveApp } = this
      that.loading = true
      validateFields((err, values) => {
        if (!err) {
          console.log('表单 values', values)
          values.team_id = Number(values.team_id)
          SaveApp(values).then((res) => {
              console.log('save res:', res)
              that.timer = setTimeout(function () {
                that.loading = false
                that.$emit('nextStep')
              }, 1500)
          }).catch(err => {
            console.log('save err:', err)
          })
        } else {
          that.loading = false
        }
      })
    },
    getTeamInfo (id) {
      if (id) {
        const { TeamInfo } = this
        TeamInfo(id).then(res => {
          this.teamData = res.result
        }).catch((err) => {
          console.log('team list', err)
        })
      }
    },
    prevStep () {
      this.$emit('prevStep')
    }
  },
  beforeDestroy () {
    clearTimeout(this.timer)
  }
}
</script>

<style lang="less" scoped>
  .stepFormText {
    margin-bottom: 24px;

    .ant-form-item-label,
    .ant-form-item-control {
      line-height: 22px;
    }
  }

</style>
