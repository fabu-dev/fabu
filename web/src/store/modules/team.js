import { create, index } from '@/api/team'

const team = {
  actions: {
    TeamCreate ({ commit }, params) {
      return new Promise((resolve, reject) => {
        create(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    TeamIndex ({ commit }, params) {
      console.log('team get list')
      return new Promise((resolve, reject) => {
        index(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default team
