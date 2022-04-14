<template>
  <component :is="component" />
</template>

<script>
export default {
  scrollToTop: true,
  layout: ({ app }) => (app.$device.isMobileOrTablet ? 'default' : 'default'),
  beforeCreate () {
    this.component = this.$device.isMobileOrTablet ? 'Desktop' : 'Desktop'
  },
  head () {
    return {
      bodyAttrs: {
        class: this.$device.isMobileOrTablet
          ? 'mobile swal2-toast-shown'
          : 'desktop swal2-toast-shown',
        ref: 'body'
      }
    }
  }
}
</script>
