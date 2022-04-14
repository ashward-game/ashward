<template>
  <button @click="$emit('click')" class="btn-red-hover">
    <slot />
  </button>
</template>
<script>
export default {
  name: 'BtnRedHover'
}
</script>
<style lang="scss" scoped>
.btn-red-hover{
  cursor:pointer;
  background-image: url("/assets/button/S4_Button_Default.png");
  background-size: 100% 100%;
  display: inline-block;
  color: #fff;
  transition: all .4s;
  &:hover,&.active{
  background-image: url("/assets/button/S4_Button_Active.png");
  color: #fff225;
  transition: all .4s;
  }
  }
</style>
