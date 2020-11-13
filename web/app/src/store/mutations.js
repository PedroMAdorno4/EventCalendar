export default {
    setToken: (state, token) => {
        state.token = token;
        state.authenticated = true;
    },
    logout: (state) => {
        state.token = "";
        state.authenticated = false;
    },
    setCurrentDate: (state, date) => {
        state.currentDate = date;
    },
    setEvents: (state, events) => {
        state.events = events;
    },
    setCurrentEvent: (state, event) => {
        state.currentEvent = event;
    },
    addEvent: (state, event) => {
        if (event !== null && Object.keys(event).length > 0) {
            if (state.events === null) {
                state.events = [];
            }
            state.events.push(event);
        }
    },
    removeEvent: (state, id) => {
        if (id !== null && id !== "") {   
            state.events = state.events.filter(e => e._id !== id);
        }
    },
}
