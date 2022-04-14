<template>
  <div class="wrapper staking">
    <div class="staking__page">
      <div class="container">
        <div class="el__header text-center mb-5 pb-3">
          <div
            class="el__title"
            data-bs-toggle="modal"
            data-bs-target="#benefitsModal"
            style="
              text-decoration-line: underline;
              color: #dcb52c;
              cursor: pointer !important;
              pointer-events: unset !important;
            "
          >
            Rewards
          </div>

          <div
            class="el__header__note"
            data-bs-toggle="modal"
            data-bs-target="#noteModal"
          >
            <span class="icon-info-circle" />
          </div>
        </div>

        <div class="row">
          <StakingPool1></StakingPool1>
          <StakingPool2></StakingPool2>
          <StakingPool3></StakingPool3>
          <StakingPool4></StakingPool4>
        </div>
      </div>
    </div>

    <StakingModalStake></StakingModalStake>
    <StakingRewards></StakingRewards>

    <!-- Modal benefitsModal -->

    <ModalNote />
  </div>
  <!-- .wrapper -->
</template>

<script>
import ModalNote from "~/components/staking/ModalNote.vue";

export default {
  components: {
    ModalNote,
  },
  mounted() {},
  methods: {},
};
</script>
