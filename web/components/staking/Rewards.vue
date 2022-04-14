<template>
  <div
    id="benefitsModal"
    class="modal -stake fade"
    tabindex="-1"
    aria-hidden="true"
  >
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-header justify-content-center">
          <h5 class="modal-title">Benefits</h5>
        </div>
        <div class="modal-body">
          <div class="el__item">
            <p class="el__item__title">
              <!--              Gift boxes: {{ rewards.giftboxes.length }}-->
              Gift boxes: coming soon
            </p>
            <div class="el__item__content">
              <ul>
                <li v-for="(item, index) in rewards.giftboxes" :key="index">
                  {{ item.code }}
                </li>
              </ul>
            </div>
          </div>
          <div class="el__item">
            <p class="el__item__title">
              <!--              Gift codes: {{ rewards.giftcodes.length }}-->
              Gift codes: coming soon
            </p>
            <div class="el__item__content">
              <ul>
                <li v-for="(item, index) in rewards.giftcodes" :key="index">
                  {{ item.code }}
                </li>
              </ul>
            </div>
          </div>
        </div>
        <div class="modal-actions d-flex justify-content-center gap-3">
          <button type="button" class="btn-stake" data-bs-dismiss="modal">
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { ethers } from "ethers";

export default {
  name: "Rewards",
  computed: {
    ...mapState("staking", {
      rewards: (state) => state.rewards,
    }),
    provider() {
      if (this.$store.state.ether?.isConnectedWeb3) {
        return this.$web3Provider();
      } else {
        return new ethers.providers.JsonRpcProvider(process.env.RPC_PROVIDER);
      }
    },
    account() {
      if (this.$store.state.ether?.web3Account)
        return this.$store.state.ether?.web3Account;
    },
  },
  mounted() {
    if (this.account) {
      console.log("fetchRewards");
      // this.$store.dispatch("staking/fetchRewards", this.account);
    }
  },
  watch: {
    account() {
      // this.$store.dispatch("staking/fetchRewards", this.account);
    },
  },
};
</script>

<style scoped></style>
