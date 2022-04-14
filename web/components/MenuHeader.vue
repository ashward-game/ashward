<template>
  <div
    class="menu-header lg:mx-8 flex lg:space-x-8 2xl:space-x-10 uppercase xl:text-xl xl:mx-0 2xl:text-2xl"
  >
    <a href="/" @click="showMenu">
      <text-border-black class="text-white"> Home </text-border-black>
    </a>
    <a href="/genesis" @click="showMenu">
      <text-border-black class="text-white"> Genesis </text-border-black>
    </a>
    <a href="/inventory" @click="showMenu">
      <text-border-black class="text-white"> Inventory </text-border-black>
    </a>
    <!--    <nuxt-link to="#features" @click.native="showMenu">-->
    <!--      <text-border-black class="text-white"> Features </text-border-black>-->
    <!--    </nuxt-link>-->
    <!--    <nuxt-link to="#road-map" @click.native="showMenu">-->
    <!--      <text-border-black class="text-white"> Roadmap </text-border-black>-->
    <!--    </nuxt-link>-->
    <!--    <nuxt-link to="#our-team" @click.native="showMenu">-->
    <!--      <text-border-black class="text-white"> Team </text-border-black>-->
    <!--    </nuxt-link>-->
    <!--    <nuxt-link to="#partner" @click.native="showMenu">-->
    <!--      <text-border-black class="text-white"> Partners </text-border-black>-->
    <!--    </nuxt-link>-->
    <hr class="line" />
  </div>
</template>
<script>
export default {
  name: "MenuHeader",
  methods: {
    showMenu() {
      this.$emit("showMenu", false);
    },
  },
};
</script>
<style lang="scss" scoped>
.line {
  display: none;
}
.show-menu {
  .line {
    display: block;
    width: 48px;
    height: 2px;
  }
}
.menu-header {
  white-space: nowrap;
}
</style>
