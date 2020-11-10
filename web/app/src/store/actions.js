export default {
	setToken: ({ commit }, token) => {
		if (token !== '') {
			commit('setToken', token)
		}
	},
	logout: ({ commit }) => {
			commit('logout')
	}
}
