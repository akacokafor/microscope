import Vue from 'vue';
import Base from './base';
import axios from 'axios';
import Routes from './routes';
import VueRouter from 'vue-router';
import VueJsonPretty from 'vue-json-pretty';
import moment from 'moment-timezone';

require('bootstrap');

// let token = document.head.querySelector('meta[name="csrf-token"]');

// if (token) {
//     axios.defaults.headers.common['X-CSRF-TOKEN'] = token.content;
// }

// window.Microscope = {
//     timezone: "Africa/Lagos",
//     path: "/microscope",
//     recording: false,
// }

Vue.use(VueRouter);

window.Popper = require('popper.js').default;

moment.tz.setDefault(Microscope.timezone);

window.Microscope.basePath = '/' + window.Microscope.path;

let routerBasePath = window.Microscope.basePath + '/';

if (window.Microscope.path === '' || window.Microscope.path === '/') {
    routerBasePath = '/';
    window.Microscope.basePath = '';
}

const router = new VueRouter({
    routes: Routes,
    mode: 'hash',
    base: routerBasePath,
});

Vue.component('vue-json-pretty', VueJsonPretty);
Vue.component('related-entries', require('./components/RelatedEntries.vue').default);
Vue.component('index-screen', require('./components/IndexScreen.vue').default);
Vue.component('preview-screen', require('./components/PreviewScreen.vue').default);
Vue.component('alert', require('./components/Alert.vue').default);

Vue.mixin(Base);

new Vue({
    el: '#microscope',
    router,
    data() {
        return {
            alert: {
                type: null,
                autoClose: 0,
                message: '',
                confirmationProceed: null,
                confirmationCancel: null,
            },

            autoLoadsNewEntries: localStorage.autoLoadsNewEntries === '1',
            recording: Microscope.recording,
        };
    },

    methods: {
        autoLoadNewEntries() {
            if (!this.autoLoadsNewEntries) {
                this.autoLoadsNewEntries = true;
                localStorage.autoLoadsNewEntries = 1;
            } else {
                this.autoLoadsNewEntries = false;
                localStorage.autoLoadsNewEntries = 0;
            }
        },

        toggleRecording() {
            axios.post(Microscope.basePath + '/microscope-api/toggle-recording');
            window.Microscope.recording = !Microscope.recording;
            this.recording = !this.recording;
        },
    },
});