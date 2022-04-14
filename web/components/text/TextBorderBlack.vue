<template>
  <div class="text-border-black">
    <slot />
  </div>
</template>
<script>
export default {
  name: 'TextBorderBlack'
}
</script>
<style lang="scss">
@import "~/static/scss/functions";
.text-border-black{
  text-shadow:  stroke(2,#000,1px);
  @media (min-width:1024px) {
  text-shadow: stroke(3,#000,2px);
  }
&.bold{
  text-shadow: stroke(3,#000,2px);
}
}
</style>
