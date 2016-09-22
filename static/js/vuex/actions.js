export const showNotice = ({commit}, payload) => {
        const show = true;
        const message = payload.message;
        const type = payload.type;
        
        commit('SHOW_NOTICE', {show, message, type});
}

export const hideNotice = ({commit}) => {
        commit('HIDE_NOTICE');
}

export const removeUid = ({commit}) => {
        commit('REMOVE_UID');
}