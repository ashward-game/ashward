<template>
  <div class="gate">
    <img v-if="minion" src="/assets/Gate.png" alt="">
    <img v-else src="/assets/gate-not-minion.png" alt="">
  </div>
</template>
<script>
export default {
  name: 'CityGate',
  props: {
    minion: {
      required: false,
      default: true,
      type: Boolean
    }
  }
}
</script>
