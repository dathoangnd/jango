<template>
  <div class="jg-discuss-wrapper">
    <div class="jg-discuss-vote">
      <span @click="upvote"><md-icon class="md-size-2x">arrow_drop_up</md-icon></span>
        {{ item.vote }}
      <span @click="downvote"><md-icon class="md-size-2x">arrow_drop_down</md-icon></span>
    </div>
    <div class="jg-discuss-content">
      <div class="jg-avatar-wrapper">
        <md-avatar class="md-avatar-icon md-accent">
          <img :src="item.userAvatar" :alt="item.userName" v-if="item.userAvatar != ''" />
          <span v-else>{{ item.userName[0] }}</span>
        </md-avatar>
        <div>
          <div class="md-subheading"><nuxt-link :to="`/ho-so?id=${item.userId}`">{{ item.userName }}</nuxt-link></div>
          <div class="md-caption">{{ calculateTimeDuration(item.updatedAt) }}</div>
        </div>
      </div>
      <div class="jg-discuss">
        <div class="md-body-1">{{ item.content }}</div>
        <div class="jg-item-actions">
          <no-ssr>
            <span class="md-primary-color" v-if="$store.state.user.id != null && ($store.state.user.id == item.userId || $store.state.user.id == adminId)" @click="remove">Xóa</span>
          </no-ssr>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [calculateTimeDuration, localStorageManager],
    props: {
      item: Object,
      type: String
    },
    computed: {
      adminId() {
        return process.env.adminUserId
      }
    },
    methods: {
      upvote() {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          if (this.type == 'thread') {
            if (this.getLikedThreads().indexOf(this.item.id) != -1) {
              this.$store.commit('setSnackbarMessage', 'Đã bình chọn chủ đề')
              this.$store.commit('toggleSnackbar', true)
              return
            }
          } else {
            if (this.getLikedReplies().indexOf(this.item.id) != -1) {
              this.$store.commit('setSnackbarMessage', 'Đã bình chọn trả lời')
              this.$store.commit('toggleSnackbar', true)
              return
            }
          }
          let formData = new FormData()
          formData.append(this.type == 'thread' ? 'threadId' : 'replyId', this.item.id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'post',
            url: `${process.env.apiUrl}/discussion/${this.type == 'thread' ? 'upvotethread' : 'upvotereply'}`, 
            data: formData
          })
          .then(function (response) {
            if (that.type == 'thread') {
              that.addLikedThread(that.item.id)
            } else {
              that.addLikedReply(that.item.id)
            }
            that.item.vote++
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
      downvote() {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          if (this.type == 'thread') {
            if (this.getLikedThreads().indexOf(this.item.id) != -1) {
              this.$store.commit('setSnackbarMessage', 'Đã bình chọn chủ đề')
              this.$store.commit('toggleSnackbar', true)
              return
            }
          } else {
            if (this.getLikedReplies().indexOf(this.item.id) != -1) {
              this.$store.commit('setSnackbarMessage', 'Đã bình chọn trả lời')
              this.$store.commit('toggleSnackbar', true)
              return
            }
          }
          let formData = new FormData()
          formData.append(this.type == 'thread' ? 'threadId' : 'replyId', this.item.id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'post',
            url: `${process.env.apiUrl}/discussion/${this.type == 'thread' ? 'downvotethread' : 'downvotereply'}`, 
            data: formData
          })
          .then(function (response) {
            if (that.type == 'thread') {
              that.addLikedThread(that.item.id)
            } else {
              that.addLikedReply(that.item.id)
            }
            that.item.vote--
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
      remove() {
        if (this.type == 'thread' && !confirm('Chủ đề đã xóa sẽ không thể khôi phục. Bạn có thực sự muốn xóa?')) {
          return
        }

        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let formData = new FormData()
          formData.append(this.type == 'thread' ? 'threadId' : 'replyId', this.item.id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'delete',
            url: `${process.env.apiUrl}/discussion/${ this.type == 'thread' ? 'single' : 'reply' }`, 
            data: formData
          })
          .then(function (response) {
            that.$store.commit('setSnackbarMessage', response.data.status)
            that.$store.commit('toggleSnackbar', true)
            that.$emit('remove', that.item.id)
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
  .jg-discuss-wrapper {
    display: flex;
    align-items: center;
    margin: 16px 0;
    .jg-discuss-vote {
      min-width: 64px;
      display: flex;
      flex-direction: column;
      align-items: center;
      .md-icon {
        margin: 0;
        width: 24px;
        min-width: 24px;
        height: 24px;
      }
    }
    .jg-discuss-content {
      .jg-avatar-wrapper {
        display: flex;
        justify-content: flex-start;
        margin-bottom: 8px;
        .md-avatar {
          margin: 0 12px 0 0;
        }
      }
      .jg-discuss {
        .jg-item-actions {
          cursor: pointer;
        }
      }
    }
  }
</style>
