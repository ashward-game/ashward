<template>
  <div class="text-border-brown">
    <slot />
  </div>
</template>
<script>
export default {
  name: 'TextBorderBrown'
}
</script>
<style lang="scss">
@import "~/static/scss/functions";
.text-border-brown{
text-shadow: stroke(2,#592a04,2px);
}
</style>
