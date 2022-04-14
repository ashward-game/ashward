export default {
  data() {
    return {
      isLoading: false,
    };
  },
  computed: {},
  methods: {
    $_globalMixin_loading(time) {
      this.$store.commit("loading", time);
    },
    $_globalMixin_load() {
      this.$store.commit("loading", false);
    },
  },
};
