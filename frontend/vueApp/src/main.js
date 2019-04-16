import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import moment from 'moment'

Vue.config.productionTip = false;
Vue.filter('moment', function(value) {
  if (value) {
    return moment(String(value)).format('HH:mm DD/MM/YYYY')
  }
});

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
