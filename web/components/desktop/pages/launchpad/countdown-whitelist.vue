<template>
  <div>
    <span id="time_whitelist"></span>
  </div>
</template>
<script>
import util from "~/plugins/lib/util";

export default {
  components: {},
  props: {
    time: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {};
  },
  mounted() {
    util.startTimer(
      "time_whitelist",
      process.env.EVENT_TIME_IDO_WHITELIST,
      "00:00:00"
    );
  },
  computed: {},
  methods: {},
};
</script>

<style scoped lang="scss" />
