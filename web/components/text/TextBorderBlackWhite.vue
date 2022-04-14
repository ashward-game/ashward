<template>
  <div class="text-border-black-white">
    <slot />
  </div>
</template>
<script>
export default {
  name: 'TextBorderBlackWhite'
}
</script>
<style lang="scss">
@import "~/static/scss/functions";
.text-border-black-white{
  text-shadow:  stroke(2, #000,2px),stroke(4, #fff,2px);
  &.bold{
  text-shadow:  stroke(4, #000,2px),stroke(8, #fff,2px);
  }
}
</style>
