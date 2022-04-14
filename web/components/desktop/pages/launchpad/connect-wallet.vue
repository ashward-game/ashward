<template>
  <div>
    <div class="el__box__tab ps-4">
      <ul id="myTab" class="nav nav-custom -box" role="tablist">
        <li class="nav-item" role="presentation">
          <button
            id="box1-tab"
            class="nav-link active"
            data-bs-toggle="tab"
            data-bs-target="#box1"
            type="button"
            role="tab"
            aria-controls="box1"
            aria-selected="true"
          >
            Info
          </button>
        </li>
        <li class="nav-item" role="presentation">
          <button
            id="box2-tab"
            class="nav-link"
            data-bs-toggle="tab"
            data-bs-target="#box2"
            type="button"
            role="tab"
            aria-controls="box2"
            aria-selected="false"
          >
            Schedule
          </button>
        </li>
      </ul>
    </div>
    <div class="el__box -tab mb-4">
      <div id="myTabContent" class="tab-content">
        <div
          id="box1"
          class="tab-pane fade show active"
          role="tabpanel"
          aria-labelledby="box1-tab"
        >
          <div class="el__box__time text-center mb-3">
            <p class="text-2xl text-white">Sale finished</p>
          </div>

<!--          <div class="el__box__time text-center mb-3" v-else>-->
            <!--            <div v-if="!isStopped && !isOpenIDO">-->
            <!--              <p>Sale starts in</p>-->
            <!--              <p class="-time">-->
            <!--                <CountdownWhitelist />-->
            <!--              </p>-->
            <!--            </div>-->

            <!--            <div v-else-if="!isStopped && isOpenIDO">-->
            <!--              <p>FCFS starts in</p>-->
            <!--              <p class="-time">-->
            <!--                <CountdownFCFS />-->
            <!--              </p>-->
            <!--            </div>-->

<!--            <div v-if="!isStopped">-->
<!--              <p>FCFS ends in</p>-->
<!--              <p class="-time">-->
<!--                <CountdownEnd />-->
<!--              </p>-->
<!--            </div>-->
<!--          </div>-->
<!--          <p class="mb-3">IDO and distribution on Binance Smart Chain.</p>-->

<!--          <div class="progress__wrap mb-4">-->
<!--            <div class="progress">-->
<!--              <div-->
<!--                class="progress-bar"-->
<!--                role="progressbar"-->
<!--                :style="{-->
<!--                  width: progressStatus + '%',-->
<!--                }"-->
<!--                aria-valuenow="20"-->
<!--                aria-valuemin="0"-->
<!--                aria-valuemax="100"-->
<!--              />-->
<!--            </div>-->
<!--            <div class="d-flex justify-content-between mt-3">-->
<!--              <p>{{ parseFloat(progressStatus).toFixed(2) }}%</p>-->
<!--              <p>-->
<!--                {{ numberWithCommas(amountTotal - amountRemain) }} /-->
<!--                {{ numberWithCommas(amountTotal) }} ASC-->
<!--              </p>-->
<!--            </div>-->
<!--          </div>-->

          <div class="el__wrap px-sm-5">
            <div v-if="isBought && dataBought">
              <span class="text-md text-white text-center">
                You have bought a {{ dataBought.pack }} BUSD pack, which
                corresponding to {{ dataBought.amount }} ASC.
                <br />
                You can claim your ASC tokens starting on March 18, 2022
              </span>
              <a class="text-md text-blue-500" href="/vesting/ido"> here.</a>
            </div>
<!--            <div class="mb-3" v-else>-->
<!--              <p class="mb-3">Choose 1 option to buy</p>-->

