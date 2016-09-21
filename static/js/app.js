import Vue from 'vue/dist/vue';
import Vuex from 'vuex';
import {routes} from './router';
import App from './App.vue';

Vue.use(Vuex); // Expose Vuex to Vue

/* Instead of including a polyfill for `...` use
 * Vue.util.extend
 */
new Vue(Vue.util.extend({
        router: routes
}, App)).$mount('.root');