import { upload, getBase, save, getList, getInfo, deleteApp } from '@/api/app'

const apps = {
  actions: {
    UploadApp ({ commit }, params) {
      return new Promise((resolve, reject) => {
        upload(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    SaveApp ({ commit }, params) {
      return new Promise((resolve, reject) => {
        save(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    GetBase ({ commit }, params) {
      console.log('apps get base params', params)
      return new Promise((resolve, reject) => {
        getBase(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    GetList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    GetAppInfo ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getInfo(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    DeleteApp ({ commit }, params) {
      return new Promise((resolve, reject) => {
        deleteApp(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default apps
