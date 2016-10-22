import Vue from 'vue';
import VueRouter from 'vue-router';
import ShortenView from '../views/ShortenView.vue';
import UnshortenView from '../views/UnshortenView.vue';
import ReportView from '../views/ReportView.vue';
import HistoryView from '../views/HistoryView.vue';
import AboutView from '../views/AboutView.vue';
import RulesView from '../views/RulesView.vue';

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
        {
                path: '/pages/history',
                component: HistoryView
        },
        {
                path: '/pages/unshorten',
                component: UnshortenView
        },
        {
                path: '/pages/report',
                component: ReportView
        },
        {
                path: '/pages/about',
                component: AboutView
        },
        {
                path: '/pages/rules',
                component: RulesView
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