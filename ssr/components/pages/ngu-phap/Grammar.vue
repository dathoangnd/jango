<template>
  <md-card :class="accent ? 'md-accent' : 'md-primary'">
    <md-card-header>
      <md-card-header-text>
        <div class="md-title">{{ data.text }}</div>
        <div class="md-subheading">{{ data.meaning }}</div>
      </md-card-header-text>

      <md-avatar class="md-avatar-icon md-small">{{ level.toUpperCase() }}</md-avatar>
    </md-card-header>
    
    <md-list>
      <md-list-item @click="$emit('open')">
        <md-icon class="md-xsmall-hide">assessment</md-icon>
        <span class="md-list-item-text" >Xem chi tiết</span>
      </md-list-item>
    </md-list>
  </md-card>
  
</template>

<script>
  export default {
    props: {
      data: Object,
      accent: Boolean,
      level: String
    },
    data() {
      return {
        expandLessons: false
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-card-header {
    display: flex;
  }

  .md-list {
    padding: 0;
  }
</style>