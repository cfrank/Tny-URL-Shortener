import Vue from 'vue/dist/vue';
import App from '../App.vue';

const app = new Vue({
        el: '.root',
        components: {
                App
        }
});

app.$mount('.root');