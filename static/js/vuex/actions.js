export const showNotice = ({commit}, payload) => {
        const active = true;
        const message = payload.message;
        const type = payload.type;
        
        commit('SHOW_NOTICE', {active, message, type});
}

export const hideNotice = ({commit}) => {
        commit('HIDE_NOTICE');
}

export const removeUid = ({commit}) => {
        commit('REMOVE_UID');
}

export const showLinkSuccess = ({commit}, payload) => {
        const active = true;
        const title = payload.title;
        const linkHref = payload.linkHref;
        
        commit('SHOW_LINK_SUCCESS', {active, title, linkHref});
}

export const hideLinkSuccess = ({commit}) => {
        commit('HIDE_LINK_SUCCESS');
}

export const updateFormValue = ({commit}, payload) => {
        commit('UPDATE_FORM_VALUE', payload);
}