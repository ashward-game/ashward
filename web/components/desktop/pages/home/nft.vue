<template>
  <section class="home-nft">
    <div class="container container-900">
      <header class="sc__header text-center">
        <h2 class="sc__header__title wow fadeInUp">NFT</h2>
      </header>
      <ul
        id="nftTab"
        class="nav nav-character justify-content-center wow fadeInUp"
        role="tablist"
      >
        <li class="nav-item" role="presentation">
          <button
            id="nftTab1-tab"
            :class="['nav-link', { 'active': component === 'NftTab1' }]"
            data-bs-toggle="tab"
            data-bs-target="#nftTab1"
            type="button"
            role="tab"
            aria-controls="nftTab1"
            aria-selected="true"
            @click="component = 'NftTab1'"
          >
            HEROES
          </button>
        </li>
        <li class="nav-item" role="presentation">
          <button
            id="nftTab2-tab"
            :class="['nav-link', { 'active': component === 'NftTab2' }]"
            data-bs-toggle="tab"
            data-bs-target="#nftTab2"
            type="button"
            role="tab"
            aria-controls="nftTab2"
            aria-selected="false"
            @click="component = 'NftTab2'"
          >
            PETS
          </button>
        </li>
      </ul>
      <div id="nftTabContent" class="tab-content wow fadeInUp">
        <client-only>
          <component
            :is="component"
          />
        </client-only>
      </div>
    </div>
  </section>
</template>
<script>
import NftTab1 from './nft-tab1.vue'
import NftTab2 from './nft-tab2.vue'
export default {
  components: {
    NftTab1,
    NftTab2
  },
  props: {
    data: {
      type: Array,
      default: () => []
    }
  },
  data () {
    return {
      component: 'NftTab1'
    }
  },
  computed: {},
  methods: {}
}
</script>

<style scoped lang="scss" />
