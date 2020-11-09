import { upload, getBase, save } from '@/api/app'

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
    }
  }
}

export default apps
