<template>
  <div class="md-layout">
    <div class="md-layout-item md-size-100">
      <md-table class="md-scrollbar" v-model="filteredThreads" md-sort="updatedAt" md-sort-order="desc">
        <md-table-toolbar v-if="filteredThreads.length == 0">
          <h3 class="md-title md-primary-color">Không có chủ đề</h3>
        </md-table-toolbar>

        <md-table-row slot="md-table-row" slot-scope="{ item }">
          <md-table-cell>
            <span @click="upvoteThread(item.id)"><md-icon class="md-size-2x">arrow_drop_up</md-icon></span>
            {{ item.vote }}
            <span @click="downvoteThread(item.id)"><md-icon class="md-size-2x">arrow_drop_down</md-icon></span>
          </md-table-cell>
          <md-table-cell md-label="Chủ đề">
            <nuxt-link :to="`/thao-luan/chu-de-${item.id}`">
              <div class="md-body-1">{{ extractDescription(item.content, 80) }}</div>
            </nuxt-link>
            <div class="jg-thread-meta">
              <nuxt-link :to="`/ho-so?id=${item.userId}`">{{ item.userName }}</nuxt-link>
              &bull; {{ item.replyCount }} trả lời
            </div>
          </md-table-cell>
          <md-table-cell md-label="Cập nhật" md-sort-by="updatedAt">{{ calculateTimeDuration(item.updatedAt) }}</md-table-cell>
        </md-table-row>
      </md-table>
    </div>
    <md-dialog :md-active="showNewThread" @md-closed="closeDialog">
      <md-dialog-title class="md-primary-color">Chủ đề mới</md-dialog-title>
      <div class="md-layout">
        <div class="md-layout-item md-size-100">
          <form class="jg-new-thread-form">
            <md-field :class="validateThread && $v.newThreadContent.$invalid ? 'md-invalid' : ''">
              <label>Nội dung</label>
              <md-textarea class="md-scrollbar" v-model="newThreadContent" maxlength="500" md-counter="500"></md-textarea>
              <span class="md-error">{{ $v.newThreadContent.required ? 'Câu hỏi quá dài' : 'Câu hỏi là bắt buộc' }}</span>
            </md-field>
              <div class="md-subheading md-primary-color">Mục</div>
              <md-checkbox v-model="newThreadCategories" value="kien-thuc">Kiến thức</md-checkbox>
              <md-checkbox v-model="newThreadCategories" value="bai-tap">Bài tập</md-checkbox>
         
          </form>
        </div>
      </div>

      <md-dialog-actions>
        <md-button @click="closeDialog">Hủy</md-button>
        <md-button class="md-raised md-primary" @click="publish">Đăng</md-button>
      </md-dialog-actions>
    </md-dialog>
    <FabButton icon="add" @click="showNewThread = true" />
  </div>
</template>

