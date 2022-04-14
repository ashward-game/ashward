<template>
  <div>
    <Loading
      v-if="$store.state.isLoading"
    />
    <TheHeader />
    <Toast />
    <div class="wrapper">
      <nuxt />
    </div>
  </div>
</template>
<script>
import TheHeader from '@/components/desktop/header/market'
import globalMixin from '@/mixins/global'
import Loading from './Loading'

export default {
  components: {
    TheHeader,
    Loading
  },
  mixins: [globalMixin],
  created () {
    this.$_globalMixin_loading(1000)
  },
  mounted () {
    const bodyElement = document.querySelector('body')
    this.$nextTick(() => {
      setTimeout(() => bodyElement.classList.remove('modal-open'), 1000)
    })
    setTimeout(() => {
      this.$_globalMixin_load()
    }, 1000)
  }
}
</script>
