export default {
	setToken: ({ commit }, token) => {
		if (token !== '') {
			commit('setToken', token)
		}
	},
	logout: ({ commit }) => {
			commit('logout')
	},
	setCurrentDate: ({ commit }, date) => {
		commit('setCurrentDate', date)
	},
	setEvents: ({ commit }, events) => {
		commit('setEvents', events)
	},
	setCurrentEvent: ({ commit }, event) => {
		commit('setCurrentEvent', event)
	},
	addEvent: ({ commit }, event) => {
		commit('addEvent', event)
	},
	removeEvent: ({ commit }, id) => {
		commit('removeEvent', id)
	},
}
