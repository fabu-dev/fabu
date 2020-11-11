import { create, edit, index, getMember, exit, del, addMember, getTeamInfo } from '@/api/team'

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
    TeamDelete ({ commit }, params) {
      return new Promise((resolve, reject) => {
        del(params).then(response => {
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
    TeamInfo ({ commit }, params) {
      return new Promise((resolve, reject) => {
        getTeamInfo(params).then(response => {
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
    },
    TeamMemberExit ({ commit }, params) {
      return new Promise((resolve, reject) => {
        exit(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    TeamMemberAdd ({ commit }, params) {
      return new Promise((resolve, reject) => {
        addMember(params).then(response => {
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    }
  }
}

export default team
