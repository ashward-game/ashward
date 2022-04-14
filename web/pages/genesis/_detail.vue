<template>
  <div class="container mx-auto">
    <div class="header h-20"></div>

    <div
      class="flex flex-col sm:flex-row sm:grid sm:grid-cols-12 w-full mt-20 sm:mt-40"
    >
      <div class="col-span-2 mb-4 px-4 sm:px-0">
        <div class="flex items-center">
          <svg
            width="32"
            height="32"
            viewBox="0 0 32 32"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M27 16H5"
              stroke="#BDBDBD"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
            <path
              d="M14 7L5 16L14 25"
              stroke="#BDBDBD"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
          <NuxtLink to="/genesis" class="single__back">
            <button class="text-white ml-4">back</button>
          </NuxtLink>
        </div>
      </div>
      <div class="col-span-3 flex justify-center mb-4 px-4 sm:px-0">
        <img
          class="w-full sm:w-auto"
          style="max-width: 250px; max-height: 360px"
          :src="nft.image"
          alt=""
        />
      </div>
      <div class="col-span-5 xl:col-span-7 mb-4 px-4 sm:px-0">
        <div class="flex justify-between mb-6">
          <div class="flex flex-col">
            <span class="text-3xl text-white">{{ nft.name }}</span>
            <span class="mt-2 text-gray-300">TokenID: {{ nft.token_id }}</span>
          </div>
          <div class="bg-grade flex justify-center items-start">
            <span class="text-white text-2xl mt-2">{{ nft.rarity }}</span>
          </div>
        </div>
        <TabView>
          <TabPanel>
            <template #header>
              <span class="text-2xl text-yellow"> Information</span>
            </template>
            <div class="flex grid grid-cols-3 my-4">
              <span class="text-gray-300"> Owner </span>
              <span class="text-gray-300"> {{ textAddress(nft.owner) }} </span>
            </div>
            <hr style="border-color: #512b35" />
            <div class="flex grid grid-cols-3 my-4">
              <span class="text-gray-300"> Rarity </span>
              <span class="text-gray-300"> {{ nft.rarity }} </span>
            </div>
            <hr style="border-color: #512b35" />
            <div class="flex grid grid-cols-3 my-4">
              <span class="text-gray-300"> Class </span>
              <span class="text-gray-300"> {{ nft.class }} </span>
            </div>
            <hr style="border-color: #512b35" />
            <div class="flex grid grid-cols-3 my-4">
              <span class="text-gray-300"> Awaken level </span>
              <span class="text-gray-300">
                {{ nft.properties.awaken_level.value }}
              </span>
            </div>
            <div class="my-3">
              <span class="text-2xl text-white"> Attributes </span>
            </div>
            <div class="my-3 flex flex-col">
              <!--              <div class="flex grid grid-cols-4 my-4">-->
              <!--                <div class="flex flex-col">-->
              <!--                  <span class="text-yellow">Overview</span>-->
              <!--                  <span class="text-gray-300">1200</span>-->
              <!--                </div>-->
              <!--              </div>-->
              <div class="flex grid grid-cols-4 my-4">
                <div class="flex flex-col">
                  <span class="text-yellow">Damage Point</span>
                  <span class="text-gray-300">
                    {{ nft.properties.damage.value }}</span
                  >
                </div>
                <div class="flex flex-col">
                  <span class="text-yellow">HP</span>
                  <span class="text-gray-300">
                    {{ nft.properties.hp.value }}</span
                  >
                </div>
                <div class="flex flex-col">
                  <span class="text-yellow">Mana</span>
                  <span class="text-gray-300">{{
                    nft.properties.mana.value
                  }}</span>
                </div>
              </div>

              <div class="flex grid grid-cols-4 my-4">
                <div class="flex flex-col">
                  <span class="text-yellow">Attack speed</span>
                  <span class="text-gray-300">{{
                    nft.properties.atk_speed.value
                  }}</span>
                </div>
                <div class="flex flex-col">
                  <span class="text-yellow">Crit chance</span>
                  <span class="text-gray-300">
                    {{ nft.properties.crit_chance.value }}</span
                  >
                </div>
                <div class="flex flex-col">
                  <span class="text-yellow">Armor</span>
                  <span class="text-gray-300">
                    {{ nft.properties.armor.value }}</span
                  >
                </div>
              </div>

              <!--              <div class="flex my-4">-->
              <!--                <div class="flex flex-col">-->
              <!--                  <span class="text-yellow">Ability</span>-->
              <!--                  <span class="text-gray-300"-->
              <!--                    >{{ nft.properties.abilities.value }}-->
              <!--                  </span>-->
              <!--                </div>-->
              <!--                <div class="flex flex-col">-->
              <!--                  <span class="text-yellow">Ultimate</span>-->
              <!--                  <span class="text-gray-300">-->
              <!--                    {{ nft.properties.ultimate.value }}-->
              <!--                  </span>-->
              <!--                </div>-->
              <!--              </div>-->
            </div>
          </TabPanel>
        </TabView>
      </div>
      <div class="xl:col-span-2"></div>
    </div>
  </div>
</template>

<script>
import { ethers } from "ethers";

const BN = require("bn.js");
import { mapGetters, mapState } from "vuex";

export default {
  layout: "genesis",
  async asyncData({ store, params }) {
    await store.dispatch("nft/fetchAll", params.detail);
  },
  data() {
    return {};
  },
  async mounted() {},
  computed: {
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
  },
  methods: {
    textAddress(text) {
      return text.slice(0, 5) + "..." + text.slice(-5);
    },
  },
};
</script>
<style lang="scss">
@import "~/static/scss/pages/index";
</style>

<style lang="css">
body {
  background-image: url("/assets/Pattern.png"), url("/assets/bg-brown.png");
  background-repeat: repeat;
  background-color: #381d24;
}
.bg-grade {
  background-image: url("/assets/nft/grade_box.png");
  background-size: 100% 100%;
  background-position: center center;
  background-repeat: no-repeat;
  width: 180px;
}
.p-tabview .p-tabview-panels {
  background: #2d171d;
}
.p-tabview .p-tabview-nav {
  background: unset;
  border: 0;
}
.p-tabview .p-tabview-nav li.p-highlight .p-tabview-nav-link {
  background: #2d171d;
}
.p-component {
  font-family: unset;
}
.p-tabview .p-tabview-nav li.p-highlight .p-tabview-nav-link {
  border-color: unset;
  color: unset;
}
</style>
