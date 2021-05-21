<template>
  <md-card md-with-hover>
    <md-card-area md-inset>
      <md-card-media md-ratio="16:9">
        <nuxt-link :to="`/bai-viet/${slug}`">
          <img :src="article.avatar != '' ? article.avatar : defaultAvatar" :alt="article.title">
        </nuxt-link>
      </md-card-media>

      <md-card-header>
        <nuxt-link :to="`/bai-viet/${slug}`">
          <h2 class="md-title md-primary-color">{{ article.title }}</h2>
        </nuxt-link>
        <div class="md-subhead">
          <div class="md-subhead-section">
            <span><nuxt-link class="md-normal-color" :to="`/ho-so?id=${article.userId}`">{{ article.userName }}</nuxt-link> &bull; {{ calculateTimeDuration(article.updatedAt) }}</span>
          </div>
        </div>
      </md-card-header>
    </md-card-area>

    <md-card-expand>
      <md-card-actions md-alignment="space-between">
        <Reaction :liked="liked" :likeCount="article.likeCount" @update="updateLike" />

        <md-card-expand-trigger>
          <md-button class="md-icon-button">
            <md-icon>keyboard_arrow_down</md-icon>
          </md-button>
        </md-card-expand-trigger>
      </md-card-actions>

      <md-card-expand-content>
        <div>
          <md-card-content>
            <no-ssr>
              {{ article.description != '' ? article.description : extractDescription(article.content, 80) }}
            </no-ssr>
          </md-card-content>
          <md-card-actions>
            <div>
              <nuxt-link :to="`/bai-viet/${slug}`">
                <md-button>Xem thêm</md-button>
              </nuxt-link>
            </div>
          </md-card-actions>
        </div>
      </md-card-expand-content>
    </md-card-expand>
  </md-card>
</template>

<script>
  import Reaction from '~/components/pages/bai-viet/Reaction'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import extractDescription from '~/assets/mixins/extractDescription.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [calculateTimeDuration, extractDescription, localStorageManager],
    components: {
      Reaction
    },
    props: {
      article: Object
    },
    computed: {
      slug() {
        return `${this.article.slug}-${this.article.id}`
      },
      defaultAvatar() {
        return process.env.defaultArticleAvatar
      },
      liked() {
        return (this.$store.state.user.likedArticles.indexOf(Number(this.article.id)) != -1)
      }
    },
    methods: {
      updateLike: function() {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let formData = new FormData()
          formData.append('articleId', this.article.id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: (that.liked ? 'delete' : 'post'),
            url: `${process.env.apiUrl}/article/like`, 
            data: formData
          })
          .then(function (response) {
            if (that.liked) {
              that.removeLikedArticle(that.article.id)
              that.$store.commit('user/setLikedArticles', that.getLikedArticles())
              that.article.likeCount--
            } else {
              that.addLikedArticle(that.article.id)
              that.$store.commit('user/setLikedArticles', that.getLikedArticles())
              that.article.likeCount++
            }
          })
          .catch(function (error) {
            that.$store.commit('setSnackbarMessage', 'Đã xảy ra sự cố, vui lòng thử lại.')
            that.$store.commit('toggleSnackbar', true)
          });
        } else {
          this.$store.commit('toggleLoginForm', true)
          this.$store.commit('setSnackbarMessage', 'Vui lòng đăng nhập để tiếp tục')
          this.$store.commit('toggleSnackbar', true)
        }
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