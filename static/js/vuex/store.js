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

/*
 * State for history cache
 * 
 * Keeps a cache of the history links array, and length as well
 * as a variable telling if the cache is stale.
 * 
 * Used so that a db/api call is needed everytime the user goes to the history
 * page
 * 
 * links -> the history link array
 * length -> int representing the number of history items
 * stale -> if true refresh the cache
 */
const HISTORY_CACHE = {
        links: [],
        length: 0,
        stale: true,
}

const state = {
        notice: NOTIFICATION_HEADER,
        linkSuccess: LINK_SUCCESS,
        formValue: '',
        historyCache: HISTORY_CACHE
};

const store =  new Vuex.Store({
        actions,
        getters,
        mutations,
        state,
});

export default store;