<template>
  <div class="account-settings-info-view">
    <a-form layout="vertical">
      <a-form-item
        label="API Token"
      >
        <a-input placeholder="API Token" />
      </a-form-item>

      <a-form-item>
        <a-button type="primary">重置</a-button>
      </a-form-item>
    </a-form>
    <div>
      <b>通过API上传文件的参数说明：</b><br/>
      curl -F "file=@filepath.apk" -F "token=123" -F "team_id=1" -F "env=1" -F "desc=我来啦" https://HOST/app/upload <br/>
      其中 file 和 token 为必填，其余为选填。token 即为本页的 API Token。<br/>
      team_id（团队ID）如果不传则匹配有权限的某个团队，env（环境，1为生产2为测试）默认为测试环境，desc（更新说明）默认为空。
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      api_token: ''
    }
  },
  methods: {

  }
}
</script>

<style lang="less" scoped>

</style>
