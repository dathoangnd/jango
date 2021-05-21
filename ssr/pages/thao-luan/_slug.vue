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
            <nuxt-link to="/thao-luan" itemprop="item"><span itemprop="name">Thảo luận</span></nuxt-link>
            <meta itemprop="position" content="2" />
          </li>

          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <template v-if="thread.categories != ''">
              <nuxt-link v-for="(category, index) in thread.categories.trim().split(' ')" :key="index" :to="`/thao-luan?tab=${category}`" itemprop="item"><span itemprop="name">{{ category == 'kien-thuc' ? (thread.categories.indexOf('bai-tap') != -1 && index == 0 ? 'Kiến thức, ' : 'Kiến thức') : (thread.categories.indexOf('kien-thuc') != -1 && index == 0 ? 'Bài tập, ' : 'Bài tập') }}</span></nuxt-link>
            </template>
            <nuxt-link v-else to="/thao-luan?tab=khac" itemprop="item"><span itemprop="name">Khác</span></nuxt-link>
            <meta itemprop="position" content="3" />
          </li>
        </ol>
      </div>
      <Discuss :item="thread" type="thread" @remove="removeThread" />
      <md-divider></md-divider>
      <Discuss v-for="(reply, index) in sortedReplies" :key="index" :item="reply" type="reply" @remove="removeReply" />
      <br />
      <md-field :class="validateReply && $v.reply.$invalid ? 'md-invalid' : ''">
        <label>Trả lời của bạn</label>
        <md-textarea v-model="reply"></md-textarea>
        <span class="md-error">{{ $v.reply.required ? 'Trả lời quá dài' : 'Vui lòng nhập trả lời' }}</span>
      </md-field>
      <md-button class="md-raised md-primary" @click="postReply">Trả lời</md-button>
      <br /><br /><br />
    </div>

    <div class="md-layout-item md-size-30 md-small-size-100 jg-newest-threads-wrapper">
      <h4>Chủ đề mới nhất</h4>
      <div class="jg-newest-threads">
        <div class="jg-newest-thread" v-for="(thread, index) in newestThreads" :key="index">
          <div class="jg-newest-thread-summary">
            <nuxt-link :to="`/thao-luan/chu-de-${thread.id}`"><h3>{{ extractDescription(thread.content, 80) }}</h3></nuxt-link>
            <p>{{thread.userName}}</p>
          </div> 
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Discuss from '~/components/pages/thao-luan/Discuss'
  import { required, maxLength } from 'vuelidate/lib/validators'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  import extractDescription from '~/assets/mixins/extractDescription.js'

  export default {
    mixins: [calculateTimeDuration, localStorageManager, extractDescription],
    components: {
      Discuss
    },
    asyncData({ $axios, params, error }) {
      let id = parseInt(params.slug.slice(params.slug.lastIndexOf('-') + 1))
      if (id == NaN) {
        error({ statusCode: 404, message: 'Trang không tồn tại' })
        return
      }

      return $axios.get(`${process.env.apiUrl}/discussion/single?id=${id}&newest=5`)
        .then((res) => {
          return res.data
        })
        .catch(e => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    created() {
      let id = parseInt(this.$route.params.slug.slice(this.$route.params.slug.lastIndexOf('-') + 1))
      if (id != NaN) {
        let that = this
        this.$axios.get(`${process.env.apiUrl}/discussion/reply?id=${id}`)
          .then((res) => {
            that.replies = [].concat(res.data.replies)
          })
      }
    },
    data: function() {
      return {
        replies: [],
        reply: "",
        validateReply: false
      }
    },
    computed: {
      updatedAtSchemaData: function() {
        let updatedDate = new Date(this.thread.updatedAt);
        return `${updatedDate.getFullYear()}-${updatedDate.getMonth() + 1}-${updatedDate.getDay()}`
      },
      sortedReplies() {
        return this.replies.sort(function(replyA, replyB) {
          return replyB.vote - replyA.vote
        })
      }
    },
    validations: {
      reply: {
        required,
        maxLength: maxLength(500)
      }
    },
    methods: {
      postReply() {
        this.validateReply = true;
        if (!this.$v.reply.$invalid) {
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let formData = new FormData()
            formData.append('threadId', this.thread.id)
            formData.append('content', this.reply.trim())
            let that = this
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: 'post',
              url: `${process.env.apiUrl}/discussion/reply`, 
              data: formData
            })
            .then(function (response) {
              that.replies.unshift(response.data.reply)
              that.$store.commit('setSnackbarMessage', response.data.status)
              that.$store.commit('toggleSnackbar', true)
              that.validateReply = false
              that.reply = ''
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
      removeThread(id) {
        this.$router.push(`/thao-luan`)
      },
      removeReply(id) {
        this.replies = this.replies.filter(function(reply) {
          return reply.id != id
        })
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

    .jg-newest-threads-wrapper {
      padding-left: 50px;
      h4 {
        margin-top: 0;
        font-size: 18px;
      }

      .jg-newest-thread {
        margin-bottom: 25px;
        .jg-newest-thread-summary {
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