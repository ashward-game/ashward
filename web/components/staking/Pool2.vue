<template>
  <div class="col-md-3 d-flex">
    <div class="el__box">
      <div class="el__box__header text-center">
        <h2 class="el__box__title">{{ pool.name }}</h2>
      </div>
      <div class="el__box__content">
        <div class="el__box__thumb">
          <div class="dnfix__thumb">
            <img src="/assets/images/staking/staking-item-02.png" alt="" />
          </div>
        </div>

        <ul class="el__list">
          <li>
            <div class="li__label">Lock</div>
            <div class="li__value">{{ pool.period }} days</div>
          </li>
          <li>
            <div class="li__label">Joined</div>
            <div class="li__value">
              <CommonNumber
                :to="pool.totalStakers"
                :from="0"
                :delay="2"
                :digital="0"
                easing="Power1.easeOut"
              />
            </div>
          </li>
          <li>
            <div class="li__label">APY</div>
            <div class="li__value">{{ pool.apy }}%</div>
          </li>
          <li>
            <div class="li__label">Stake limit</div>
            <div class="li__value">{{ pool.min }} - {{ pool.max }} ASC</div>
          </li>
          <li>
            <div class="li__label">Pool</div>
            <div class="li__value">
              <CommonNumber
                :to="pool.totalStaked"
                :from="0.0"
                :delay="2"
                :digital="2"
                easing="Power1.easeOut"
              />
              /
              {{ numberWithCommas(pool.pool) }} ASC
            </div>
          </li>
        </ul>

        <button
          class="btn-stake w-100 mb-3"
          data-bs-toggle="modal"
          data-bs-target="#stakeModal"
          @click="setPoolSelected()"
          :disabled="pool.myStakes == pool.max || !checkEndTime"
        >
          Stake
        </button>

        <ul class="el__list">
          <li>
            <div class="li__label">My stakes</div>
            <div class="li__value">
              <CommonNumber
                :to="pool.myStakes"
                :from="0.0"
                :delay="2"
                :digital="2"
                easing="Power1.easeOut"
              />
              ASC
            </div>
          </li>
          <li>
            <div class="li__label">Profit</div>
            <div class="li__value">
              <CommonNumber
                :to="pool.earned"
                :from="0.0"
                :delay="2"
                :digital="5"
                easing="Power1.easeOut"
              />
              ASC
            </div>
          </li>
          <li>
            <div class="li__label">
              <span> Rewards </span>
            </div>
            <div class="li__value" v-if="pool.myStakes && pool.myStakes !== 0">
              <p>{{ parseInt(pool.myStakes / pool.min) }} gift box</p>
              <p>{{ pool.myStakes > 0 ? 1 : 0 }} gift code</p>
            </div>
            <div class="li__value" v-else>
              <p>__ gift box</p>
              <p>__ gift code</p>
            </div>
          </li>
        </ul>

        <button
          class="btn-stake2 w-100"
          @click="withdraw('pool2')"
          :disabled="checkEndTime"
        >
          Harvest
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import stake from "~/mixins/stake";
import util from "@/plugins/lib/util";
const { numberWithCommas } = util;

export default {
  name: "Pool2",
  mixins: [stake],
  data() {
    return {
      pool: {
        name: "Macsen",
        type: "pool2",
        period: 21,
        apy: 912.5,
        min: 1000,
        max: 2000,
        pool: 200000,
      },
    };
  },
  mounted() {
    this.initPool("pool2");
  },
  methods: { numberWithCommas },
};
</script>

<style scoped></style>
