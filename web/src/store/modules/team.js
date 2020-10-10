import { create } from '@/api/team'

const team = {
  actions: {
    Create ({ commit }, params) {
      return new Promise((resolve, reject) => {
        create(params).then(response => {
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default team
