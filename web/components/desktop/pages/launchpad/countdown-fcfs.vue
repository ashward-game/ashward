<template>
  <div>
    <span id="time_fcfs"></span>
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
    util.startTimer("time_fcfs", process.env.EVENT_TIME_IDO_FCFS, "00:00:00");
  },
  computed: {},
  methods: {},
};
</script>

<style scoped lang="scss" />
