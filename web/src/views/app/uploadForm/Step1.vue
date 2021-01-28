<template>
  <div>
    <a-form :form="form" style="max-width: 800px; margin: 40px auto 0;">
      <a-form-item>
        <a-upload-dragger
          name="file"
          :action="uploadUrl"
          :multiple="false"
          :fileList="fileList"
          :customRequest="chunkUpload"
          @change="handleChange"
          v-decorator="['file', { rules: [{required: true, message: '请上传文件'}] }]"
        >
          <p class="ant-upload-drag-icon">
            <a-icon type="inbox" />
          </p>
          <p class="ant-upload-text">
            点击或者拖拽APP文件开始上传
          </p>
          <p class="ant-upload-hint">
            一次仅支持上传一个文件，支持 ipa, apk 文件类型。
          </p>
        </a-upload-dragger>
      </a-form-item>
      <a-form-item>
        <a-input type="hidden" v-decorator="['team_id', { initialValue: this.$route.query.teamId, rules: [{required: true, message: '参数错误'}] }]"/>
      </a-form-item>
      <a-form-item style="text-align: center;">
        <a-button ref="nextButton" type="primary" @click="nextStep">下一步</a-button>
      </a-form-item>
    </a-form>
    <a-divider />
    <div class="step-form-style-desc">
      <h3>说明</h3>
      <p>1.只可以上传 ipa 或 apk 格式的应用文件</p>
      <p>2.文件大小不超过 1GB</p>
      <p>3.上传过程中，不要关闭或刷新页面</p>
    </div>
  </div>
</template>

<script>

import { mapActions } from 'vuex'
import chunkUpload from '@/utils/chunkupload'

export default {
  name: 'Step1',
  data () {
    return {
      fileList: [],
      uploadUrl: process.env.VUE_APP_API_BASE_URL + '/v1/app/upload',
      labelCol: { lg: { span: 5 }, sm: { span: 5 } },
      wrapperCol: { lg: { span: 19 }, sm: { span: 19 } },
      chunkUpload: chunkUpload, // 分片上传自定义方法，在头部引入了
      form: this.$form.createForm(this)
    }
  },
  created () {
    console.log('router', this.$route)
  },
  methods: {
    ...mapActions(['UploadApp', 'getBase']),
    handleChange (info) {
      let fileList = [...info.fileList]
      fileList = fileList.slice(-1)
      this.fileList = fileList

      const status = info.file.status
      if (status !== 'uploading') {
        console.log(info.file, info.fileList)
      }

      if (status === 'done') {
        this.$message.success(`${info.file.name} 上传成功`)
      } else if (status === 'error') {
        this.$message.error(`${info.file.name} 上传失败`)
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
    onError () {
      this.$alert('文件上传失败，请重试', '错误', {
        confirmButtonText: '确定'
      })
    },
    nextStep () {
    const { form: { validateFields } } = this
      // 先校验，通过表单校验后，才进入下一步
      validateFields((err, values) => {
        console.log('values', values)
        if (!err) {
          this.$emit('nextStep', values)
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
