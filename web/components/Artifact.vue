<template>
  <div class="artifact">
    <slot />
    <!-- <div v-if="dot" class="dot" /> -->
  </div>
</template>
<script>
export default {
  name: "ArtifactField",
  props: {
    dot: {
      type: Boolean,
      required: false,
      default: true,
    },
  },
};
</script>
<style lang="scss" scoped>
.artifact {
  position: relative;
  cursor: pointer;
  background-image: url("/assets/artifact.png");
  background-size: 100% 100%;
  min-width: 150px;
  align-items: center;
  div:first-child {
    text-align: center;
  }
  // .dot {
  //   display: none;
  //   position: absolute;
  //   left: -40%;
  //   top: 50%;
  //   transform: translateY(-50%);
  //   width: 54px;
  //   height: 54px;
  //   background-image: url("/assets/dot.png");
  //   background-repeat: no-repeat;
  //   background-position: center center;
  // }
  &:hover,
  &.active {
    color: #fff;
    background-image: url("/assets/artifact_active.png");
    // .dot {
    //   left: -39%;
    //   width: 60px;
    //   height: 60px;
    //   background-image: url("/assets/dot_active.png");
    // }
  }
  @media (min-width: 1024px) {
    font-size: 32px;
    padding: 8px 25px;
    min-width: 200px;
    // .dot {
    //   display: block;
    //   width: 40px;
    //   height: 40px;
    //   left: -24%;
    // }
    // &.active {
    //   .dot {
    //     left: -26%;
    //     width: 40px;
    //     height: 40px;
    //     background-size: 100% 100%;
    //   }
    // }
  }
  @media (min-width: 1280px) {
    padding: 16px 50px;
    min-width: 180px;
    height: unset;
  }
  @media (min-width: 1536px) {
    min-width: 180px;
    height: 70px;
    padding: 10px 50px;
    // .dot {
    //   width: 40px;
    //   height: 40px;
    //   left: -24%;
    // }
    // &.active {
    //   .dot {
    //     left: -39%;
    //     width: 60px;
    //     height: 60px;
    //   }
    // }
  }
}
</style>
