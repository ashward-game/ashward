<template>
  <button class="btn-green" @click="triggerEvent">
    <slot />
  </button>
</template>
<script>
export default {
  name: 'BtnGreen',
  methods: {
    triggerEvent () {
      this.$emit('click')
    }
  }
}
</script>
<style lang="scss" scoped>
.btn-green{
  background-image: url("/assets/button/button-green.png");
  background-size: 100% 100%;
  background-position: center center;
  width: auto;
  height: auto;
  display: inline-block;
  color: #fff;
  padding: 14px 24px;
  }
</style>
