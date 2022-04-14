<template>
  <div class="menu-social">
    <div class="header">
      <text-border-black class="text text-31px uppercase">
        {{ name }}
      </text-border-black>
      <img src="/assets/bg-menu-social.png" alt="">
    </div>
    <div class="content">
      <slot />
    </div>
  </div>
</template>
<script>
import TextBorderBlack from '../text/TextBorderBlack.vue'
export default {
  name: 'MenuSocial',
  components: { TextBorderBlack },
  props: {
    name: {
      default: '',
      required: false,
      type: String
    }
  }
}
</script>
<style lang="scss">

</style>
