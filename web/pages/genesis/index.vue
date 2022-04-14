<template>
  <div>
    <Dialog
      :visible.sync="isShowNFT"
      :modal="true"
      class="popup-nft"
      :dismissableMask="true"
      :showHeader="false"
      v-if="nft"
    >
      <div class="flex flex-col justify-center items-center mt-3">
        <strong class="uppercase text-yellow mt-5 text-xl"
          >congratulations!</strong
        >
        <span class="uppercase text-white">you get</span>
        <img :src="nft.image" alt="" class="img-nft" />
        <NuxtLink
          class="btn-claim-nft mt-6 flex items-center justify-center"
          :to="`/genesis/${nft.token_id}`"
        >
          <button>Detail</button>
        </NuxtLink>
      </div>
    </Dialog>

    <div class="header h-20"></div>
    <div class="banner">
      <img src="/assets/openbox/banner.png" alt="" />
    </div>

    <div
      class="
        mt-16
        mx-auto
        px-5
        flex
        justify-around
        items-center
        w-11/12
        md:w-9/12
        max-w-xs
        md:max-w-6xl
      "
    >
      <div class="flex flex-col items-center">
        <span
          class="
            mb-4
            text-2xl
            md:text-4xl
            banner-text_header
            bold
            text-border-brown
            uppercase
            text-yellow
          "
          >Countdown</span
        >
        <span
          id="time"
          class="
            text-white text-2xl
            md:text-4xl
            banner-text_header
            bold
            text-border-brown
            uppercase
          "
        ></span>
      </div>
      <div class="flex flex-col items-center">
        <span
          class="
            mb-4
            text-2xl
            md:text-4xl
            banner-text_header
            bold
            text-border-brown
            uppercase
            text-yellow
          "
          >joined</span
        >
        <span
          class="
            text-white text-2xl
            md:text-4xl
            banner-text_header
            bold
            text-border-brown
            uppercase
          "
          >3,888</span
        >
      </div>
    </div>
    <div
      class="
        box-sale
        mx-auto
        px-5
        mt-16
        border-2 border-t-0
        rounded-3xl
        flex flex-col
        w-11/12
        md:w-9/12
        max-w-xs
        md:max-w-6xl
      "
      style="background-color: #3b0e0e; border-color: #581e1e"
    >
      <div class="flex justify-center mb-60">
        <div
          class="
            px-5
            bg-genesis
            flex flex-col
            justify-center
            items-center
            w-11/12
            md:w-9/12
            max-w-xs
            md:max-w-6xl
          "
        >
          <span
            class="
              banner-text_header
              bold
              text-border-brown
              uppercase
              text-yellow text-xl
              md:text-2xl
              mt-5
              mb-2
            "
            >Benefits</span
          >
          <div class="flex max-w-4xl justify-between items-center md:w-1/2">
            <div class="flex flex-col">
              <span class="text-white">Emblem on NFT card </span>
              <span class="text-white">Drop rate 100% </span>
              <span class="text-white">Attributes bonus 5% </span>
            </div>
            <div class="flex flex-col m-2">
              <span class="text-white">Free 1st awaken </span>
              <span class="text-white">Join testnet </span>
              <span class="text-white">Join IDO </span>
            </div>
          </div>
        </div>
      </div>
      <div
        class="
          list-box
          flex flex-col
          md:grid md:grid-cols-2
          xl:grid-cols-3
          place-content-center place-items-center
        "
      >
        <div
          class="box box1 sm:px-4 md:px-3 lg:px-5 relative mb-5"
          v-for="(item, index) in boxes"
          :key="index"
          :class="index === 2 ? ' col-span-2 xl:col-span-1 max-w-sm' : ''"
        >
          <div class="image">
            <img :src="item.image" alt="" />
          </div>
          <span class="txt-amount" v-if="!isLoadingData"> 0 </span>
          <span class="txt-amount" v-else> {{ item.amount }} </span>
          <button
            class="btn-buy-off"
            v-if="isLoadingData"
            @click="buyBox(index)"
          >
            <vue-element-loading
              :active="isBuying"
              spinner="bar-fade-scale"
              color="#FF6700"
            />
            Purchase
          </button>
          <div v-else>
            <button class="btn-buy-off" v-if="amountLeftBoxes[item.name] === 0">
              <vue-element-loading
                :active="isBuying"
                spinner="bar-fade-scale"
                color="#FF6700"
              />
              Sold out
            </button>
            <button
              class="btn-buy flex justify-center items-center"
              @click="buyBox(index)"
              v-else
            >
              <vue-element-loading
                :active="isBuying"
                spinner="bar-fade-scale"
                color="#FF6700"
              />
              <img src="/assets/openbox/BNB.svg" alt="" class="mr-4" />
              {{ formatPriceBN(priceBoxes[item.name]) }} BNB
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full flex justify-center mt-16" v-if="!account">
      <span class="text-2xl text-white">Connect your wallet to join </span>
    </div>
    <div class="w-full flex justify-center mt-16" v-else>
      <span class="text-2xl text-white" v-if="!inWhiteList"
        >You are not in whitelist</span
      >
    </div>
  </div>