<!--              <ul class="row row-col-1 row-cols-sm-2 gx-2 mb-4 pb-2">-->
<!--                <li class="mb-3 mb-sm-0">-->
<!--                  <label :class="['checkmark__item', { '-active': !isBought }]">-->
<!--                    <input-->
<!--                      v-model="packageSelected"-->
<!--                      type="radio"-->
<!--                      name="cate"-->
<!--                      value="0"-->
<!--                    /><span class="checkmark" />-->
<!--                    <span>$100 (3333,34 ASC)</span>-->
<!--                  </label>-->
<!--                </li>-->
<!--                <li>-->
<!--                  <label :class="['checkmark__item', { '-active': !isBought }]">-->
<!--                    <input-->
<!--                      v-model="packageSelected"-->
<!--                      type="radio"-->
<!--                      name="cate"-->
<!--                      value="1"-->
<!--                    /><span class="checkmark" />-->
<!--                    <span>$200 (6666,67 ASC)</span>-->
<!--                  </label>-->
<!--                </li>-->
<!--              </ul>-->
<!--              <div class="text-center" v-if="!account">-->
<!--                <button class="btnz w-100 mb-3" @click="connectWallet()">-->
<!--                  Connect wallet-->
<!--                </button>-->
<!--                <div class="-text-note">Connect wallet to verify</div>-->
<!--              </div>-->
<!--              <div v-else-if="!isStopped">-->
<!--                <div v-if="isBought" class="text-center">-->
<!--                  <button-->
<!--                    class="btnz w-100 mb-3 me-3 js-attr-button"-->
<!--                    disabled=""-->
<!--                  >-->
<!--                    Bought-->
<!--                  </button>-->
<!--                </div>-->
<!--                <div-->
<!--                  v-else-if="-->
<!--                    (isOpenIDO && inWhitelist) || (isOpenFCFS && !inWhitelist)-->
<!--                  "-->
<!--                  class="text-center flex justify-center"-->
<!--                >-->
<!--                  <div v-if="packageSelected === undefined">-->
<!--                    <button class="btnz mb-3 me-2" disabled>Enable</button>-->
<!--                  </div>-->
<!--                  <div class="text-center" v-else>-->
<!--                    <button-->
<!--                      class="btnz mb-3 me-2"-->
<!--                      @click="allow()"-->
<!--                      v-if="!isEnable"-->
<!--                    >-->
<!--                      <vue-element-loading-->
<!--                        :active="isLoadingCall"-->
<!--                        spinner="bar-fade-scale"-->
<!--                        color="#FF6700"-->
<!--                      />Enable-->
<!--                    </button>-->
<!--                  </div>-->
<!--                  <div class="text-center">-->
<!--                    <button-->
<!--                      class="btnz w-100 mb-2 me-2 js-attr-button"-->
<!--                      @click="buy()"-->
<!--                      v-if="isEnable"-->
<!--                    >-->
<!--                      <vue-element-loading-->
<!--                        :active="isLoadingCall"-->
<!--                        spinner="bar-fade-scale"-->
<!--                        color="#FF6700"-->
<!--                      />-->
<!--                      Buy-->
<!--                    </button>-->
<!--                    <button-->
<!--                      class="btnz mb-3 me-2"-->
<!--                      disabled-->
<!--                      v-if="!isEnable || isBought"-->
<!--                    >-->
<!--                      Buy-->
<!--                    </button>-->
<!--                  </div>-->
<!--                </div>-->
<!--              </div>-->
<!--            </div>-->

            <ul class="el__list -light mb-3">
              <li class="d-block d-sm-flex">
                <div class="li__label -text-light">1 BUSD = 33.33 ASC</div>
                <div class="li__val font14 ms-sm-auto text-start text-sm-end">
                  1 ASC = 0.03 BUSD
                </div>
              </li>
              <li>
                <div class="li__label">Whitelist time</div>
                <div class="li__val ms-auto">
                  <p>18-Mar-2022 - 13:00 (UTC)</p>
                  <p>18-Mar-2022 - 14:30 (UTC)</p>
                </div>
              </li>
              <li>
                <div class="li__label">FCFS time</div>
                <div class="li__val ms-auto">
                  <p>18-Mar-2022 - 14:30 (UTC)</p>
                  <p>18-Mar-2022 - 15:00 (UTC)</p>
                </div>
              </li>
            </ul>

