import Vue from 'vue';
import {sync} from 'vuex-router-sync';

import router from './router/routeConfig';
import store from './vuex/store';
import App from './App.vue';

sync(store, router);

console.log(store);

const app = new Vue({
        router: router,
        store: store,
        render: h => h(App)
}).$mount('.root');