<template>
  <div>
    <!-- Modal Stake-->
    <div
      id="stakeModal"
      ref="modalStake"
      class="modal -stake fade"
      tabindex="-1"
      aria-hidden="true"
      data-backdrop="static"
      data-keyboard="false"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header justify-content-center">
            <h5 class="modal-title">{{ poolSelected.name }}</h5>
          </div>
          <div class="modal-body">
            <div class="d-flex mb-2">
              <div>Amount</div>
              <div class="ms-auto">
                Available: {{ numberWithCommas(balance) }} ASC
              </div>
            </div>
            <div class="el__input d-flex align-items-center mb-3">
              <input
                type="text"
                class="form-control"
                placeholder="Enter your amount"
                v-model="amountStake"
              />
              <div class="el__input--unit">ASC</div>
              <div class="el__input--prefix cursor-pointer" @click="maxStake()">
                Max
              </div>
            </div>

            <div
              class="alert alert-danger mb-3"
              role="alert"
              v-if="
                amountStake !== 0 &&
                amountStake != '' &&
                (amountStake > poolSelected.max ||
                  amountStake < poolSelected.min)
              "
            >
              <strong>Oops! Something went wrong</strong>
              <p>
                Staking amount: {{ poolSelected.min }} - {{ poolSelected.max }}
              </p>
            </div>

            <ul class="el__list">
              <li>
                <div class="li__label">APY</div>
                <div class="li__value">{{ poolSelected.apy }}%</div>
              </li>
            </ul>

            <hr />

            <div class="mb-3">
              <div class="el__text">
                <p class="text-white">Conditions & Benefits:</p>
                <p>
                  {{ formatDate(poolSelected.startDate) }}
                  to {{ formatDate(poolSelected.endDate) }}
                </p>
                <p>
                  User's limit: {{ poolSelected.min }} -
                  {{ poolSelected.max }} ASC
                </p>
                <p>Testnet access</p>
                <p>Gift box (1 per {{ poolSelected.min }} ASC)</p>
                <p>Gift code (1)</p>
              </div>
            </div>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn-stake2 w-100 mw-0"
              data-bs-dismiss="modal"
            >
              Close
            </button>
            <button
              type="button"
              class="btn-stake w-100 mw-0"
              data-bs-target="#stake2Modal"
              @click="stake(poolSelected.type)"
            >
              <vue-element-loading
                :active="isLoading"
                spinner="bar-fade-scale"
                color="#FF6700"
              />
              Confirm
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal Stake Step 2-->
    <div
      id="stake2Modal"
      class="modal -stake fade"
      tabindex="-1"
      aria-hidden="true"
      ref="modalStakeConfirm"
    >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-stt text-center my-3">
            <i class="icon-check-circle" />
            <p class="modal-stt__text">Completed</p>
            <h5 class="-title text-white">{{ poolSelected.name }}</h5>
          </div>

          <div class="modal-body">
            <ul class="el__list">
              <!-- <li>
                <div class="li__label">Earned</div>
                <div class="li__value">
                  {{ this.amountStake % 500 }} gift box
                </div>
              </li> -->
              <li>
                <div class="li__label">Added</div>
                <div class="li__value">{{ this.amountStake }} ASC</div>
              </li>
              <!-- <li>
                <div class="li__label">Total</div>
                <div class="li__value">1,000 ASC</div>
              </li>
              <li>
                <div class="li__label">Redemption Date</div>
                <div class="li__value">2022-03-26 15:48</div>
              </li> -->
            </ul>
          </div>
          <div class="modal-actions d-flex justify-content-center gap-3">
            <button
              type="button"
              class="btn-stake w-100 mw-0"
              data-bs-dismiss="modal"
            >
              OK
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { ethers } from "ethers";
import Token from "~/plugins/contracts/services/Token";
import util from "@/plugins/lib/util";
import moment from "moment";
const { numberWithCommas } = util;
export default {
  name: "ModalStake",

  data() {
    return {
      balance: 0,
      isLoading: false,
      amountStake: "",
      modalStake: null,
      modalStakeConfirm: null,
    };
  },
  async mounted() {
    this.modalStake = this.getBootstrapModal(this.$refs.modalStake);
    this.modalStakeConfirm = this.getBootstrapModal(
      this.$refs.modalStakeConfirm
    );
  },
  computed: {
    ...mapState("staking", {
      poolSelected: (state) => state.poolSelected ?? {},
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
    numberWithCommas,
    formatDate(date) {
      return moment.unix(date).format("DD-MMM-YY");
    },
    async getBalance() {
      try {
        this.balance = await Token.balanceOf(this.provider, this.account);
      } catch (e) {
        console.error(e);
      }
    },
    maxStake() {
      if (parseFloat(this.balance) >= parseFloat(this.poolSelected.max))
        this.amountStake = this.poolSelected.max;
      else this.amountStake = parseFloat(this.balance);
    },
    async stake() {
      if (!this.account) {
        this.$toast.add({
          severity: "warn",
          summary: "You must connect your wallet",
          life: 3000,
        });

        return;
      }
      if (this.amountStake === "" || parseFloat(this.amountStake) <= 0) {
        this.$toast.add({
          severity: "warn",
          summary: "You must enter amount to stake",
          life: 3000,
        });
        return;
      }
      try {
        this.isLoading = true;
        const amountStakeToWei = await ethers.utils.parseUnits(
          this.amountStake.toString(),
          18
        );

        const receipt = await Token.stake(
          this.provider,
          amountStakeToWei,
          this.poolSelected.type
        );
        const tx = await receipt.wait();

        if (tx) {
          this.isLoading = false;
          this.modalStake.hide();
          this.modalStakeConfirm.show();
          this.$toast.add({
            severity: "success",
            summary: "Staked successful",
            life: 3000,
          });
        }
      } catch (error) {
        if (error.data) {
          console.log(error);
          let reason = "";
          let message = "Oops! Something went wrong";
          if (error.data.message) reason = error.data.message;
          else reason = error.data.data.stack;
          if (
            reason.includes(
              "StakingRewards: min stake must be less than or equal max stake"
            )
          ) {
            message = "Min stake must be less than or equal max stake";
          }

          if (
            reason.includes("StakingRewards: Cannot stake greater than min")
          ) {
            message = "Cannot stake greater than min";
          }

          if (
            reason.includes(
              "StakingRewards: Cannot stake exceed amount of staker"
            )
          ) {
            message = "Cannot stake exceed amount of staker";
          }

          if (reason.includes("StakingRewards: pool is full")) {
            message = "Pool is full";
          }

          if (reason.includes("StakingRewards: end time stake")) {
            message = "End time stake";
          }

          if (reason.includes("StakingRewards: Allowance is not enough")) {
            message = "Allowance is not enough";
          }
          this.$toast.add({
            severity: "error",
            summary: message,
            life: 3000,
          });
        }
      }
      this.isLoading = false;
    },
  },
  watch: {
    async account() {
      await this.getBalance();
    },
  },
};
</script>

<style scoped></style>
