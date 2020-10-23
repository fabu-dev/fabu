<template>
  <a-form @submit="handleSubmit" :form="form">
    <a-form-item label="ID" v-if="record.id" :labelCol="labelCol" :wrapperCol="wrapperCol">
      <a-input v-decorator="['id', {initialValue: record.id, rules:[]}]" :allowClear="true" :disabled="true" />
    </a-form-item>
    <a-form-item
      label="团队名称"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-input v-decorator="['name', {rules:[{required: true, message: '请输入团队名称'}]}]" :allowClear="true" :disabled="true"/>
    </a-form-item>
    <a-form-item
      label="邮箱"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-input v-decorator="['email', {rules: [{ required: true, type: 'email', message: '请输入邮箱地址' }], validateTrigger: ['change', 'blur']}]"/>
    </a-form-item>
    <a-form-item
      label="角色"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-select style="width: 120px" v-decorator="['role', {initialValue: 1, rules:[]}]">
        <a-select-option :value="1">
          团队成员
        </a-select-option>
        <a-select-option :value="2">
          管理员
        </a-select-option>
      </a-select>
    </a-form-item>
  </a-form>
</template>

<script>
import pick from 'lodash.pick'
import { mapActions } from 'vuex'

const fields = ['name']

export default {
  name: 'InviteForm',
  props: {
    record: {
      type: Object,
      default: null
    }
  },
  data () {
    return {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      },
      form: this.$form.createForm(this)
    }
  },
  mounted () {
    console.log('record', this.record)
    this.record && this.form.setFieldsValue(pick(this.record, fields))
  },
  methods: {
    ...mapActions(['TeamMemberAdd']),
    onOk () {
      console.log('监听了 modal ok 事件')
      const { form: { validateFields }, TeamMemberAdd } = this
      this.visible = true
      validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          TeamMemberAdd(values)
            .then((res) => this.success(res))
            .catch(err => this.failed(err))
        } else {
          this.visible = false
        }
      })
      if (this.visible) {
        return new Promise(resolve => {
          resolve(true)
        })
      }
    },
    getMemberByEmail () {
      const { form: { validateFields }, TeamCreate, TeamEdit } = this
      this.visible = true
      validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          if (this.record.id > 0) {
            TeamEdit(values)
              .then((res) => this.success(res))
              .catch(err => this.failed(err))
          } else {
            TeamCreate(values)
              .then((res) => this.success(res))
              .catch(err => this.failed(err))
          }
        } else {
          this.visible = false
        }
      })
    },
    onCancel () {
      console.log('监听了 modal cancel 事件')
      return new Promise(resolve => {
        resolve(true)
      })
    },
    handleSubmit (e) {
      console.log('监听了 submit 事件')
      const { form: { validateFields } } = this
      this.visible = true
      validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
        }
      })
    },
    success (res) {
      console.log(res)
    },
    failed (err) {
      console.log(err)
    }
  }
}
</script>
