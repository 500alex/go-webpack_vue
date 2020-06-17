import Vue from 'vue';
import $ from "jquery"
// import app from './app.vue';

import component from '../components';
component(Vue);

$(document).ready(function () {
    // window.bus = new Vue();

    var app = new Vue({
        el: '#app',
        delimiters: ['${', '}'],
        data: {
          message: 'Hello Vue!'
        }
      })
})
