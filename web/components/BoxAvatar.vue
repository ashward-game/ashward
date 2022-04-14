<template>
  <div class="box-avatar">
    <div class="image mx-auto">
      <img :src="avatar" alt="" />
    </div>
    <div
      class="
        text text-center
        mt-4
        mx-auto
        px-14
        max-w-md
        sm:px-4
        md:max-w-xs
        lg:max-w-sm
        xl:max-w-md
      "
    >
      <text-border-black class="name text-yellow">
        {{ name }}
      </text-border-black>
      <text-border-black v-if="role" class="desc text-white mt-2 lg:mt-2">
        {{ role }}
      </text-border-black>
      <text-border-black v-if="desc" class="desc text-white mt-2 lg:mt-2">
        {{ desc }}
      </text-border-black>
    </div>
  </div>
</template>
<script>
import TextBorderBlack from "./text/TextBorderBlack.vue";
export default {
  name: "BoxAvatar",
  components: { TextBorderBlack },
  props: {
    avatar: {
      type: String,
      required: false,
      default: "",
    },
    name: {
      type: String,
      required: true,
      default: "",
    },
    role: {
      type: String,
      required: false,
      default: "",
    },
    desc: {
      type: String,
      required: false,
      default: "",
    },
  },
};
</script>
<style lang="scss">
.box-avatar {
  width: auto;
  margin-bottom: 31px;
  height: 350px;
  .image {
    width: 170px;
    height: 170px;
    overflow: hidden;
    img {
      width: 100%;
    }
  }
  .desc {
    line-height: 21px;
    font-size: 18px;
  }
  .name {
    font-size: 20px;
  }
  @media (max-width: 360px) {
    .image {
      width: 150px;
      height: 150px;
    }
  }
  @media (min-width: 1024px) {
    height: 400px;
    .image {
      width: 205px;
      height: 205px;
    }
    .desc {
      line-height: 27px;
      font-size: 21px;
    }
    .name {
      font-size: 25px;
    }
  }
  @media (min-width: 1280px) {
    min-height: 294px;
  }
}
</style>
