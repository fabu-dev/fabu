<template>
  <div>
    <a-form :form="form" style="max-width: 800px; margin: 40px auto 0;">
      <a-form-item
        label=""
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-upload-dragger
          name="file"
          :action="uploadUrl"
          :multiple="false"
          :defaultFileList="defaultFileList"
          :customRequest="chunkUpload"
          @change="handleChange"
        >
          <p class="ant-upload-drag-icon">
            <a-icon type="inbox" />
          </p>
          <p class="ant-upload-text">
            点击或者拖拽APP文件开始上传
          </p>
          <p class="ant-upload-hint">
            一次仅支持上传一个文件，支持IPA、APK文件类型。
          </p>
        </a-upload-dragger>
      </a-form-item>
      <a-form-item :wrapperCol="{span: 19, offset: 5}">
        <a-button type="primary" @click="nextStep">下一步</a-button>
      </a-form-item>
    </a-form>
    <a-divider />
    <div class="step-form-style-desc">
      <h3>说明</h3>
      <h4>转账到支付宝账户</h4>
      <p>如果需要，这里可以放一些关于产品的常见问题说明。如果需要，这里可以放一些关于产品的常见问题说明。如果需要，这里可以放一些关于产品的常见问题说明。</p>
      <h4>转账到银行卡</h4>
      <p>如果需要，这里可以放一些关于产品的常见问题说明。如果需要，这里可以放一些关于产品的常见问题说明。如果需要，这里可以放一些关于产品的常见问题说明。</p>
    </div>
  </div>
</template>

<script>

import { mapActions } from 'vuex'
import chunkUpload from '@/utils/chunkupload'
// import { appApi } from '@/api/app'

export default {
  name: 'Step1',
  data () {
    return {
      defaultFileList: [],
      uploadUrl: process.env.VUE_APP_API_BASE_URL + '/app/upload',
      labelCol: { lg: { span: 5 }, sm: { span: 5 } },
      wrapperCol: { lg: { span: 19 }, sm: { span: 19 } },
      chunkUpload: chunkUpload, // 分片上传自定义方法，在头部引入了
      form: this.$form.createForm(this)
    }
  },
  created () {
  },
  methods: {
    ...mapActions(['UploadApp']),
    handleChange (info) {
      const status = info.file.status
      if (status !== 'uploading') {
        console.log(info.file, info.fileList)
      }
      if (status === 'done') {
        this.$message.success(`${info.file.name} file uploaded successfully.`)
      } else if (status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    beforeUpload (file) {
      const isJpgOrPng = file.type === 'application/vnd.android.package-archive' || file.type === ''
      // if (!isJpgOrPng) {
      //   this.$message.error('文件不是APK或IPA')
      // }
      const isLt2M = file.size / 1024 / 1024 / 1024 < 1
      if (!isLt2M) {
        this.$message.error('上传文件能能超过1GB!')
      }
      return isJpgOrPng && isLt2M
    },
    upload (data) {
      const { UploadApp } = this
      const params = new FormData()

      console.log('upload data', data)

      params.append('file', data.file)
      this.uploading = true
      console.log('upload params data', params)
      UploadApp(params).then(res => {
        this.uploading = false
        console.log('upload res', res)
      }).catch((err) => {
        this.uploading = false
        console.log('team list', err)
      })
    },
    onError () {
      this.$alert('文件上传失败，请重试', '错误', {
        confirmButtonText: '确定'
      })
    },
    nextStep () {
    const { form: { validateFields } } = this
    // 先校验，通过表单校验后，才进入下一步
    validateFields((err, values) => {
      if (!err) {
        this.$emit('nextStep')
      }
    })
  }
  }
}
</script>

<style lang="less" scoped>
.step-form-style-desc {
  padding: 0 56px;
  color: rgba(0,0,0,.45);

  h3 {
    margin: 0 0 12px;
    color: rgba(0,0,0,.45);
    font-size: 16px;
    line-height: 32px;
  }

  h4 {
    margin: 0 0 4px;
    color: rgba(0,0,0,.45);
    font-size: 14px;
    line-height: 22px;
  }

  p {
    margin-top: 0;
    margin-bottom: 12px;
    line-height: 22px;
  }
}
</style>
