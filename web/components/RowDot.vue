<template>
  <div class="row-dot">
    <div class="dot"></div>
    <slot />
  </div>
</template>
<style lang="scss" scoped>
.row-dot {
    display: flex;
    align-items: center;
    & + .row-dot{
        margin-top: 8px;
    }
  .dot {
    position: relative;
    margin-right: 8px;
    width: 12px;
    height: 12px;
    background-image: url("/assets/dot.png");
    background-repeat: no-repeat;
    background-position: center center;
    background-size: 100% 100%;
  }
}
</style>