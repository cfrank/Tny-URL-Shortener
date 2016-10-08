import Vue from 'vue';
import Vuex from 'vuex';
import * as getters from './getters';
import * as actions from './actions';
import mutations from './mutations';

Vue.use(Vuex);

/*
 * State for the notice-header component
 * 
 * Active -> If the component should be shown
 * Message -> The message to show
 * Type -> The type of notification
 *      error -> Problem occured
 *      notification -> Site notification
 *      success -> Successful operation occured
 */
const NOTIFICATION_HEADER = {
        active: false,
        message: '',
        type: '',
}
/*
 * State for the message shown on successful link return
 * 
 * Only shown on link success, errors returning a link show in notification
 * 
 * Active -> If the component should be shown
 * Message -> Title
 * Type -> Link of returned short url
 */
const LINK_SUCCESS = {
        active: false,
        title: '',
        linkHref: '',
}

const state = {
        notice: NOTIFICATION_HEADER,
        linkSuccess: LINK_SUCCESS,
        formValue: ''
};

const store =  new Vuex.Store({
        actions,
        getters,
        mutations,
        state,
});

export default store;