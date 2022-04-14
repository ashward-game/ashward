<template>
  <div>
    <span id="time_end"></span>
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
    util.startTimer("time_end", "2022-03-18 15:00:00", "00:00:00");
  },
  computed: {},
  methods: {},
};
</script>

<style scoped lang="scss" />
