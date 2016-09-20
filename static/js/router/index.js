import ShortenView from '../views/ShortenView.vue';

/*
 * --IMPORTENT!--
 * Besides the root route, everything must be under the
 * /pages/ directory or else it will not be routed correctly
 * by nginx!
 */
const routes = [
        {
                path: '/',
                component: ShortenView
        }
]

export const routerData = routes.map((obj) => {
        let routeObj = {}
        routeObj['path'] = obj.path;
        routeObj['component'] = obj.component;
        return routeObj;
});