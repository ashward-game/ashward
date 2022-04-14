<template>
  <section class="home-roadmap">
    <div class="container container-1200">
      <header class="sc__header text-center">
        <h2 class="sc__header__title wow fadeInUp">ROADMAP</h2>
      </header>
      <div class="row pt-2 pt-md-5 gx-5">
        <div class="col-md-7 col-lg-8 wow fadeInLeft">
          <div class="nav-roadmap__wrap">
            <ul>
              <client-only>
                <flickity
                  id="roadmapTab"
                  ref="flickityRoadmap"
                  class="nav nav-roadmap row"
                  role="tablist"
                  :options="options"
                >
                  <li class="nav-item col-12 col-md-6" role="" @click="component = 'roadmap1'">
                    <button
                      id="roadmap1-tab"
                      :class="['nav-link', { 'tab-active': component === 'roadmap1'}]"
                      data-bs-toggle="tab"
                      data-bs-target="#roadmap1"
                      type="button"
                      role="tab"
                      aria-controls="roadmap1"
                      aria-selected="true"
                    >
                      <div class="dnfix__thumb -contain">
                        <img src="/assets/images/roadmap-1.png" alt="">
                      </div>
                    </button>
                  </li>
                  <li class="nav-item col-12 col-md-6" role="" @click="component = 'roadmap2'">
                    <button
                      id="roadmap2-tab"
                      :class="['nav-link', { 'tab-active': component === 'roadmap2'}]"
                      data-bs-toggle="tab"
                      data-bs-target="#roadmap2"
                      type="button"
                      role="tab"
                      aria-controls="roadmap2"
                      aria-selected="true"
                    >
                      <div class="dnfix__thumb -contain">
                        <img src="/assets/images/roadmap-2.png" alt="">
                      </div>
                    </button>
                  </li>
                  <li class="nav-item col-12 col-md-6" role="" @click="component = 'roadmap3'">
                    <button
                      id="roadmap3-tab"
                      :class="['nav-link', { 'tab-active': component === 'roadmap3'}]"
                      data-bs-toggle="tab"
                      data-bs-target="#roadmap3"
                      type="button"
                      role="tab"
                      aria-controls="roadmap3"
                      aria-selected="true"
                    >
                      <div class="dnfix__thumb -contain">
                        <img src="/assets/images/roadmap-3.png" alt="">
                      </div>
                    </button>
                  </li>
                  <li class="nav-item col-12 col-md-6" role="" @click="component = 'roadmap4'">
                    <button
                      id="roadmap4-tab"
                      :class="['nav-link', { 'tab-active': component === 'roadmap4'}]"
                      data-bs-toggle="tab"
                      data-bs-target="#roadmap4"
                      type="button"
                      role="tab"
                      aria-controls="roadmap4"
                      aria-selected="true"
                    >
                      <div class="dnfix__thumb -contain">
                        <img src="/assets/images/roadmap-4.png" alt="">
                      </div>
                    </button>
                  </li>
                </flickity>
              </client-only>
            </ul>
          </div>
        </div>
        <div class="col-md-5 col-lg-4 wow fadeInRight">
          <div id="roadmapTabContent" class="tab-content">
            <div
              id="roadmap1"
              class="tab-pane fade show active"
              role="tabpanel"
              aria-labelledby="roadmap1-tab"
            >
              <div class="el__title">Q4, 2021</div>
              <div class="el__excerpt">
                Build core team<br>Complete project concept<br>Build
                community
              </div>
            </div>
            <div
              id="roadmap2"
              class="tab-pane fade"
              role="tabpanel"
              aria-labelledby="roadmap2-tab"
            >
              <div class="el__title">Q1, 2022</div>
              <div class="el__excerpt">
                Private community round<br>Genesis NFT sale<br>Public sale &
                Listing<br>PancakeSwap listing<br>Game testnet
              </div>
            </div>
            <div
              id="roadmap3"
              class="tab-pane fade"
              role="tabpanel"
              aria-labelledby="roadmap3-tab"
            >
              <div class="el__title">Q2, 2022</div>
              <div class="el__excerpt">
                Launch Marketplace<br>Update PvE<br>Add characters<br>Hero
                awakening<br>Game mainet V1.1
              </div>
            </div>
            <div
              id="roadmap4"
              class="tab-pane fade"
              role="tabpanel"
              aria-labelledby="roadmap4-tab"
            >
              <div class="el__title">Q3, 2022</div>
              <div class="el__excerpt">New items<br>Test PvP</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script>
import $ from 'jquery'

export default {
  components: {},
  props: {
  },
  data () {
    return {
      slideIndex: 0,
      component: 'roadmap1',
      options: {
        lazyLoad: true,
        autoPlay: false,
        pauseAutoPlayOnHover: false,
        wrapAround: false,
        initialIndex: 0,
        prevNextButtons: false,
        pageDots: true,
        watchCSS: true,
        cellSelector: '.nav-item',
        on: {
          change: (index) => {
            const index2 = index + 1
            $('.nav-roadmap li:nth-child(' + index2 + ') .nav-link').click()
          }
        }
      }
    }
  },
  computed: {},
  methods: {

  }
}
</script>
