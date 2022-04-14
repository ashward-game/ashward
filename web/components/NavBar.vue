<template>
  <div class="nav-bar">
    <logo />
    <button
      :class="`hamburger hamburger--collapse ${showMenu ? 'is-active' : ''}`"
      type="button"
      @click="toggleMenu(!showMenu)"
    >
      <span class="hamburger-box">
        <span class="hamburger-inner" />
      </span>
    </button>
    <div v-show="showMenu" class="collapse w-full flex justify-center">
      <menu-header class="ml-0 lg:ml-4 xl:ml-10" @showMenu="toggleMenu" />
    </div>
    <div v-show="!showMenu" class="flex connect-wallet-button justify-end">
      <div
        class="flex flex-col"
        @mouseover="isOpenWallet = true"
        @mouseleave="isOpenWallet = false"
      >
        <button v-if="account" class="btn-connect-wallet flex">
          <span>{{ textAddress(account) }}</span>
        </button>

        <button v-else class="btn-connect-wallet flex" @click="connectWallet()">
          <span>CONNECT WALLET</span>
        </button>
        <div class="absolute top-24 pt-0 md:pt-4">
          <NuxtLink
            to="/inventory"
            v-show="isOpenWallet"
            class="btn-disconnect-wallet flex bg-green-500 rounded-md border-2 border-rose-900"
          >
            <div class="flex">
              <span class="ml-2">Inventory</span>
            </div>
          </NuxtLink>
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
    </div>
  </div>
</template>
<script>
export default {
  name: "NavBar",
  data() {
    return {
      showMenu: false,
      isOpenWallet: false,
    };
  },
  computed: {
    account() {
      if (this.$store.state.ether?.web3Account)
        return this.$store.state.ether?.web3Account;
    },
  },
  async mounted() {
    const web3ModalConnected = localStorage.getItem("web3ModalConnected");
    if (web3ModalConnected) {
      await this.$ether();
    }
  },
  methods: {
    toggleMenu(closeMenu) {
      this.showMenu = closeMenu;
      this.$emit("toggle-menu", closeMenu);
    },
    openAccount() {
      this.$router.push({ path: "/account" });
    },
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
  },
};
</script>
<style lang="scss" scope>
.nav-bar {
  background-color: transparent;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  position: relative;
  padding-left: 18px;
  padding-right: 20px;
  padding-top: 18px;
  transition: 0.4s;
  .hamburger {
    padding: 0;
    .hamburger-box {
      width: 32px;
      .hamburger-inner,
      .hamburger-inner::before,
      .hamburger-inner::after {
        width: 32px;
        background: #fff;
        border-radius: 0;
      }
    }
  }
  .collapse {
    .menu-header {
      flex-direction: column;
      align-items: center;
    }
  }
  .logo {
    width: 72px;
    height: 58px;
  }
  @media (min-width: 640px) {
    padding-left: 23px;
  }
  @media (min-width: 1024px) {
    padding-left: 43px;
    .hamburger {
      display: none;
    }
    justify-content: flex-start;
    .connect-wallet {
      font-size: 16px;
      padding: 10px 18px;
    }
    .logo {
      width: 130px;
      height: 90px;
    }
    .collapse {
      background: transparent;
      display: flex !important;
      align-items: center;
      position: relative;
      flex-direction: row;
      .close-menu {
        display: none;
      }
      .menu-header {
        flex-direction: row;
        align-items: center;
      }
    }
  }
  @media (min-width: 1280px) {
    .connect-wallet {
      font-size: 18px;
      padding: 14px 20px;
    }
    .logo {
      width: 140px;
      height: 100px;
    }
  }
  @media (min-width: 1536px) {
    .connect-wallet {
      padding: 14px 24px;
      font-size: 28px;
    }
    .logo {
      width: 145px;
      height: 105px;
    }
  }
}
.show-menu {
  .nav-bar {
    overflow: auto;
    transition: 0.4s;
    height: 100%;
    width: 100%;
    flex-direction: column;
    justify-content: flex-start;
    background: rgba(#000000, 85%);
    padding-top: 30px;
    .hamburger {
      order: 0;
      margin-left: auto;
      position: fixed;
      top: 2.2rem;
      right: 1rem;
      .hamburger-inner,
      .hamburger-inner::before,
      .hamburger-inner::after {
        background: #ffe205;
      }
    }
    .logo {
      order: 1;
      width: 127px;
      height: 93px;
      margin-bottom: 30px;
      margin-top: -3px;
    }
    .collapse {
      order: 3;
      .menu-header {
        font-size: 24px;

        & > a {
          margin-top: 12px;
          margin-bottom: 12px;
        }
      }
    }
  }
}
.connect-wallet-button {
  text-align: right;
  display: flex;
  right: 20px;
  top: 20px;
}
.btn-connect-wallet {
  background-image: url("/assets/wallet.png");
  background-size: 100% 100%;
  background-position: center center;
  background-repeat: no-repeat;
  height: 40px;
  width: 200px;
  color: white;
  align-items: center;
  justify-content: center;
  padding: 30px 0;
}
.btn-disconnect-wallet {
  height: 50px;
  width: 150px;
  color: white;
  align-items: center;
  justify-content: center;
}
</style>
