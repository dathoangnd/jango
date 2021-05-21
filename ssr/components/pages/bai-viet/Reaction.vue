<template>
  <div class="jg-reactions">
    <span class="md-subheading jg-reaction-indicator">{{ likeCount }}</span>
    <md-button class="md-icon-button" @click="$emit('update')">
      <md-icon :class="liked ? 'md-accent-color' : ''">favorite</md-icon>
    </md-button>
    
    <md-button class="md-icon-button">
      <md-icon>share</md-icon>
    </md-button>
  </div>
</template>

<script>
  export default {
    props: {
      liked: Boolean,
      likeCount: Number
    }
  }
</script>


<style lang="scss" scoped>
  .jg-reactions {
    background-color: white;
    .md-button {
      margin: 0;
    }
    .jg-reaction-indicator {
      padding-left: 8px;
      font-size: 16px;
      line-height: 40px;
      color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
    }
  }
</style>

