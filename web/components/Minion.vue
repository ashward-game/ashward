<template>
  <div class="minion">
    <img v-if="left" src="/assets/minion_1.png" alt="">
    <img v-else src="/assets/minion_2.png" alt="">
  </div>
</template>
<script>
export default {
  name: 'MinionStatue',
  props: {
    left: {
      required: false,
      type: Boolean,
      default: true
    }
  }
}
</script>
