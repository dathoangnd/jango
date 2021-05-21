<template>
  <md-card class="md-primary jg-no-select" md-with-hover v-show="text != ''">
    <md-ripple>
      <md-card-header>
        <div class="md-display-1">{{ text }}</div>
      </md-card-header>

      <md-card-content class="md-title">
        {{ read }}
      </md-card-content>
    </md-ripple>
  </md-card>
</template>

<script>
  export default {
    props: {
      text: String,
      read: String
    }
  }
</script>

<style lang="scss" scoped>
  .md-card {
    margin: 0 8px 24px 8px;
    @media (max-width: 1280px) {
      & {
        margin: 0 24px 30px 24px;
      }
    }
    @media (max-width: 960px) {
      & {
        margin: 0 4px 30px 4px;
      }
    }
    /deep/ .md-card-header {
      padding: 8px 0;
      div {
        text-align: center;
        color: white;
        @media (max-width: 960px) {
          & {
            font-size: 30px;
          }
        }
        @media (max-width: 599px) {
          & {
            font-size: 28px;
          }
        }
      }
    }

    /deep/ .md-card-content {
      text-align: center;
      padding: 0 0 4px 0;
      color: var(--md-theme-default-text-accent-on-background-variant, rgba(0, 0, 0, 0.54));
      @media (max-width: 599px) {
        & {
          font-size: 18px;
        }
      }
    }
  }
</style>


