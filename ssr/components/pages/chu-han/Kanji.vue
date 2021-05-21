<template>
  <md-card class="md-primary jg-no-select" md-with-hover>
    <md-ripple>
      <md-card-header>
        <div class="md-display-1">{{ kanji.text }}</div>
      </md-card-header>

      <md-card-content class="md-subheading">
        {{ kanji.meaning }}
      </md-card-content>
    </md-ripple>
  </md-card>
</template>

<script>
  export default {
    props: {
      kanji: Object
    }
  }
</script>

<style lang="scss" scoped>
  .md-card {
    margin: 8px;
    /deep/ .md-card-header {
      text-align: center;
      padding-left: 0;
      padding-right: 0;
      padding-bottom: 4px;
      div {
        text-align: center;
        color: white;
      }
    }

    /deep/ .md-card-content {
      font-size: 16px;
      text-align: center;
      color: var(--md-theme-default-text-accent-on-background-variant, rgba(0, 0, 0, 0.54));
      padding: 0 0 8px 0;
    }
  }
</style>