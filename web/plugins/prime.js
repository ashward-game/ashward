import Vue from 'vue'
import PrimeVue from 'primevue/config'
import Toast from 'primevue/toast'
import Button from 'primevue/button'
import ToastService from 'primevue/toastservice'
import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.css'

Vue.use(PrimeVue, { ripple: true })
Vue.use(ToastService)
Vue.component('Toast', Toast)
Vue.component(Button)
// export default (context, inject) => {
//   const toast = () => {}
//   // Inject $hello(msg) in Vue, context and store.
//   inject('toast', toast)
//   // For Nuxt <= 2.12, also add 👇
//   context.$toast = toast
// }
