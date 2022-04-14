<template>
  <div class="el__box">
    <p class="font18 mb-2">Information</p>
    <ul class="el__list mb-4 pb-3">
      <li>
        <div class="li__label">IDO price</div>
        <div class="li__val ms-auto">1 ASC = 0.03 USD</div>
      </li>
      <li>
        <div class="li__label">Hard cap</div>
        <div class="li__val ms-auto">$90,000</div>
      </li>
      <li>
        <div class="li__label">Initial supply</div>
        <div class="li__val ms-auto">28,600,000 ASC</div>
      </li>
    </ul>

    <p class="font18 mb-2">Token info</p>
    <ul class="el__list -s2 mb-4 pb-3">
      <li>
        <div class="li__label">Address</div>
        <div class="li__val ms-auto">
          {{ ascToken }}
        </div>
      </li>
      <li>
        <div class="li__label">Token name</div>
        <div class="li__val ms-auto">Ashward</div>
      </li>
      <li>
        <div class="li__label">Token symbol</div>
        <div class="li__val ms-auto">ASC</div>
      </li>
      <li>
        <div class="li__label">Type</div>
        <div class="li__val ms-auto">BEP-20</div>
      </li>
      <li>
        <div class="li__label">Token decimals</div>
        <div class="li__val ms-auto">18</div>
      </li>
      <li>
        <div class="li__label">Total supply</div>
        <div class="li__val ms-auto">1,000,000,000 ASC</div>
      </li>
      <li>
        <div class="li__label">Listing</div>
        <div class="li__val ms-auto">PancakeSwap</div>
      </li>
      <li class="pt-4">
        <div class="li__label">Vesting</div>
        <div class="li__val ms-auto">
          TGE 20%, 1 month cliff, <br />vest monthly in 5 months
        </div>
      </li>
    </ul>

    <p class="font18 mb-3">About</p>
    <p>
      ASC is an essential part of The Ashward platform. It is an BEP-20 utility
      token built on the Binance Smart Chain (BSC) that serves as the basis for
      transactions within Ashward, and has the two main uses: Access Ashward
      platform and staking. ASC has a fixed supply of 1,000,000,000 (one
      billion) tokens.
    </p>
  </div>
</template>
<script>
import { ADDRESS_TOKEN } from "~/plugins/contracts/Token";

export default {
  components: {},
  props: {},
  data() {
    return {
      ascToken: "",
    };
  },
  mounted() {
    this.ascToken = ADDRESS_TOKEN;
  },
  computed: {},
  methods: {},
};
</script>

<style scoped lang="scss" />
