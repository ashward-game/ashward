import Vue from "vue";
import "primeicons/primeicons.css";
import "primevue/resources/themes/saga-blue/theme.css";
import "primevue/resources/primevue.min.css";
import "primeicons/primeicons.css";

import Toast from "primevue/toast";
import Dialog from "primevue/dialog";
import ToastService from "primevue/toastservice";
import TabView from "primevue/tabview";
import TabPanel from "primevue/tabpanel";

Vue.use(ToastService);
Vue.component("Toast", Toast);
Vue.component("Dialog", Dialog);
Vue.component("TabView", TabView);
Vue.component("TabPanel", TabPanel);
