<template>
  <input :value="value">
</template>

<script>
import $ from 'jquery'

export default {
  // eslint-disable-next-line vue/require-prop-types
  props: ['options', 'value'],
  watch: {
    options (value) {
      this.$slider.update(value)
    }
  },
  mounted () {
    const vm = this

    $(this.$el)
      .ionRangeSlider(this.options)
      // emit event on change.
      .on('change', function () {
        vm.$emit('input', this.value)
      })

    this.$slider = $(this.$el).data('ionRangeSlider')
  },
  destroyed () {
    this.$slider.destroy()
  }
}
</script>
