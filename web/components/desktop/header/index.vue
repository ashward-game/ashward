<template>
  <div>
    <div class="header__top -absolute">
      <div class="container d-flex align-items-center justify-content-center">
        <div class="d-flex align-items-center">
          <div class="header__top__address text-truncate me-2">
            BEP-20: {{ ascToken }}
          </div>
          <i class="ic-copy" @click="copyText('token-address')" />
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
            <li class="menu-item"><n-link to="/">HOME</n-link></li>
            <li class="menu-item">
              <a href="https://doc.ashward.io/" target="_blank"> LITEPAPER </a>
            </li>
            <!--            <li class="menu-item"><n-link to="/">MARKETPLACE</n-link></li>-->
            <li class="menu-item"><n-link to="/staking">STAKING</n-link></li>
            <li class="menu-item">
              <n-link to="/launchpad">LAUNCHPAD</n-link>
            </li>
          </ul>
        </nav>

        <!--        <div class="ms-auto">-->
        <!--          <a href="" class="header__btn btn">Connect wallet</a>-->
        <!--        </div>-->

        <div
          class="ms-auto"
          @mouseover="isOpenWallet = true"
          @mouseleave="isOpenWallet = false"
        >
          <button class="header__btn btn" v-if="account">
            {{ textAddress(account) }}
          </button>
          <button class="header__btn btn" v-else @click="connectWallet()">
            Connect wallet
          </button>
          <div class="absolute top-16 pt-0 md:pt-4" v-if="account">
            <button
              v-show="isOpenWallet"
              class="btn-disconnect-wallet flex bg-gray-500 rounded-md mt-2 border-2 border-rose-900 mt-2"
              @click="disconnectWallet()"
            >
              <div class="flex">
                <img src="/assets/logout.svg" alt="" />
                <span class="ml-2">Disconnect</span>
              </div>
            </button>
          </div>
        </div>

        <div
          class="menu-mb__btn ms-3"
          data-id="menu__mobile"
          @click="showMenu = !showMenu"
        >
          <span class="iconz-bar" />
          <span class="iconz-bar" />
          <span class="iconz-bar" />
        </div>
      </div>
    </header>
    <nav
      id="menu__mobile"
      :class="['nav__mobile', { active: showMenu }]"
      @click="showMenu = !showMenu"
    >
      <div class="nav__mobile__close icon-close js-menu__close" />
      <div class="nav__mobile__content">
        <ul class="nav__mobile--ul mb-4">
          <li class="menu-item"><n-link to="/">HOME</n-link></li>
          <li class="menu-item">
            <n-link to="https://doc.ashward.io/" target="_blank">
              LITEPAPER
            </n-link>
          </li>
          <li class="menu-item"><n-link to="/staking">STAKING</n-link></li>
          <li class="menu-item"><n-link to="/launchpad">LAUNCHPAD</n-link></li>
        </ul>
      </div>
      <input type="hidden" id="token-address" :value="ascToken" readonly />
    </nav>
  </div>
</template>
<script>
import $ from "jquery";
import { ADDRESS_TOKEN } from "~/plugins/contracts/Token";

export default {
  components: {},
  mixins: [],
  data() {
    return {
      showMenu: false,
      scrollHeight: 0,
      isSticky: false,
      interval: null,
      isOpenWallet: false,
      ascToken: "",
    };
  },

  watch: {
    // scrollHeight(val) {
    //   this.isSticky = val > 100;
    // },
  },
  computed: {
    account() {
      if (this.$store.state.ether?.web3Account)
        return this.$store.state.ether?.web3Account;
    },
  },
  async mounted() {
    this.ascToken = ADDRESS_TOKEN;
    const web3ModalConnected = localStorage.getItem("web3ModalConnected");
    if (web3ModalConnected) {
      await this.$ether();
    }
    // this.scrollHeight = this.$refs.headerTop.offsetTop;
    // this.interval = setInterval(() => {
    //   this.scrollHeight = $(".header.-fix").offset().top;
    // }, 100);
  },
  destroyed() {
    clearInterval(this.interval);
  },
  methods: {
    textAddress(text) {
      return text.slice(0, 5) + "..." + text.slice(-5);
    },
    async connectWallet() {
      this.isOpenWallet = false;
      await this.$ether();
    },
    async disconnectWallet() {
      await this.$disconnectWallet();
    },
    async openDisconnect() {
      this.isOpenWallet = !this.isOpenWallet;
    },
    copyText(id) {
      let testingCodeToCopy = document.querySelector(`#${id}`);
      testingCodeToCopy.setAttribute("type", "text"); // 不是 hidden 才能複製
      testingCodeToCopy.select();

      try {
        var successful = document.execCommand("copy");
        var msg = successful ? "successful" : "unsuccessful";
        this.$toast.add({
          severity: "success",
          summary: "Copied " + msg,
          life: 3000,
        });
      } catch (err) {
        alert("Oops, unable to copy");
      }

      testingCodeToCopy.setAttribute("type", "hidden");
      window.getSelection().removeAllRanges();
    },
  },
};
</script>
<style scoped lang="scss"></style>
