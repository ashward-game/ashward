/* eslint-disable */
import Vue from 'vue'
import * as spine from '@esotericsoftware/spine-player'
import '@esotericsoftware/spine-player/dist/spine-player.css'

// Register Components
Vue.use(spine, 'spine')

export default (context, inject) => {
  const initSpine = (name, orbitJson, atlasUrl, animation) => {
    new spine.SpinePlayer(name, {
      jsonUrl: orbitJson,
      atlasUrl: atlasUrl,
      alpha: true,
      showControls: false,
      backgroundColor: "#00000000",
      animation: animation,
    });
  }
  // Inject $hello(msg) in Vue, context and store.
  inject('initSpine', initSpine)
  // For Nuxt <= 2.12, also add 👇
  context.$initSpine = initSpine
}
