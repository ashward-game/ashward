<template>
  <div class="container mx-auto min-h-screen">
    <div class="header h-20"></div>
    <Dialog
      :visible.sync="isShow"
      :modal="true"
      class="popup-intro"
      :dismissableMask="true"
      :showHeader="false"
    >
      <div class="p-dialog-content__content">
        <div>
          <!--          <div class="title">Import your NFT to wallet</div>-->
          <div class="paragraph">
            Import NFTs <br />
            <div class="flex items-end justify-between mb-2">
              Contract address: {{ textAddress(nftAddress) }}
              <img
                src="/assets/icon/copy.svg"
                alt=""
                class="ml-2 w-7 h-7"
                @click="copyText('nft-address')"
              />
            </div>
            Token ID: <br> Check inventory
          </div>
        </div>
        <input type="hidden" id="nft-address" :value="nftAddress" readonly />
        <div class="foot">
          <button class="button-close" @click="isShow = !isShow">Close</button>
        </div>
      </div>
    </Dialog>
    <div class="flex pt-16 sm:mt-20 justify-center w-full">
      <button
        class="text-xl md:text-3xl text-white ml-2 text-yellow"
        @click="isShow = !isShow"
      >
        How to import NFTs?
      </button>
    </div>
    <div
      class="flex flex-col sm:flex-row justify-center w-full mt-10"
      v-if="nfts.length > 0"
    >
      <div
        class="flex justify-center items-center m-4 flex-col"
        v-for="(item, key) in nfts"
        :key="key"
      >
        <NuxtLink :to="`/genesis/${item.token_id}`">
          <img
            class="w-full sm:w-auto"
            style="max-width: 250px; max-height: 360px"
            :src="item.image"
            alt=""
          />
        </NuxtLink>
        <span class="text-2xl text-white mt-4">{{ item.name }}</span>
        <span class="text-2xl text-white">#: {{ item.token_id }}</span>
      </div>
    </div>
    <div v-else class="flex flex-col sm:flex-row justify-center w-full mt-10">
      <span class="text-3xl text-white text-center"
        >Open box to claim NFT!</span
      >
    </div>
  </div>
</template>

<script>
import { ethers } from "ethers";
import { ADDRESS_NFT } from "~/plugins/contracts/NFT";
const BN = require("bn.js");
import { mapGetters, mapState } from "vuex";

export default {
  layout: "genesis",
  data() {
    return {
      isShow: false,
      nftAddress: ADDRESS_NFT,
      chainID: process.env.CHAIN_ID,
      RPC_PROVIDER: process.env.RPC_PROVIDER,
    };
  },
  watch: {
    async account() {
      await this.$store.dispatch("account/setQuery", {
        address: this.account,
      });
    },
  },
  async mounted() {
    console.log(ADDRESS_NFT);
    if (this.account)
      await this.$store.dispatch("account/setQuery", {
        address: this.account,
      });
  },
  computed: {
    ...mapState("nft", {
      nft: (state) => state.nft,
    }),
    ...mapState("account", {
      isLoading: (state) => state.isLoading,
      total: (state) => state.total,
      queryObject: (state) => state.queryObject,
    }),
    ...mapGetters({
      nfts: "account/listNfts",
    }),
    provider() {
      if (this.$store.state.ether?.isConnectedWeb3) {
        return this.$web3Provider();
      } else {
        return new ethers.providers.JsonRpcProvider(process.env.RPC_PROVIDER);
      }
    },
    account() {
      if (
        this.$store.state.ether?.web3Account &&
        this.$store.state.ether?.isConnectedWeb3
      )
        return this.$store.state.ether?.web3Account;
    },
  },
  methods: {
    textAddress(text) {
      return text.slice(0, 5) + "..." + text.slice(-5);
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
<style lang="scss">
@import "~/static/scss/pages/index";
</style>

<style lang="scss">
body {
  background-image: url("/web/static/assets/Pattern.png"),
    url("/web/static/assets/bg-brown.png");
  background-repeat: repeat;
  background-color: #381d24;
}
.popup-intro {
  box-shadow: none;
  &.p-dialog {
    width: 800px;
    height: 700px;
    position: absolute;
    top: 57% !important;
    left: 50% !important;
    transform: translate(-50%, -50%);
    @media (max-width: 1024px) {
      width: 90%;
    }
    .p-dialog-content {
      background: url("/assets/openbox/popup_claimnft.png") no-repeat !important;
      background-size: 100% 100% !important;
      max-width: 350px;
      max-height: 450px;
      min-height: 350px;
      position: relative;
      padding-top: 3rem;
      background-color: transparent;
      margin: 0 auto;
      margin-top: 30px;
      position: relative;
      overflow: hidden;
      .p-dialog-content__content {
        position: relative;
        width: 100%;
        height: 100%;
        left: 0;
        padding: 0 0.5rem;
        display: flex;
        flex-direction: column;
        .title {
          color: #ffcf64;
          font-family: "UTM Bienvenue", sans-serif;
          font-size: 37px;
          text-align: center;
          font-weight: bold;
          margin-bottom: 1rem;
        }
        .paragraph {
          font-family: "Pixel";
          color: #fff;
          font-size: 18px;
          line-height: 24px;
          word-wrap: break-word;
        }
        .foot {
          flex: 1;
          display: flex;
          align-items: flex-end;
          .button-close {
            position: relative;
            left: 50%;
            bottom: 0rem;
            transform: translateX(-50%);
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 20px;
            color: #5f2e1a;
            background: #ffdf09;
            border: 3px solid #5f2e1a;
            box-shadow: inset 0px -5px 0px #be791f;
            border-radius: 8px;
            padding: 0.15rem 3rem;
            width: 200px !important;
            height: 54px;
          }
        }
      }
      @media (max-width: 640px) {
        padding-left: 0.2rem;
        padding-right: 0.2rem;
        .p-dialog-content__content {
          padding: 0;
          .title {
            font-size: 20px;
          }
          .paragraph {
            font-size: 16px;
          }
          .foot .button-close {
            width: 80%;
          }
        }
      }
    }
  }
}
</style>