<!--            <div v-if="account && !isPublicSale">-->
<!--              <div-->
<!--                class="el__box__time text-center mx-4"-->
<!--                v-if="!inWhitelist && !isStopped"-->
<!--              >-->
<!--                <span class="text-2xl text-white"-->
<!--                  >You are not in whitelist</span-->
<!--                >-->
<!--              </div>-->
<!--              <div-->
<!--                class="el__box__time text-center mx-4"-->
<!--                v-else-if="inWhitelist && !isStopped"-->
<!--              >-->
<!--                <span class="text-2xl text-white">You are in whitelist</span>-->
<!--              </div>-->
<!--            </div>-->
          </div>
        </div>
        <div
          id="box2"
          class="tab-pane fade"
          role="tabpanel"
          aria-labelledby="box2-tab"
        >
          <div class="el__schedule -col3">
            <div class="-th row row-cols-3 d-none d-sm-flex">
              <div>Round</div>
              <div>Opens</div>
              <div>Closes</div>
            </div>
            <div class="-tr row">
              <div class="-label col-12 col-sm-4 mb-2">Register</div>
              <div class="col-6 col-sm-4">26-Feb-2022<br />05:00 (UTC)</div>
              <div class="col-6 col-sm-4">06-Mar-2022<br />15:00 (UTC)</div>
            </div>
            <div class="-tr row row-cols-2 row-cols-sm-3">
              <div class="-label col-12 col-sm-4 mb-2">Whitelist</div>
              <div class="col-6 col-sm-4">18-Mar-2022<br />13:00 (UTC)</div>
              <div class="col-6 col-sm-4">18-Mar-2022<br />14:30 (UTC)</div>
            </div>
            <div class="-tr row row-cols-2 row-cols-sm-3">
              <div class="-label col-12 col-sm-4 mb-2">FCFS</div>
              <div class="col-6 col-sm-4">18-Mar-2022<br />14:30 (UTC)</div>
              <div class="col-6 col-sm-4">18-Mar-2022<br />15:00 (UTC)</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import CountdownWhitelist from "./countdown-whitelist.vue";
