<template>
  <div
    class="header-top flex flex-col text-sm items-start px-2 py-3 sm:text-base md:flex-row md:items-center md:text-md md:justify-center md:py-6 xl:py-3 lg:text-lg xl:text-xl 2xl:text-3xl"
  >
    <div class="text-white uppercase">Contract address (BEP-20):</div>
    <div class="text-yellow flex items-center md:ml-2 2xl:ml-4">
      <span ref="addressContract"></span>
      <btn-yellow
        class="uppercase text-white ml-2 text-border-black lg:px-4 lg:py-2 lg:text-lg xl:text-xl 2xl:text-3xl 2xl:px-6 2xl:ml-4"
        @click="copy"
      >
        copy
      </btn-yellow>
    </div>
  </div>
</template>

<script>
import BtnYellow from "./button/BtnYellow.vue";
export default {
  name: "HeaderTop",
  components: { BtnYellow },
  methods: {
    copy() {
      const address = this.$refs.addressContract.innerText;
      navigator.clipboard.writeText(address);
    },
  },
};
</script>

<style lang="scss" scope>
.header-top {
  background-color: #c60217;
  color: #fff;
  position: relative;
}
</style>
