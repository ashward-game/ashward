<template>
  <div class="flag-atf artifact-red">
    <slot />
  </div>
</template>
<script>
export default {
  name: 'ArtifactRed'
}
</script>
<style lang="scss" scoped>
.artifact-red{
    position: relative;
    background-image: url("/assets/artifact_red.png");
    background-size: 100% 100%;
}
</style>