import CountdownFCFS from "./countdown-fcfs.vue";
import { ethers } from "ethers";
import BUSD from "@/plugins/contracts/services/BUSD";
import IDO from "@/plugins/contracts/services/IDO";
import util from "@/plugins/lib/util";
import { ADDRESS_IDO } from "~/plugins/contracts/IDO";
import moment from "moment";
import CountdownEnd from "~/components/desktop/pages/launchpad/countdown-end";
const BN = require("bn.js");
export default {
  components: {
    CountdownEnd,
    CountdownWhitelist,
    CountdownFCFS,
  },
  data() {
    return {
      packages: [
        {
          price: 100,
          id: 0,
          amount: "3333,34",
        },
        {
          price: 200,
          id: 1,
          amount: "6666,67",
        },
      ],
      packageSelected: undefined,
      isEnable: false,
      isLoadingCall: false,
      isBought: false,
      inWhitelist: false,
      amountTotal: 3000000,
      amountRemain: 3000000,
      isStopped: false,
      dataBought: undefined,
      isPublicSale: false,
      isOpenIDO: false,
      isOpenFCFS: false,
    };
  },
  async mounted() {
    this.startSaleIn(process.env.EVENT_TIME_IDO_WHITELIST);
    this.startFCFSIn(process.env.EVENT_TIME_IDO_FCFS);
    const web3ModalConnected = localStorage.getItem("web3ModalConnected");

    if (web3ModalConnected) {
      await this.$ether();
    }

    if (this.provider && this.account) {
      this.isPublicSale = await IDO.getIsPublicSale(this.provider);
    }
    if (this.provider) {
      const amountRemainBN = await IDO.getAmountRemain2(this.provider);
      this.amountRemain = parseInt(ethers.utils.formatEther(amountRemainBN));
    }

    if (this.account) {
      try {
        let storeAddress = localStorage.getItem("address");

        let storePack = localStorage.getItem("pack");

        let storeAmount = localStorage.getItem("amount");

        if (storeAddress && storeAddress.includes(this.account)) {
          this.isBought = true;

          const index = JSON.parse(storeAddress).indexOf(this.account);
          const address = JSON.parse(storeAddress);
          const pack = JSON.parse(storePack);
          const amount = JSON.parse(storeAmount);
          this.dataBought = {
            address: address[index],
            pack: pack[index],
            amount: amount[index],
          };
        }

        this.inWhitelist = util.checkInWhiteListIDO(this.account);
        const amountRemainBN = await IDO.getAmountRemain(this.provider);
        this.amountRemain = parseInt(ethers.utils.formatEther(amountRemainBN));

        this.isStopped = await IDO.getPaused(this.provider);
      } catch (e) {
        console.log(e);
      }
    }
  },
  computed: {
    amountText() {
      return new Intl.NumberFormat().format(this.amountRemain);
    },
    progressStatus() {
      return ((this.amountTotal - this.amountRemain) / this.amountTotal) * 100;
    },
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
    async allow() {
      try {
        if (this.packageSelected !== undefined && this.account) {
          this.isLoadingCall = true;
          const priceToWei = await ethers.utils.parseUnits(
            this.packages[this.packageSelected].price.toString(),
            18
          );
          const receipt = await BUSD.approveBUSD(
            this.provider,
            priceToWei.toString()
          );
          const tx = await receipt.wait();
          if (tx) {
            this.isEnable = true;

            this.isLoadingCall = false;
            this.$toast.add({
              severity: "success",
              summary: "Enable successful",
              life: 3000,
            });
          }
          this.isLoadingCall = false;
        }
      } catch (e) {
        console.log(e);
      }
    },
    async buy() {
      if (this.account && this.packageSelected !== undefined) {
        try {
          this.isLoadingCall = true;
          const allowance = await BUSD.getAllowance(
            this.provider,
            this.account,
            ADDRESS_IDO
          );

          if (allowance) {
            const receipt = await IDO.buy(this.provider, this.packageSelected);
            const tx = await receipt.wait();
            if (tx) {
              const amountRemainBN = await IDO.getAmountRemain(this.provider);
              this.amountRemain = parseInt(
                ethers.utils.formatEther(amountRemainBN)
              );

              let storeAddress = localStorage.getItem("address");
              let storePack = localStorage.getItem("pack");
              let storeAmount = localStorage.getItem("amount");

              let storedAddress = [];
              let storedPack = [];
              let storedAmount = [];

              if (storeAddress) {
                storedAddress = JSON.parse(storeAddress);
                storedPack = JSON.parse(storePack);
                storedAmount = JSON.parse(storeAmount);
              }
              storedAddress.push(this.account);
              storedPack.push(
                this.packages[this.packageSelected].price.toString()
              );
              storedAmount.push(this.packages[this.packageSelected].amount);

              localStorage.setItem("address", JSON.stringify(storedAddress));
              localStorage.setItem("pack", JSON.stringify(storedPack));
              localStorage.setItem("amount", JSON.stringify(storedAmount));

              this.isBought = true;
              this.$toast.add({
                severity: "success",
                summary: "Buy successful",
                life: 3000,
              });
            }
          }
        } catch (error) {
          console.log(error);
          if (error.data) {
            let reason = "";
            let message = "";
            if (error.data.message) reason = error.data.message;
            else reason = error.data.data.stack;
            if (
              reason.includes(
                "IDO: either caller is not in the whitelist or public sale is not ready"
              )
            ) {
              message = "You are not in whitelist";
            }

            if (reason.includes("IDO: caller has reached her quota")) {
              message = "Each wallet is eligible to purchase 1 package";
            }

            if (reason.includes("amount exceeds balance")) {
              message = "Balance is not enough";
            }

            if (reason.includes("Pausable: paused")) {
              message = "Sale ended";
            }

            if (message !== "") {
              this.$toast.add({
                severity: "error",
                summary: message,
                life: 3000,
              });
            }
          }
        }
        this.isLoadingCall = false;
      }
    },
    async connectWallet() {
      this.isOpenWallet = false;
      await this.$ether();
    },
    numberWithCommas(x) {
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    },
    startSaleIn(eventTime) {
      let countDownDate = moment(eventTime, "YYYY-MM-DD HH:mm:ss");
      const self = this;
      let x = setInterval(function () {
        let now = moment(moment.utc().format("YYYY-MM-DD HH:mm:ss"));

        let distance = moment.duration(countDownDate.diff(now));

        if (distance < 0) {
          clearInterval(x);

          self.isOpenIDO = true;
          return true;
        }
      }, 1000);
      return false;
    },
    startFCFSIn(eventTime) {
      let countDownDate = moment(eventTime, "YYYY-MM-DD HH:mm:ss");
      const self = this;
      let x = setInterval(function () {
        let now = moment(moment.utc().format("YYYY-MM-DD HH:mm:ss"));

        let distance = moment.duration(countDownDate.diff(now));

        if (distance < 0) {
          clearInterval(x);

          self.isOpenFCFS = true;
          return true;
        }
      }, 1000);
      return false;
    },
  },
  watch: {
    account(newValue, oldValue) {
      let address = localStorage.getItem("address");
      let storedAddress = JSON.parse(address);
      if (storedAddress && storedAddress.includes(newValue)) {
        this.isBought = true;
      }
    },
  },
};
</script>

<style scoped lang="scss" />
