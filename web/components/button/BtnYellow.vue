<template>
  <button class="btn-yellow px-4 py-2" @click="triggerEvent">
    <slot />
  </button>
</template>
<script>
export default {
  name: 'BtnYellow',
  methods: {
    triggerEvent () {
      this.$emit('click')
    }
  }
}
</script>
<style lang="scss" scoped>
.btn-yellow{
  background-image: url("/assets/button/button-yellow.png");
  background-size: 100% 100%;
  }
</style>
