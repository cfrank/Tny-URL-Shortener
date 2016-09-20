import Vue from 'vue/dist/vue';
import App from './App.vue';

const app = new Vue({
        // Mount to element
        el: '.root',
        
        // Register components
        components: {
                App
        }
});