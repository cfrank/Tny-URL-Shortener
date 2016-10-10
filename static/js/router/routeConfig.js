import Vue from 'vue';
import VueRouter from 'vue-router';
import ShortenView from '../views/ShortenView.vue';
import UnshortenView from '../views/UnshortenView.vue';
import ReportView from '../views/ReportView.vue';

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
        },
        // Todo: history
        {
                path: '/pages/unshorten',
                component: UnshortenView
        },
        {
                path: '/pages/report',
                component: ReportView
        }
]

const routerData = configRoutes.map((obj) => {
        let routeObj = {}
        routeObj['path'] = obj.path;
        routeObj['component'] = obj.component;
        return routeObj;
});

const routes = new VueRouter({
        hashbang: false,
        mode: 'history',
        routes: routerData
});

export default routes;