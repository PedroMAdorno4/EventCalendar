export default {
    setToken: (state, token) => {
        state.token = token;
        state.authenticated = true;
    },
    logout: (state) => {
        state.token = "";
        state.authenticated = false;
    }
}
