<template>
  <md-button class="md-fab md-plain" :to="to" @click="$emit('click')">
    <md-icon>{{ icon }}</md-icon>
  </md-button>
</template>

<script>
  export default {
    props: {
      to: String,
      icon: String
    }
  }
</script>


<style lang="scss" scoped>
  .md-fab {
    position: fixed;
    bottom: 50px;
    right: 50px;
    @media(max-width: 599px) {
      & {
        bottom: 30px;
        right: 30px;
      }
    }
  }
</style>

