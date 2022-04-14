<template>
  <div class="col-md-3 col--left d-flex">
    <div class="account__sidebar w-100">
      <div class="el__user text-center mb-3">
        <div class="el__user__avatar">
          <div class="dnfix__thumb -contain">
            <img src="assets/images/market/avatar.png" alt="">
          </div>
          <div class="el__user--edit"><i class="icon-edit" /></div>
        </div>
        <div class="el__user__title">Sluthxie</div>
        <div class="el__user__address">
          <span class="text__truncate">{{ address }}}</span>
          <i class="icon-copy ms-1" @click="copy" />
        </div>
      </div>

      <div class="account__nav">
        <ul class="row">
          <li class="col-6 col-md-12">
            <n-link to="/wallet">Information</n-link>
          </li>
          <li class="col-6 col-md-12">
            <n-link to="/listing">Listing</n-link>
          </li>
          <li class="col-6 col-md-12"><n-link to="/account">Settings</n-link></li>
        </ul>
      </div>
    </div>
  </div>
</template>
<script>
import $ from 'jquery'

export default {
  data () {
    return {
      address: '0xaaaf...2cbcdc'
    }
  },
  methods: {
    copy () {
      const $temp = $('<input>')
      $('body').append($temp)
      $temp.val(this.address).select()
      document.execCommand('copy')
    }
  }
}
</script>
