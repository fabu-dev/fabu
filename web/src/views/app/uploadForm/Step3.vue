<template>
  <div>
    <a-form>
      <a-result title="操作成功" :is-success="true" sub-title="预计两小时内到账" style="max-width: 560px; margin: 40px auto 0;">
        <div class="information">
          <a-row>
            <a-col :sm="8" :xs="24">应用名称：</a-col>
            <a-col :sm="16" :xs="24">{{ sendData.name }}</a-col>
          </a-row>
          <a-row>
            <a-col :sm="8" :xs="24">团队：</a-col>
            <a-col :sm="16" :xs="24">{{ sendData.teamName }}</a-col>
          </a-row>
          <a-row>
            <a-col :sm="8" :xs="24">BundleID：</a-col>
            <a-col :sm="16" :xs="24">{{ sendData.bundle_id }}</a-col>
          </a-row>
          <a-row>
            <a-col :sm="8" :xs="24">版本：</a-col>
            <a-col :sm="16" :xs="24">{{ sendData.version }}</a-col>
          </a-row>
          <a-row>
            <a-col :sm="8" :xs="24">大小：</a-col>
            <a-col :sm="16" :xs="24">{{ (sendData.size/1024/1024).toFixed(2) }} MB</a-col>
          </a-row>
        </div>
        <template #extra>
          <a-button type="primary" @click="finish">继续上传</a-button>
          <a-button style="margin-left: 8px" @click="toOrderList">查看APP</a-button>
        </template>
      </a-result>
    </a-form>
  </div>
</template>

<script>
export default {
  name: 'Step3',
  props: ['sendData'], // 用来接收父组件传给子组件的数据
  data () {
    return {
      loading: false
    }
  },
  mounted () {
    console.log('team id', this.sendData)
  },
  methods: {
    finish () {
      this.$emit('finish')
    },
    toOrderList () {
      this.$router.push('/app/index')
    }
  }
}
</script>
<style lang="less" scoped>
  .information {
    line-height: 22px;

    .ant-row:not(:last-child) {
      margin-bottom: 24px;
    }
  }
  .money {
    font-family: "Helvetica Neue",sans-serif;
    font-weight: 500;
    font-size: 20px;
    line-height: 14px;
  }
</style>
