<template>
  <section class="home-team">
    <div class="container container-800">
      <header class="sc__header text-center">
        <h2 class="sc__header__title wow fadeInUp">MEET THE TEAM</h2>
      </header>
      <div class="wow fadeInUp">
        <flickity ref="flickityTeam" :options="options">
          <div class="slider__item__wrap w-100">
            <div class="row">
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <a
                        href="https://www.linkedin.com/in/andyphi/"
                        target="_blank"
                      >
                        <img src="/assets/images/team/Avatar_hai.png" alt="" />
                      </a>
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Andy Phi</h3>
                    <div class="el__item__excerpt">Founder & CEO</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/avatar_Huy.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Huy Pham</h3>
                    <div class="el__item__excerpt">CPO</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/Avatar_Tung.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Drimagine</h3>
                    <div class="el__item__excerpt">Head of Design</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/Avatar_Nam.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Nam Ngo</h3>
                    <div class="el__item__excerpt">Back-end Leader</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/AVatar_Lien.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Mei Lin</h3>
                    <div class="el__item__excerpt">CMO</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/AvatarDev13.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Dat Nguyen</h3>
                    <div class="el__item__excerpt">Game dev leader</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="slider__item__wrap w-100">
            <div class="row">
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/AvatarDev17.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Danny</h3>
                    <div class="el__item__excerpt">Community manager</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/AvatarDev15.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Duc Ngo</h3>
                    <div class="el__item__excerpt">Front-end leader</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/AvatarDev5.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Long Vu</h3>
                    <div class="el__item__excerpt">UX/UI Designer</div>
                  </div>
                </div>
              </div>
              <div class="col-6 col-md-4">
                <div class="el__item">
                  <div class="el__item__thumb">
                    <div class="dnfix__thumb -contain">
                      <img src="/assets/images/team/AvatarDev14.png" alt="" />
                    </div>
                  </div>
                  <div class="el__item__meta">
                    <h3 class="el__item__title">Minh Anh</h3>
                    <div class="el__item__excerpt">Artist</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </flickity>
      </div>
    </div>
  </section>
</template>
<script>
export default {
  components: {},
  props: {
    data: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      options: {
        // initialIndex: 4,
        groupCells: true,
        lazyLoad: true,
        pauseAutoPlayOnHover: false,
        wrapAround: true,
        cellAlign: "left",
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
