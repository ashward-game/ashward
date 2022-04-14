import $ from 'jquery'

export default {
  methods: {
    toggleClassFilter () {
      $('.market__sidebar').toggleClass('active')
      $('body').toggleClass('filter-open')
    }
  }
}
