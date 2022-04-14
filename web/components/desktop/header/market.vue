<template>
  <div>
    <header class="header -v2">
      <div class="container container-1200 d-flex align-items-center">
        <div class="logo">
          <n-link to="/">
            <img
              src="/assets/images/logo.png"
              width="198"
              height="152"
              alt="Logo"
            >
          </n-link>
        </div>

        <nav class="main__nav -v2">
          <ul class="el__menu">
            <li class="menu-item"><n-link to="/">Hero</n-link></li>
            <li class="menu-item"><n-link to="/event">Events box</n-link></li>
            <li class="menu-item"><n-link to="/pets">Pets</n-link></li>
            <li class="menu-item"><n-link to="/items">Items</n-link></li>
          </ul>
        </nav>

        <div class="ms-auto">
          <a href="" class="header__btn btn">Connect wallet</a>
        </div>

        <div
          :class="['menu-mb__btn ms-3', { active: isOpen }]"
          @click="isOpen = !isOpen"
        >
          <span class="iconz-bar" />
          <span class="iconz-bar" />
          <span class="iconz-bar" />
        </div>
      </div>
    </header>
    <nav
      id="menu__mobile"
      :class="['nav__mobile', { active: isOpen }]"
      @click="isOpen = !isOpen"
    >
      <div class="nav__mobile__close icon-close js-menu__close" />
      <div class="nav__mobile__content">
        <ul class="nav__mobile--ul mb-4">
          <li class="menu-item"><n-link to="/">Hero</n-link></li>
          <li class="menu-item"><n-link to="/event">events box</n-link></li>
          <li class="menu-item"><n-link to="/pets">Pets</n-link></li>
          <li class="menu-item"><n-link to="/items">items</n-link></li>
        </ul>
      </div>
    </nav>
  </div>
</template>
<script>
export default {
  components: {},
  mixins: [],
  data () {
    return {
      isOpen: false
    }
  }
}
</script>
<style scoped lang="scss"></style>
