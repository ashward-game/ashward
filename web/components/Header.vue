<template>
  <div :class="`header ${showMenu?'show-menu':''}`">
    <NavBar @toggle-menu="toggleMenu" />
  </div>
</template>

<script>
export default {
  name: 'HeaderPage',
  data () {
    return {
      showMenu: false
    }
  },
  methods: {
    toggleMenu (showMenu) {
      this.showMenu = showMenu
    }
  }
}
</script>
<style lang="scss" scoped>
.header{
  background-image: url("/assets/Header_BG.png");
  background-repeat: repeat-x;
  background-size: auto 100%;
  position: fixed;
  top:0;
  left: 0;
  z-index: 1030;
  width: 100%;
  padding-bottom: 4rem;
  &.show-menu{
    height: 100%;
    padding-bottom: 0;
    background-image: none;
  }
}
</style>
