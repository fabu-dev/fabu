import { getList, deleteVersion } from '@/api/version'

const version = {
  actions: {
    GetVersionList ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getList(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    DeleteVersion ({ commit }, params) {
      return new Promise((resolve, reject) => {
        deleteVersion(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default version
