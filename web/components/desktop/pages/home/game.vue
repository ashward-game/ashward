<template>
  <section class="home-game">
    <div class="container-full">
      <header class="sc__header text-center">
        <h2 class="sc__header__title wow fadeInUp">GAME PLAY</h2>
      </header>
      <div class="game-slider__wrap wow fadeInUp">
        <client-only>
          <flickity ref="flickityRoadmap" :options="options">
            <div class="slider__item__wrap col-12 col-md-46">
              <div class="slider__item">
                <div class="slider__item__thumb">
                  <div class="dnfix__thumb">
                    <img src="/assets/images/gameplay2.jpg" alt="" />
                  </div>
                </div>
                <div class="slider__item__meta text-center">
                  <p class="slider__item__title">Casual</p>
                  <p class="slider__item__excerpt">
                    Join the most exciting team adventure, <br />get rewards
                    from the game mode
                  </p>
                </div>
              </div>
            </div>
            <div class="slider__item__wrap col-12 col-md-46">
              <div class="slider__item">
                <div class="slider__item__thumb">
                  <div class="dnfix__thumb">
                    <img src="/assets/images/gameplay3.jpg" alt="" />
                  </div>
                </div>
                <div class="slider__item__meta text-center">
                  <p class="slider__item__title">Raid</p>
                  <p class="slider__item__excerpt">
                    Players in a clan join boss fight together.
                  </p>
                </div>
              </div>
            </div>
            <div class="slider__item__wrap col-12 col-md-46">
              <div class="slider__item">
                <div class="slider__item__thumb">
                  <div class="dnfix__thumb">
                    <img src="/assets/images/gameplay1.jpg" alt="" />
                  </div>
                </div>
                <div class="slider__item__meta text-center">
                  <p class="slider__item__title">ADVENTURE</p>
                  <p class="slider__item__excerpt">
                    Join the most exciting team adventure, <br />get rewards
                    from the game mode
                  </p>
                </div>
              </div>
            </div>
          </flickity>
        </client-only>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  props: {
    data: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      options: {
        lazyLoad: true,
        autoPlay: false,
        pauseAutoPlayOnHover: false,
        wrapAround: true,
        initialIndex: 2,
        prevNextButtons: false,
        pageDots: true,
        cellSelector: ".slider__item__wrap",
      },
    };
  },
  computed: {},
  methods: {},
};
</script>

<style scoped lang="scss" />
