<template>
  <component :is="component" />
</template>

<script>
import Base from '~/components/common/base.vue'
export default {
  layout: ({ app }) => (app.$device.isMobileOrTablet ? 'mobile' : 'default'),
  components: {
    Mobile: () => import('~/components/desktop/404'),
    Desktop: () => import('~/components/desktop/404')
  },
  extends: Base
}
</script>
