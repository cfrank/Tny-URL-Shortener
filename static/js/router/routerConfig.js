import Vue from 'vue/dist/vue';
import VueRouter from 'vue-router';
import ShortenView from '../views/ShortenView.vue';

Vue.use(VueRouter); // Expose Vuerouter to Vue

/*
 * --IMPORTENT!--
 * Besides the root route, everything must be under the
 * /pages/ directory or else it will not be routed correctly
 * by nginx!
 */

const configRoutes = [
        {
                path: '/',
                component: ShortenView
        }
]

const routerData = configRoutes.map((obj) => {
        let routeObj = {}
        routeObj['path'] = obj.path;
        routeObj['component'] = obj.component;
        return routeObj;
});

export const routes = new VueRouter({
        hashbang: false,
        mode: 'history',
        routes: routerData
});