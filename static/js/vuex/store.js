import Vue from 'vue';
import Vuex from 'vuex';
import * as getters from './getters';

Vue.use(Vuex);

const state = {
        testing: 'Hello World!',
};

const store =  new Vuex.Store({
        state,
        getters,
});

export default store;