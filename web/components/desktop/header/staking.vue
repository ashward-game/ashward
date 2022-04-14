<template>
  <div>
    <div class="header__top">
      <div class="container d-flex align-items-center justify-content-center">
        <div class="d-flex align-items-center">
          <div class="header__top__address text-truncate me-2">
            BEP-20: To be announced
          </div>
          <!-- <i class="ic-copy"></i> -->
        </div>
      </div>
    </div>
    <header ref="headerTop" class="header">
      <div class="container container-1200 d-flex align-items-center">
        <div class="logo">
          <n-link to="/">
            <img
              src="/assets/images/logo.png"
              width="198"
              height="152"
              alt="Logo"
            />
          </n-link>
        </div>

        <nav class="main__nav">
          <ul class="el__menu">
            <li class="menu-item active"><n-link to="/">HOME</n-link></li>
            <li class="menu-item">
              <a href="https://doc.ashward.io/" target="_blank"> LITEPAPER </a>
            </li>
            <li class="menu-item">
              <a href="https://events.ashward.io/" target="_blank"> EVENTS </a>
            </li>
            <li class="menu-item"><n-link to="/trade">MARKETPLACE</n-link></li>
          </ul>
        </nav>

        <div class="ms-auto">
          <n-link to="/staking" class="header__btn btn">Staking</n-link>
        </div>

        <div
          class="menu-mb__btn ms-3"
          data-id="menu__mobile"
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
          <li class="menu-item"><n-link to="/">HOME</n-link></li>
          <li class="menu-item">
            <a href="https://doc.ashward.io/" target="_blank"> LITEPAPER </a>
          </li>
          <li class="menu-item">
            <a href="https://events.ashward.io/" target="_blank"> EVENTS </a>
          </li>
          <li class="menu-item"><n-link to="/trade">MARKETPLACE</n-link></li>
        </ul>
      </div>
    </nav>
    <div
      id="stakingModal"
      class="modal fade -staking"
      tabindex="-1"
      aria-labelledby="exampleModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content">
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            aria-label="Close"
          />
          <div class="modal-body">
            <button class="btn w-100">Coming soon.</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  components: {},
  mixins: [],
  data() {
    return {
      isOpen: false,
      textAddress: "BEP-20: 0x9fd87aefe02441b123c3c32466cd9db4c578618f",
    };
  },
  mounted() {},
  destroyed() {},
  methods: {
    copy() {
      this.textAddress = "copied";
      document.execCommand("0x9fd87aefe02441b123c3c32466cd9db4c578618f");
      setTimeout(() => {
        this.textAddress = "BEP-20: 0x9fd87aefe02441b123c3c32466cd9db4c578618f";
      }, 3000);
    },
  },
};
</script>
<style scoped lang="scss"></style>
