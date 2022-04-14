<template>
  <button class="btn-brown-sm">
    <slot />
  </button>
</template>
<script>
export default {
  name: 'BtnBrownSm'
}
</script>
<style lang="scss" scoped>
.btn-brown-sm{
  background-image: url("/assets/button/button-brown-sm.png");
  background-size: 100% 100%;
  padding: 40px 80px;
  display: inline-block;
  cursor: pointer;
  }
</style>
