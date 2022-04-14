<template>
  <div class="all-divider" />
</template>
<script>
export default {
  name: 'AllDivider'
}
</script>
<style lang="scss" scoped>
.all-divider{
    background-image: url("/assets/All_Divider.png");
    width: 100%;
    height: 80px;
    background-size: auto 100%;
    background-repeat: repeat-x;
    @media (min-width:1024px) {
    height: 150px;
    }
}
</style>