</template>

<script>
import { ethers } from "ethers";

const BN = require("bn.js");
import OpenboxGenesis from "@/plugins/contracts/services/OpenboxGenesisV2";
import NFT from "@/plugins/contracts/services/NFT";
import { mapGetters, mapState } from "vuex";
import util from "@/plugins/lib/util";
export default {
  layout: "genesis",
  data() {
    return {
      boxes: [
        {
          name: "Rare",
          boxType: 0,
          image: "/assets/Box1.png",
          amount: 888,
        },
        {
          name: "Legend",
          boxType: 1,
          image: "/assets/Box2.png",
          amount: 88,
        },
        {
          name: "Myth",
          boxType: 2,
          image: "/assets/Box3.png",
          amount: 8,
        },
      ],
      tokenId: undefined,
      boxSelected: undefined,
      isShowNFT: false,
      amountLeftBoxes: undefined,
      priceBoxes: undefined,
      isBuying: false,
      inWhiteList: false,
    };
  },
  async mounted() {
    util.startTimer(
      "time",
      process.env.EVENT_TIME_IDO_WHITELIST,
      "Sale started"
    );

    const web3ModalConnected = localStorage.getItem("web3ModalConnected");
    if (web3ModalConnected) {
      await this.$ether();
    }
    this.amountLeftBoxes = await OpenboxGenesis.getAmountLeftBox(this.provider);
    this.priceBoxes = await OpenboxGenesis.getPriceBox(this.provider);

    if (this.account) {
      let self = this;
      try {
        this.inWhiteList = await OpenboxGenesis.inWhiteList(
          this.provider,
          this.account
        );
      } catch (e) {
        console.log(e);
      }

      const contract = await OpenboxGenesis.getContract(this.provider);
      contract.on("BoxOpened", (buyer, grade, tokenId, sHash) => {
        if (
          buyer.toLowerCase() === this.account.toLowerCase() &&
          sHash === this.hash
        ) {
          this.$toast.add({
            severity: "success",
            summary: "Box is opening",
            life: 3000,
          });
          self.tokenId = tokenId;
        }
      });
    }
  },
  computed: {
    ...mapGetters({
      hash: "openboxGenesis/getHash",
      signature: "openboxGenesis/getSignature",
    }),
    ...mapState("nft", {
      nft: (state) => state.nft,
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
    isLoadingData() {
      //stop event
      return true;
      // return !(
      //   this.amountLeftBoxes !== undefined && this.priceBoxes !== undefined
      // );
    },
  },
  methods: {
    formatPriceBN(price) {
      return parseFloat(ethers.utils.formatEther(price)).toFixed(1);
    },
    async buyBox(boxIndex) {
      if (!this.account) {
        this.$toast.add({
          severity: "warn",
          summary: "You must connect your wallet",
          life: 3000,
        });
        return;
      }
      this.isBuying = true;
      this.boxSelected = this.boxes[boxIndex];
      await this.$store.dispatch("openboxGenesis/getCommit");
    },
  },
  watch: {
    async hash() {
      if (this.boxSelected && this.priceBoxes !== undefined) {
        try {
          const receipt = await OpenboxGenesis.buyBox(
            this.provider,
            this.boxSelected.boxType,
            this.hash,
            this.signature,
            this.priceBoxes[this.boxSelected.name]
          );

          const tx = await receipt.wait();
          if (tx) {
            this.$toast.add({
              severity: "success",
              summary: "Box bought successful",
              life: 1000,
            });
          }
          console.log(tx);
        } catch (e) {
          console.log(e);
          if (e.data) {
            let reason = "";
            if (e.data.message) reason = e.data.message;
            else reason = e.data.data.stack;
            if (
              reason.includes(
                "OpenboxGenesis: either caller is not in the whitelist or public sell is not ready"
              )
            ) {
              this.$toast.add({
                severity: "error",
                summary: "Address is not in whitelist",
                life: 3000,
              });
            }
            if (
              reason.includes("OpenboxGenesis: can only buy at most 2 boxes")
            ) {
              this.$toast.add({
                severity: "error",
                summary: "Can only buy at most 2 boxes",
                life: 3000,
              });
            }
            if (reason.includes("insufficient funds for transfer")) {
              this.$toast.add({
                severity: "error",
                summary: "Balance is not enough",
                life: 3000,
              });
            }

            if (reason.includes("there is no")) {
              this.$toast.add({
                severity: "error",
                summary: "Box sold out",
                life: 3000,
              });
            }
          }
          this.isBuying = false;
        }
      }
    },
    async tokenId() {
      if (this.tokenId !== undefined) {
        // const tokenURI = await NFT.getUriString(
        //   this.provider,
        //   this.tokenId.toNumber()
        // );
        await this.$store.dispatch("nft/fetchAll", this.tokenId.toNumber());
        this.isBuying = false;
        this.isShowNFT = true;
      }
    },
  },
};
</script>
<style lang="scss">
.bg-genesis {
  position: absolute;
  background-image: url("/assets/mobile/genesis/bg_boxsale.png");
  background-size: 100% 100%;
  background-position: center center;
  background-repeat: no-repeat;
  height: 200px;
}

body {
  background-image: url("/assets/Pattern.png"), url("/assets/bg-brown.png") !important;
  background-repeat: repeat;
  background-color: #381d24;
}

.btn-buy {
  position: absolute;
  background-image: url("/assets/openbox/btn_buy_on.png");
  background-size: 100% 100%;
  background-position: center center;
  background-repeat: no-repeat;
  width: 73%;
  height: 60px;
  bottom: 105px;
  left: 13%;
  font-size: 25px;
  color: white;
}
.btn-buy-off {
  position: absolute;
  background-image: url("/assets/openbox/btn_buy_off.png");
  background-size: 100% 100%;
  background-position: center center;
  background-repeat: no-repeat;
  width: 73%;
  height: 60px;
  bottom: 105px;
  left: 13%;
  font-size: 25px;
  color: white;
}

.btn-claim-nft {
  /*position: absolute;*/
  background-image: url("/assets/openbox/btn_claim_nft.png");
  background-size: 100% 100%;
  background-position: center center;
  background-repeat: no-repeat;
  width: 73%;
  height: 62px;
  left: 13%;
  font-size: 25px;
  color: white;
}

.popup-nft {
  background-image: url("/assets/openbox/popup_claimnft.png");
  background-size: 380px 610px;
  background-position: center center;
  background-repeat: no-repeat;
  margin: 0 auto;
}

.txt-amount {
  position: absolute;
  width: 73%;
  height: 50px;
  bottom: 173px;
  left: 13%;
  font-size: 38px;
  color: #61040b;
  text-align: center;
}

.p-dialog-content {
  background: unset !important;
}

.p-dialog-header {
  background: unset !important;
}

.img-nft {
  max-width: 350px;
  max-height: 400px;
}
.icon-bnb {
  position: absolute;
  left: 20px;
}
</style>
