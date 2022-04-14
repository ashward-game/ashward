<template>
  <div class="item-menu-social">
    <div class="image">
      <div class="chain">
        <img class="mx-auto" src="/assets/double-chain.png" alt="">
      </div>
    </div>
    <div class="bg">
      <img class="mx-auto" src="/assets/item-menu-social.png" alt="">
      <div class="icon">
        <slot />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'ItemMenuSocial'
}
</script>
<style lang="scss" scoped>
.item-menu-social{
    width: fit-content;
    margin-left: auto;
    position: relative;
    img{
      width: 100%;
    }
    &+.item-menu-social{
        margin-top: -17px;
    }
    .bg{
        position: relative;
        margin-top: -20px;
        z-index: 8;
    }
    .icon{
        position: absolute;
        top:50%;
        left: 50%;
        transform: translate(-50%,-50%);
        color:#fff;
        padding: 5px;
        img{
          max-height: 35px;
        }
    }
}
</style>
