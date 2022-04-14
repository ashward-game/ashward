<template>
  <div class="loading-page">
    <div class="loading-page__logo">
      <img
        src="/assets/images/logo-large.png"
        alt="Logo"
        style="max-width: 383px"
      />
      <p class="loading-page__text mt-4">LOADING {{ timerCount }} %</p>
    </div>
  </div>
</template>
<script>
export default {
  data: () => ({
    timerCount: 0,
  }),
  watch: {
    timerCount: {
      handler(value) {
        const time = this.$store.state.isLoading || 1000;
        if (value < 99) {
          setTimeout(() => {
            this.timerCount += 10;
          }, time / 100);
        }
      },
      immediate: true,
    },
  },
  methods: {},
};
</script>