<script>
  import FabButton from '~/components/pages/shared/FabButton'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import extractDescription from '~/assets/mixins/extractDescription.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  import { required, maxLength } from 'vuelidate/lib/validators'

  export default {
    components: {
      FabButton
    },
    mixins: [calculateTimeDuration, extractDescription, localStorageManager],
    asyncData({ $axios }) {
      return $axios.get(`${process.env.apiUrl}/discussion/list`)
        .then((res) => {
          return {
            threads: res.data.threads
          }
        })
    },
    created() {
      let that = this
      this.filteredThreads = this.threads.filter(function(thread) {
        if (that.category == 'khac')
          return thread.categories.indexOf('kien-thuc') == -1 && thread.categories.indexOf('bai-tap') == -1
        return thread.categories.indexOf(that.category) != -1
      })
    },
    data: function() {
      return {
        category: (this.$route.query.tab != undefined ? this.$route.query.tab : ''),
        showNewThread: false,
        validateThread: false,
        newThreadContent: '',
        newThreadCategories: [],
        filteredThreads: []
      }
    },
    validations: {
      newThreadContent: {
        required,
        maxLength: maxLength(500)
      }
    },
    methods: {
      publish() {
        this.validateThread = true
        if (!this.$v.newThreadContent.$invalid) {  
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let formData = new FormData()
            formData.append('content', this.newThreadContent.trim())
            formData.append('categories', this.newThreadCategories.join(' '))
            let that = this
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: 'post',
              url: `${process.env.apiUrl}/discussion/publish`, 
              data: formData
            })
            .then(function (response) {
              that.threads.unshift(response.data.thread)
              that.closeDialog()
              that.$store.commit('setSnackbarMessage', response.data.status)
              that.$store.commit('toggleSnackbar', true)
            })
            .catch(function (error) {
              that.$store.commit('setSnackbarMessage', error.response.data.status)
              that.$store.commit('toggleSnackbar', true)
            });
          } else {
            this.$store.commit('toggleLoginForm', true)
            this.$store.commit('setSnackbarMessage', 'Vui lòng đăng nhập để tiếp tục')
            this.$store.commit('toggleSnackbar', true)
          }
        }
      },
      closeDialog() {
        this.validateThread = false
        this.newThreadContent = ''
        this.newThreadCategories = []
        this.showNewThread = false
      },
      upvoteThread(id) {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          if (this.getLikedThreads().indexOf(id) != -1) {
            this.$store.commit('setSnackbarMessage', 'Đã bình chọn chủ đề')
            this.$store.commit('toggleSnackbar', true)
            return
          }
          let formData = new FormData()
          formData.append('threadId', id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'post',
            url: `${process.env.apiUrl}/discussion/upvotethread`, 
            data: formData
          })
          .then(function (response) {
            that.addLikedThread(id)
            for (let i = 0; i < that.threads.length; i++) {
              if (that.threads[i].id == id) {
                that.threads[i].vote++
                break
              }
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
      },
      downvoteThread(id) {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          if (this.getLikedThreads().indexOf(id) != -1) {
            this.$store.commit('setSnackbarMessage', 'Đã bình chọn chủ đề')
            this.$store.commit('toggleSnackbar', true)
            return
          }
          let formData = new FormData()
          formData.append('threadId', id)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'post',
            url: `${process.env.apiUrl}/discussion/downvotethread`, 
            data: formData
          })
          .then(function (response) {
            that.addLikedThread(id)
            for (let i = 0; i < that.threads.length; i++) {
              if (that.threads[i].id == id) {
                that.threads[i].vote--
                break
              }
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
    },
    watch: {
      '$route.query.tab': function() {
        this.category = (this.$route.query.tab != undefined ? this.$route.query.tab : '')
        let that = this
        this.filteredThreads = this.threads.filter(function(thread) {
          if (that.category == 'khac')
            return thread.categories.indexOf('kien-thuc') == -1 && thread.categories.indexOf('bai-tap') == -1
          return thread.categories.indexOf(that.category) != -1
        })
      },
      threads() {
        let that = this
        this.filteredThreads = this.threads.filter(function(thread) {
          if (that.category == 'khac')
            return thread.categories.indexOf('kien-thuc') == -1 && thread.categories.indexOf('bai-tap') == -1
          return thread.categories.indexOf(that.category) != -1
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-layout {
    .md-layout-item {
      margin-bottom: 30px;
      .jg-thread-meta {
        /deep/ a {
          color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
        }
      }

      .md-table-cell {
        /deep/ .md-table-cell-container {
          padding: 6px;
        }
        &:first-child {
          width: 64px;
          /deep/ .md-table-cell-container {
            display: flex;
            flex-direction: column;
            align-items: center;
            padding: 6px 0;
            .md-icon {
              margin: 0;
              width: 24px;
              min-width: 24px;
              height: 24px;
            }
          }
        }
        &:last-child {
          width: 80px;
        }
      }
    }
  }

  .md-dialog {
    /deep/ .md-dialog-container {
      min-width: 500px;
      padding: 24px 12px 16px 12px;
      .md-dialog-title {
        padding: 10px 0;
      }
      .md-layout-item {
        margin: 0;
      }

    }
  }
</style>