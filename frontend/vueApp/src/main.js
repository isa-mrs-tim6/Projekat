import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import moment from 'moment'
import * as VueGoogleMaps from "vue2-google-maps";

Vue.config.productionTip = false;
Vue.filter('moment', function(value) {
  if (value) {
    return moment(String(value)).format('HH:mm DD/MM/YYYY')
  }
});
Vue.use(VueGoogleMaps, {
  load: {
    key: "AIzaSyAmI3ZEoWD24xABBLQCLXM4y7_2A_fvPfI",
    libraries: "places" // necessary for places input
  }
});

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
