<template>
  <div class="md-layout">
    <div class="md-layout-item md-size-70 md-small-size-100">
      <div class="md-subhead jg-breadcrumb breadcrumb">
        <ol itemscope itemtype="http://schema.org/BreadcrumbList">
          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <nuxt-link to="/" itemprop="item"><span itemprop="name">Jango</span></nuxt-link>
            <meta itemprop="position" content="1" />
          </li>

          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <nuxt-link to="/bai-viet" itemprop="item"><span itemprop="name">Bài viết</span></nuxt-link>
            <meta itemprop="position" content="2" />
          </li>

          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <template v-if="article.categories != ''">
              <nuxt-link v-for="(category, index) in article.categories.trim().split(' ')" :key="index" :to="`/bai-viet?tab=${category}`" itemprop="item"><span itemprop="name">{{ category == 'kien-thuc' ? (article.categories.indexOf('chia-se') != -1 && index == 0 ? 'Kiến thức, ' : 'Kiến thức') : (article.categories.indexOf('kien-thuc') != -1 && index == 0 ? 'Chia sẻ, ' : 'Chia sẻ') }}</span></nuxt-link>
            </template>
            <nuxt-link v-else to="/bai-viet?tab=khac" itemprop="item"><span itemprop="name">Khác</span></nuxt-link>
            <meta itemprop="position" content="3" />
          </li>
        </ol>
      </div>
      <h1 class="md-title md-primary-color" itemprop="name">{{ article.title }}</h1>
      <div class="md-subhead"><nuxt-link :to="`/ho-so?id=${article.userId}`" itemprop="author">{{ article.userName }}</nuxt-link> &bull; <time itemprop="dateModified" :datetime="updatedAtSchemaData">{{ calculateTimeDuration(article.updatedAt) }}</time></div>
      <br />
      <md-divider></md-divider>
      <br />
      <div class="jg-article-content" v-html="article.content" />
      <br />

      <md-divider></md-divider>
        <Reaction :liked="liked" :likeCount="article.likeCount" @update="updateLike" />
      <md-divider></md-divider>
      <br />
         
      <div class="jg-comments" itemprop="comment">
        <Comment v-for="(comment, index) in comments" :key="index" :comment="comment" @deletecomment="deleteComment($event)" />
      </div>
      <md-field :class="validateComment && $v.comment.$invalid ? 'md-invalid' : ''">
        <label>Bình luận của bạn</label>
        <md-textarea v-model="comment"></md-textarea>
        <span class="md-error">{{ $v.comment.required ? 'Bình luận quá dài' : 'Vui lòng nhập bình luận' }}</span>
      </md-field>
      <md-button class="md-raised md-primary" @click="postComment">Bình luận</md-button>
      <br /><br /><br />
    </div>

    <div class="md-layout-item md-size-30 md-small-size-100 jg-newest-articles-wrapper">
      <h4>Bài viết mới nhất</h4>
      <div class="jg-newest-articles">
        <div class="jg-newest-article" v-for="(article, index) in newestArticles" :key="index">
          <div class="jg-newest-article-image">
            <img :src="article.avatar != '' ? article.avatar : defaultAvatar" :alt="article.title" />
          </div>
          <div class="jg-newest-article-summary">
            <nuxt-link :to="`/bai-viet/${article.slug}-${article.id}`"><h3>{{ article.title }}</h3></nuxt-link>
            <p>{{article.userName}}</p>
          </div> 
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Comment from '~/components/pages/bai-viet/Comment'
  import Reaction from '~/components/pages/bai-viet/Reaction'
  import { required, maxLength } from 'vuelidate/lib/validators'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [calculateTimeDuration, localStorageManager],
    components: {
      Comment,
      Reaction
    },
    asyncData({ $axios, params, error }) {
      let id = parseInt(params.slug.slice(params.slug.lastIndexOf('-') + 1))
      if (id == NaN) {
        error({ statusCode: 404, message: 'Trang không tồn tại' })
        return
      }

      return $axios.get(`${process.env.apiUrl}/article/single?id=${id}&newest=5`)
        .then((res) => {
          return res.data
        })
        .catch(e => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    created() {
      let id = parseInt(this.$route.params.slug.slice(this.$route.params.slug.lastIndexOf('-') + 1))
      if (id != undefined && id != '' && id != NaN) {
        let that = this
        this.$axios.get(`${process.env.apiUrl}/article/comment?id=${id}`)
          .then((res) => {
            that.comments = [].concat(res.data.comments)
          })
      }
    },
    data: function() {
      return {
        comments: [
        ],
        comment: "",
        validateComment: false
      }
    },
    computed: {
      updatedAtSchemaData: function() {
        let updatedDate = new Date(this.article.updatedAt);
        return `${updatedDate.getFullYear()}-${updatedDate.getMonth() + 1}-${updatedDate.getDay()}`
      },
      liked() {
        return (this.$store.state.user.likedArticles.indexOf(Number(this.article.id)) != -1)
      },
      defaultAvatar() {
        return process.env.defaultArticleAvatar
      }
    },
    validations: {
      comment: {
        required,
        maxLength: maxLength(500)
      }
    },
    methods: {
      postComment() {
        this.validateComment = true;
        if (!this.$v.comment.$invalid) {
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let formData = new FormData()
            formData.append('articleId', this.article.id)
            formData.append('content', this.comment.trim())
            let that = this
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: 'post',
              url: `${process.env.apiUrl}/article/comment`, 
              data: formData
            })
            .then(function (response) {
              that.comments.unshift({
                userName: response.data.comment.userName,
                userId: response.data.comment.userId,
                userAvatar: response.data.comment.userAvatar,
                updatedAt: response.data.comment.updatedAt,
                content: response.data.comment.content
              })
              that.$store.commit('setSnackbarMessage', response.data.status)
              that.$store.commit('toggleSnackbar', true)
              that.validateComment = false
              that.comment = ''
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
      },
      deleteComment(id) {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let formData = new FormData()
          formData.append('commentId', id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'delete',
            url: `${process.env.apiUrl}/article/comment`, 
            data: formData
          })
          .then(function (response) {
            that.comments = that.comments.filter(function(comment) {
              return comment.id != id
            })
            that.$store.commit('setSnackbarMessage', response.data.status)
            that.$store.commit('toggleSnackbar', true)
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
      },
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
  .md-layout {
    .jg-breadcrumb {
      ol {
        list-style-type: none;
        margin: 0;
        padding: 0;
        li {
          display: inline-block;

          &:not(:last-child)::after {
            content: ">";
          }
          a {
            color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
          }
        }
      }
    }

    h1.md-title {
      font-size: 28px;
      margin: 10px 0;
    }

    .md-subhead {
      a {
        color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
      }
    }

    .jg-reactions {
      padding-left: 16px;
    }

    .jg-newest-articles-wrapper {
      padding-left: 50px;
      h4 {
        margin-top: 0;
        font-size: 18px;
      }

      .jg-newest-article {
        display: flex;
        margin-bottom: 25px;
        .jg-newest-article-image {
          flex-basis: 25%;
          img {
            width: 100%;
          }
        }
        .jg-newest-article-summary {
          flex-basis: 75%;
          padding-left: 10px;
          h3 {
            margin: 0;
            font-size: 14px;
          }
          p {
            margin: 0;
            font-size: 12px;
          }
        }
      }
    }
  }
  
</style>
