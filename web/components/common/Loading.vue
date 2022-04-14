<template>
  <div v-if="loading" class="loading-page">
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
    loading: false,
    timerCount: 1,
  }),
  watch: {
    timerCount: {
      handler(value) {
        if (value < 99) {
          setTimeout(() => {
            this.timerCount++;
          }, 10);
        }
      },
      immediate: true,
    },
  },
  methods: {
    start() {
      this.loading = true;
    },
    finish() {
      this.timerCount = 1;
      this.loading = false;
    },
  },
};
</script>
