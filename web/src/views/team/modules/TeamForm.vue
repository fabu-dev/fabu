<template>
  <a-form @submit="handleSubmit" :form="form">
    <a-form-item
      label="团队名称"
      :labelCol="labelCol"
      :wrapperCol="wrapperCol"
    >
      <a-input v-decorator="['name', {rules:[{required: true, message: '请输入团队名称'}]}]" />
    </a-form-item>
  </a-form>
</template>

<script>
import pick from 'lodash.pick'
import { mapActions } from 'vuex'

const fields = ['name']

export default {
  name: 'TeamForm',
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
    this.record && this.form.setFieldsValue(pick(this.record, fields))
  },
  methods: {
    ...mapActions(['TeamCreate']),
    onOk () {
      console.log('监听了 modal ok 事件')
      const { form: { validateFields }, TeamCreate } = this
      this.visible = true
      validateFields((errors, values) => {
        if (!errors) {
          console.log('values', values)
          TeamCreate(values)
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
