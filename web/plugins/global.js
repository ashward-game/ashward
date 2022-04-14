import Vue from "vue";
import VueClipboard from "vue-clipboard2";
import VueScrollTo from "vue-scrollto";
import "@fortawesome/fontawesome-free/css/all.css";
import VueLazyload from "vue-lazyload";
import { Modal } from "bootstrap";

VueClipboard.config.autoSetContainer = true;
Vue.use(VueClipboard);
Vue.use(VueScrollTo);

const loadimage = require("../static/assets/images/loading.gif");
Vue.use(VueLazyload, {
  preLoad: 1.3,
  error: loadimage,
  loading: loadimage,
  attempt: 1,
  listenEvents: ["scroll"],
});
Vue.mixin({
  methods: {
    getBootstrapModal(id) {
      return new Modal(id);
    },
  },
});
