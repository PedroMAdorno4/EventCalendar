export default {
    userInfo: (state) => {
        if (state.token !== '') {
            return JSON.parse(atob(state.token.split('.')[1]))
        }
        return {}
    },
    strCurrentDate: (state) => {
        let formatter = Intl.DateTimeFormat('pt-BR', { month: 'long', year: 'numeric' });
        return formatter.format(state.currentDate).replace(/^\w/, (c) => c.toUpperCase());
    },
}