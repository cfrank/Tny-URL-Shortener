import Vue from 'vue';
import Vuex from 'vuex';
import * as getters from './getters';
import * as actions from './actions';
import mutations from './mutations';

Vue.use(Vuex);

/*
 * State for the notice-header component
 * 
 * Show -> If the component should be shown
 * Message -> The message to show
 * Type -> The type of notification
 *      error -> Problem occured
 *      notification -> Site notification
 *      success -> Successful operation occured
 */
const NOTIFICATION_HEADER = {
        show: false,
        message: '',
        type: 'error',
}

const state = {
        notice: NOTIFICATION_HEADER,
};

const store =  new Vuex.Store({
        actions,
        getters,
        mutations,
        state,
});

export default store;