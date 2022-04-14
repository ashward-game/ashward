<template>
  <div class="flag-atf artifact-purple">
    <slot />
  </div>
</template>
<script>
export default {
  name: 'ArtifactPurple'
}
</script>
<style lang="scss" scoped>
.artifact-purple{
    position: relative;
    background-image: url("/assets/artifact_purple.png");
    background-size: 100% 100%;
}
</style>
