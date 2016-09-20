import Vue from 'vue/dist/vue';
import App from './App.vue';
import ShortenView from './views/ShortenView.vue';
import VueRouter from 'vue-router';
import {routerData} from './router';

Vue.use(VueRouter);

const routes = new VueRouter({
        history: true,
        saveScrollPosition: true,
        routes: routerData
});

/* Instead of including a polyfill for `...` use
 * Vue.util.extend
 */
new Vue(Vue.util.extend({
        router: routes
}, App)).$mount('.root');