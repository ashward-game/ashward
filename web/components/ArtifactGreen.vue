<template>
  <div class="flag-atf artifact-green">
    <slot />
  </div>
</template>
<script>
export default {
  name: 'ArtifactGreen'
}
</script>
<style lang="scss" scoped>
.artifact-green{
    position: relative;
    background-image: url("/assets/artifact_green.png");
    background-size: 100% 100%;
}
</style>
