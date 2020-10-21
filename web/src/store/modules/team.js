import { create, edit, index, getMember } from '@/api/team'

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
    TeamEdit ({ commit }, params) {
      return new Promise((resolve, reject) => {
        edit(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    TeamIndex ({ commit }, params) {
      return new Promise((resolve, reject) => {
        index(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    TeamMember ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getMember(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default team
