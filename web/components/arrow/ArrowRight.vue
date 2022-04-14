<template>
  <button class="arrow-right"  @click="$emit('click')"/>
</template>

<script>
export default {
  name: 'ArrowRight'
}
</script>
<style lang="scss" scoped>
.arrow-right{
    background-image: url("/assets/arrow.png");
    background-size: 100% 100%;
    width: 7rem;
    height: 7rem;
    transform: rotate(180deg);
    &:hover{
    background-image: url("/assets/arrow_active.png");
    transform: rotate(0);
    }
}
</style>
