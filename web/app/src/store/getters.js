export default {
    userInfo: (state) => {
        if (state.token !== '') {
            return JSON.parse(atob(state.token.split('.')[1]))
        }
        return {}
    }
}
