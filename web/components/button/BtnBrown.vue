<template>
  <button class="btn-brown">
    <slot />
  </button>
</template>
<script>
export default {
  name: 'BtnBrown'
}
</script>
<style lang="scss" scoped>
.btn-brown{
  background-image: url("/assets/button/button-brown.png");
  background-size: 100% 100%;
  padding: 40px 80px;
  display: inline-block;
  cursor: pointer;
  }
</style>
