<template>
  <button class="arrow-left" @click="$emit('click')"/>
</template>

<script>
export default {
  name: 'ArrowLeft'
}
</script>
<style lang="scss" scoped>
.arrow-left{
    background-image: url("/assets/arrow.png");
    background-size: 100% 100%;
    width: 7rem;
    height: 7rem;
    &:hover{
    background-image: url("/assets/arrow_active.png");
    transform: rotate(180deg);
    }
}
</style>
