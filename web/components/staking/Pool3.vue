<template>
  <div class="col-md-3 d-flex">
    <div class="el__box">
      <div class="el__box__header text-center">
        <h2 class="el__box__title">{{ pool.name }}</h2>
      </div>
      <div class="el__box__content">
        <div class="el__box__thumb">
          <div class="dnfix__thumb">
            <img src="/assets/images/staking/staking-item-03.png" alt="" />
          </div>
        </div>
        <div class="el__box__coming">
          <p>Coming soon</p>
        </div>

        <!--        <ul class="el__list">-->
        <!--          <li>-->
        <!--            <div class="li__label">Lock</div>-->
        <!--            <div class="li__value">{{ pool.period }} days</div>-->
        <!--          </li>-->
        <!--          <li>-->
        <!--            <div class="li__label">Joined</div>-->
        <!--            <div class="li__value">{{ totalStakers }}</div>-->
        <!--          </li>-->
        <!--          <li>-->
        <!--            <div class="li__label">APY</div>-->
        <!--            <div class="li__value">{{ pool.apy }}%</div>-->
        <!--          </li>-->
        <!--          <li>-->
        <!--            <div class="li__label">Stake limit</div>-->
        <!--            <div class="li__value">{{ pool.min }} - {{ pool.max }} ASC</div>-->
        <!--          </li>-->
        <!--          <li>-->
        <!--            <div class="li__label">Pool</div>-->
        <!--            <div class="li__value">-->
        <!--              {{ numberWithCommas(totalStaked) }} / {{ numberWithCommas(pool.pool) }} ASC-->
        <!--            </div>-->
        <!--          </li>-->
        <!--        </ul>-->

        <!--        <button-->
        <!--          class="btn-stake w-100 mb-3"-->
        <!--          data-bs-toggle="modal"-->
        <!--          data-bs-target="#stakeModal"-->
        <!--          @click="setPoolSelected()"-->
        <!--        >-->
        <!--          Stake-->
        <!--        </button>-->

        <!--        <ul class="el__list">-->
        <!--          <li>-->
        <!--            <div class="li__label">My stakes</div>-->
        <!--            <div class="li__value">{{ myStakes }} ASC</div>-->
        <!--          </li>-->
        <!--          <li>-->
        <!--            <div class="li__label">Profit</div>-->
        <!--            <div class="li__value">{{ earned }} ASC</div>-->
        <!--          </li>-->
        <!--          <li>-->
        <!--            <div class="li__label">-->
        <!--              <span> Rewards </span>-->
        <!--            </div>-->
        <!--            <div class="li__value" v-if="myStakes && myStakes !== 0">-->
        <!--              <p>{{ myStakes % pool.min }} gift box</p>-->
        <!--              <p>{{ myStakes % pool.min }} gift code</p>-->
        <!--            </div>-->
        <!--            <div class="li__value" v-else>-->
        <!--              <p>__ gift box</p>-->
        <!--              <p>__ gift code</p>-->
        <!--            </div>-->
        <!--          </li>-->
        <!--        </ul>-->

        <!--        <button class="btn-stake2 w-100" @click="withdraw('pool3')">-->
        <!--          Harvest-->
        <!--        </button>-->
      </div>
    </div>
  </div>
</template>

<script>
import stake from "~/mixins/stake";
import util from "@/plugins/lib/util";
const { numberWithCommas } = util;

export default {
  mixins: [stake],
  name: "Pool3",
  data() {
    return {
      pool: {
        name: "Hefina",
        type: "pool3",
        period: 7,
        apy: 365,
        min: 200,
        max: 1000,
        pool: 500000,
      },
    };
  },
  mounted() {},
  methods: { numberWithCommas },
};
</script>

<style scoped></style>
