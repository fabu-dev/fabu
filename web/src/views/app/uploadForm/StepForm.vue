<template>
  <page-header-wrapper>
    <!-- PageHeader 第二种使用方式 (v-slot) -->
    <template v-slot:content>
      相同 bundle_id 的文件视为同一应用
    </template>
    <a-card :bordered="false">
      <a-steps class="steps" :current="currentTab">
        <a-step title="上传APP" />
        <a-step title="填写APP信息" />
        <a-step title="完成" />
      </a-steps>
      <div class="content">
        <step1 v-if="currentTab === 0" @nextStep="nextStep"/>
        <step2 v-if="currentTab === 1" @nextStep="nextStep" @prevStep="prevStep" :sendData="stepOne"/>
        <step3 v-if="currentTab === 2" @prevStep="prevStep" @finish="finish" :sendData="stepOne"/>
      </div>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import Step1 from './Step1'
import Step2 from './Step2'
import Step3 from './Step3'

export default {
  name: 'StepForm',
  components: {
    Step1,
    Step2,
    Step3
  },
  data () {
    return {
      stepOne: {},
      currentTab: 0,
      // form
      form: null
    }
  },
  methods: {

    // handler
    nextStep (data) {
      console.log('tab is: ', this.currentTab)
      if (this.currentTab === 0) {
        console.log('target:', data.file.file.response)
        console.log('target:', data.file.file.response.result)
        const result = data.file.file.response.result
        const stepOneData = {
          team_id: data.team_id,
          team_name: data.team_name,
          build: result.build,
          bundle_id: result.bundle_id,
          icon: process.env.VUE_APP_API_BASE_URL + '/' + result.icon,
          identifier: result.identifier,
          name: result.name,
          size: result.size,
          version: result.version,
          shortKey: process.env.VUE_APP_BASE_URL + '/s/' + result.short_key,
          qrCode: process.env.VUE_APP_API_BASE_URL + '/' + result.qr_code
        }
        this.stepOne = stepOneData
      }

      if (this.currentTab === 1) {
        console.log('step2 data', data)
        this.stepOne.teamName = data.teamName
        this.stepOne.description = data.description
      }

      console.log('current tab:', this.currentTab)
      console.log('stepOne:', this.stepOne)
      if (this.currentTab < 2) {
        this.currentTab += 1
      }
    },
    prevStep () {
      if (this.currentTab > 0) {
        this.currentTab -= 1
      }
    },
    finish () {
      this.currentTab = 0
    }
  }
}
</script>

<style lang="less" scoped>
  .steps {
    max-width: 750px;
    margin: 16px auto;
  }
</style>
