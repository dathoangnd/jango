<template>
  <div class="jg-download-stats">
    <span class="md-subheading jg-download-indicator">{{ download }}</span>
    <a :href="file" target="_blank" download v-if="file != undefined">
      <md-button class="md-icon-button">
        <md-icon>cloud_download</md-icon>
      </md-button>
    </a>
    <md-button class="md-icon-button jg-download-button" v-else>
      <md-icon>cloud_download</md-icon>
    </md-button>
    
    <md-button class="md-icon-button">
      <md-icon>share</md-icon>
    </md-button>
  </div>
</template>

<script>
  export default {
    props: {
      download: Number,
      file: String
    }
  }
</script>


<style lang="scss" scoped>
  .jg-download-stats {
    background-color: white;
    .md-button {
      margin: 0;
    }
    .jg-download-indicator {
      padding-left: 8px;
      font-size: 16px;
      line-height: 40px;
      color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
    }
    .jg-download-button {
      &:hover {
        cursor: default;
      }
      &::before {
        display: none;
      }
      /deep/ .md-ripple-wave {
        display: none;
      }
    }
  }
</style>

