<template>
  <div class="wrap__page litepaper">
    <Tab />
    <div class="container container-580">
      <Token />

      <ConnectWallet
        v-if="!isConnectWallet"
        :handle-connect-wallet="handleConnectWallet"
      />

      <!-- <WalletNotWhiteList
        v-if="isConnectWallet"
        :handle-dis-connect-wallet="handleDisConnectWallet"
      /> -->

      <!--      <WalletConnected-->
      <!--        v-if="isConnectWallet"-->
      <!--        :is-enable="isEnable"-->
      <!--        :bought="bought"-->
      <!--        :handle-enable="handleEnable"-->
      <!--        :handle-buy="handleBuy"-->
      <!--        :handle-dis-connect-wallet="handleDisConnectWallet"-->
      <!--      />-->
      <Info />
    </div>
  </div>
</template>

<script>
import Tab from "@/components/desktop/pages/launchpad/tab.vue";
import Token from "@/components/desktop/pages/launchpad/token.vue";
import Info from "@/components/desktop/pages/launchpad/info.vue";
import ConnectWallet from "@/components/desktop/pages/launchpad/connect-wallet.vue";

export default {
  components: {
    Tab,
    Token,
    Info,
    ConnectWallet,
  },
  data() {
    return {
      isConnectWallet: false,
      isEnable: false,
      bought: false,
    };
  },
  methods: {
    handleConnectWallet() {
      this.$toast.add({
        severity: "success",
        summary: "Connect Wallet Success",
        life: 1000,
      });
      this.isConnectWallet = true;
    },
    handleDisConnectWallet() {
      this.isConnectWallet = false;
      this.isEnable = false;
      this.bought = false;
    },
    handleEnable() {
      // handle enable
      this.isEnable = true;
    },

    handleBuy(optionBuy) {
      // handle buy
      console.log(optionBuy);
      this.bought = true;
    },
  },
};
</script>
