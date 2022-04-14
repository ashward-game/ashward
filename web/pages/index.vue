<template>
  <div class="wrapper">
    <section class="home-header">
      <div class="ratio -center">
        <client-only>
          <LazyVideo
            src="/assets/images/banner.webm"
            :attrs="{ loop: true, muted: true, autoplay: true }"
          />
        </client-only>
      </div>
      <div class="el__wrap">
        <div class="container">
          <div class="el__meta">
            <div class="el__title wow fadeInUp">Ashward</div>
            <div class="el__excerpt text-center wow fadeInUp">
              A medieval theme game that ultilizes a play-and-earn model,
              pixelated.
            </div>

            <div class="btn__wrap d-flex align-items-center">
              <a
                @click="coming()"
                href=""
                class="btn mb-3 me-sm-3 wow fadeInLeft"
                data-bs-toggle="modal"
                data-bs-target="#stakingModal"
                >Game demo</a
              >
              <a
                href="https://docsend.com/view/kufgsi3qexnzpi3z"
                target="_blank"
                class="btn mb-3 -sepia wow fadeInRight"
                >Pitch deck</a
              >
            </div>
          </div>
        </div>
        <div class="el__scroll wow fadeInUp">Scroll down</div>
      </div>

      <div class="stars" />
      <div class="stars2" />
      <div class="stars3" />
    </section>
    <Nft />
    <Game />
    <Clan />
    <Roadmap />
    <client-only>
      <Team />
    </client-only>

    <section class="home-team">
      <div class="container container-800">
        <header class="sc__header text-center">
          <h2 class="sc__header__title wow fadeInUp">Advisor</h2>
        </header>
        <div class="wow fadeInUp">
          <div class="row justify-content-center">
            <div class="col-6 col-md-4">
              <div class="el__item">
                <div class="el__item__thumb">
                  <div class="dnfix__thumb -contain">
                    <a
                      href="https://www.linkedin.com/in/hiepnp"
                      target="_blank"
                    >
                      <img src="/assets/images/team/Avatar_Hiep.png" alt="" />
                    </a>
                  </div>
                </div>
                <div class="el__item__meta">
                  <h3 class="el__item__title">Stephen Nguyen</h3>
                  <div class="el__item__excerpt">Technology advisor</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <Partners />
  </div>
  <!-- .wrapper -->
</template>

<script>
import Nft from "~/components/desktop/pages/home/nft";
import Game from "~/components/desktop/pages/home/game";
import Clan from "~/components/desktop/pages/home/clan";
import Roadmap from "~/components/desktop/pages/home/roadmap";
import Team from "~/components/desktop/pages/home/team.vue";
import Partners from "~/components/desktop/pages/home/partners.vue";

export default {
  layout: "home",
  components: {
    Game,
    Nft,
    Clan,
    Roadmap,
    Team,
    Partners,
  },
  created() {},
  methods: {
    coming() {
      this.$toast.add({
        severity: "success",
        summary: "Coming soon",
        life: 3000,
      });
    },
  },
};
</script>
