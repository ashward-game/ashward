import Vue from 'vue'
import WOW from 'wow.js'

const wowAnime = () => {
  new WOW({
    boxClass: 'wow',
    animateClass: 'animated',
    offset: 0,
    live: true
  }).init()
}

Vue.use(wowAnime)
