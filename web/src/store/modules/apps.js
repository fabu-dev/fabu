import { upload } from '@/api/app'

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
    }
  }
}

export default apps
