<template>
  <md-card md-with-hover>
    <md-card-area md-inset>
      <md-card-media md-ratio="16:9">
        <nuxt-link :to="`/bai-viet/${slug}`">
          <img :src="document.avatar != '' ? document.avatar : defaultAvatar" :alt="document.title">
        </nuxt-link>
      </md-card-media>

      <md-card-header>
        <nuxt-link :to="`/tai-lieu/${slug}`">
          <h2 class="md-title md-primary-color">{{ document.title }}</h2>
        </nuxt-link>
        <div class="md-subhead">
          <div class="md-subhead-section">
            <span><nuxt-link class="md-normal-color" :to="`/ho-so?id=${document.userId}`">{{ document.userName }}</nuxt-link> &bull; {{ calculateTimeDuration(document.updatedAt) }}</span>
          </div>
        </div>
      </md-card-header>
    </md-card-area>

    <md-card-expand>
      <md-card-actions md-alignment="space-between">
        <DownloadStats :download="document.download" />
      </md-card-actions>

      
    </md-card-expand>
  </md-card>
</template>

<script>
  import DownloadStats from '~/components/pages/tai-lieu/DownloadStats'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'

  export default {
    mixins: [calculateTimeDuration],
    components: {
      DownloadStats
    },
    props: {
      document: Object
    },
    computed: {
      slug() {
        return `${this.document.slug}-${this.document.id}`
      },
      defaultAvatar() {
        return process.env.defaultDocumentAvatar
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-card {
    cursor: default;
    .md-card-media {
      img {
        transition: transform 0.5s;
        &:hover {
          transform: translateY(-50%) scale(1.25);
        }
      }
    }

    .md-card-header {
      margin-bottom: 0;

      .md-title {
        font-size: 20px;
      }
    }

    .md-card-actions {
      line-height: 40px;
    }
  }
</style>