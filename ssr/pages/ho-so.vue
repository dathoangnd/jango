<template>
  <div>
    <no-ssr>
      <template v-if="tab == ''">
        <form class="md-layout jg-profile-layout">
          <md-card class="md-layout-item md-size-30 md-small-size-100">
            <md-card-content>
              <div @click="changeAvatar" :class="$store.state.user.id == id ? 'jg-avatar-changable' : ''">
                <md-avatar class="md-large jg-user-avatar" v-if="user.avatar != ''">
                  <img :src="user.avatar" :alt="user.name">
                </md-avatar>
                <md-avatar class="md-avatar-icon md-large md-primary jg-user-avatar" v-else>
                  {{ user.name[0] }}
                </md-avatar>
              </div>
              <input type="file" ref="avatarFile" accept="image/*" @change="onFileSelected" v-show="false" />
              <p class="md-title md-primary-color">{{ user.name }}</p>
              <p class="md-subheading">{{ user.about }}</p>
              <a :href="user.facebookId" target="_blank" v-if="user.facebookId != ''">
                <md-button class="md-raised md-primary">Trang Facebook</md-button>
              </a>
            </md-card-content>
          </md-card>
          <md-card class="md-layout-item md-size-65 md-small-size-100">
            <md-card-header>
              <div class="md-title md-primary-color">Thông tin</div>
            </md-card-header>
            
            <md-card-content>
              <div class="md-layout md-gutter">
                <div class="md-layout-item md-small-size-100">
                  <md-field>
                    <label>Tên</label>
                    <md-input v-model="user.name" :disabled="$store.state.user.id != id" />
                  </md-field>
                </div>

                <div class="md-layout-item md-small-size-100">
                  <md-field>
                    <label>Email</label>
                    <md-input :value="user.email" disabled />
                  </md-field>
                </div>
              </div>

              <div class="md-layout md-gutter">
                <div class="md-layout-item md-size-50 md-xsmall-size-100">
                  <md-field>
                    <label>Giới tính</label>
                    <md-select v-model="user.gender" v-if="$store.state.user.id == id">
                      <md-option value="0">Không chọn</md-option>
                      <md-option value="1">Nam</md-option>
                      <md-option value="-1">Nữ</md-option>
                    </md-select>
                    <md-input :value="user.gender == 0 ? 'Không chọn' : (user.gender == 1 ? 'Nam' : 'Nữ')" disabled v-else />
                  </md-field>
                </div>

                <div class="md-layout-item md-size-50 md-xsmall-size-100">
                  <md-datepicker v-model="user.birthday" v-if="$store.state.user.id == id">
                    <label>Sinh nhật</label>
                  </md-datepicker>
                  <md-field v-else>
                    <label>Sinh nhật</label>
                    <md-input :value="user.birthday == null ? 'Không chọn' : calculateTimeDuration(user.birthday)" disabled />
                  </md-field>
                </div>

                <div class="md-layout-item md-size-50 md-xsmall-size-100">
                  <md-field>
                    <label>Đến từ</label>
                    <md-input v-model="user.city" :disabled="$store.state.user.id != id" />
                  </md-field>
                </div>

                <div class="md-layout-item md-size-50 md-xsmall-size-100">
                  <md-field>
                    <label>Địa chỉ Facebook</label>
                    <md-input v-model="user.facebookId" :disabled="$store.state.user.id != id" />
                  </md-field>
                </div>
              </div>

              <md-field v-if="$store.state.user.id == id">
                <label>Giới thiệu</label>
                <md-textarea v-model="user.about" md-autogrow></md-textarea>
              </md-field>
            </md-card-content>
            <md-card-actions v-if="$store.state.user.id == id">
              <md-button @click="showChangePassword = true">Đổi mật khẩu</md-button>
              <md-button class="md-primary" :disabled="profileUpdated" @click="updateProfile">Cập nhật</md-button>
            </md-card-actions>
          </md-card>
        </form>

        <md-dialog :md-active="showChangePassword" @md-closed="closeDialog">
          <md-dialog-title>Đổi mật khẩu</md-dialog-title>
          <div class="jg-change-password-wrapper">
            <div class="md-layout md-gutter">
              <div class="md-layout-item md-size-100">
                <form class="jg-account-register-form">
                  <md-field :class="validatePassword && $v.oldPassword.$invalid ? 'md-invalid' : ''">
                    <label>Mật khẩu cũ</label>
                    <md-input v-model="oldPassword" type="password"></md-input>
                    <span class="md-error">{{ $v.oldPassword.required ? 'Mật khẩu phải chứa ít nhất 6 kí tự' : 'Mật khẩu là bắt buộc' }}</span>
                  </md-field>
                  <md-field :class="validatePassword && $v.newPassword.$invalid ? 'md-invalid' : ''">
                    <label>Mật khẩu mới</label>
                    <md-input v-model="newPassword" type="password"></md-input>
                    <span class="md-error">{{ $v.newPassword.required ? 'Mật khẩu phải chứa ít nhất 6 kí tự' : 'Mật khẩu là bắt buộc' }}</span>
                  </md-field>
                </form>
              </div>
            </div>

            <md-dialog-actions>
              <md-button @click="closeDialog">Hủy</md-button>
              <md-button class="md-raised md-primary" @click="changePassword">Đổi mật khẩu</md-button>
            </md-dialog-actions>
          </div>
        </md-dialog>
      </template>

      <template v-if="tab == 'tai-lieu'">
        <md-table class="md-scrollbar" v-model="documents" md-sort="updatedAt" md-sort-order="desc">
          <md-table-toolbar>
            <h3 class="md-title md-primary-color" v-if="documents.length > 0">Tài liệu của {{ user.name }}</h3>
            <h3 class="md-title md-primary-color" v-else>Không có tài liệu</h3>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
              <md-table-cell md-label="Tiêu đề" md-sort-by="title">
                <template v-if="$store.state.user.id != id && $store.state.user.id != adminId">
                  <nuxt-link :to="`/tai-lieu/${item.slug}-${item.id}`" target="_blank" v-if="item.publish == 1">
                    {{ item.title }}
                  </nuxt-link>
                  <template v-else>
                    {{ item.title }}
                    <md-tooltip md-delay="200">Chưa xuất bản</md-tooltip>
                  </template>
                </template>
                <template v-else>
                  <nuxt-link :class="item.publish == 0 ? 'jg-unpublish-article' : ''" :to="`/chinh-sua/tai-lieu?id=${item.id}`">
                    {{ item.title }}
                    <md-tooltip md-delay="200" v-if="item.publish == 0">Chưa xuất bản</md-tooltip>
                  </nuxt-link> 
                </template>
              </md-table-cell>
            <md-table-cell md-label="Chuyên mục">{{ item.categories | documentCategoryList }}</md-table-cell>
            <md-table-cell md-label="Lượt tải"><md-badge class="md-accent" :md-content="item.download" /></md-table-cell>
            <md-table-cell md-label="Cập nhật" md-sort-by="updatedAt">{{ calculateTimeDuration(item.updatedAt) }}</md-table-cell>
          </md-table-row>
        </md-table>

        <FabButton icon="cloud_upload" to="/chinh-sua/tai-lieu" v-if="$store.state.user.id == id" />
      </template>

      <template v-if="tab == 'bai-viet'">
        <md-table class="md-scrollbar" v-model="articles" md-sort="updatedAt" md-sort-order="desc">
          <md-table-toolbar>
            <h3 class="md-title md-primary-color" v-if="articles.length > 0">Bài viết của {{ user.name }}</h3>
            <h3 class="md-title md-primary-color" v-else>Không có bài viết</h3>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
          
              <md-table-cell md-label="Tiêu đề" md-sort-by="title">
                <template v-if="$store.state.user.id != id && $store.state.user.id != adminId">
                  <nuxt-link :to="`/bai-viet/${item.slug}-${item.id}`" target="_blank" v-if="item.publish == 1">
                    {{ item.title }}
                  </nuxt-link>
                  <template v-else>
                    {{ item.title }}
                    <md-tooltip md-delay="200">Chưa xuất bản</md-tooltip>
                  </template>
                </template>
                <template v-else>
                  <nuxt-link :class="item.publish == 0 ? 'jg-unpublish-article' : ''" :to="`/chinh-sua/bai-viet?id=${item.id}`">
                    {{ item.title }}
                    <md-tooltip md-delay="200" v-if="item.publish == 0">Chưa xuất bản</md-tooltip>
                  </nuxt-link> 
                </template>
              </md-table-cell>
            <md-table-cell md-label="Chuyên mục">{{ item.categories | categoryList }}</md-table-cell>
            <md-table-cell md-label="Thích"><md-badge class="md-accent" :md-content="item.likeCount" /></md-table-cell>
            <md-table-cell md-label="Bình luận"><md-badge class="md-accent" :md-content="item.commentCount" /></md-table-cell>
            <md-table-cell md-label="Cập nhật" md-sort-by="updatedAt">{{ calculateTimeDuration(item.updatedAt) }}</md-table-cell>
          </md-table-row>
        </md-table>
          <FabButton icon="edit" to="/chinh-sua/bai-viet" v-if="$store.state.user.id == id" />
      </template>

      <template v-if="tab == 'thao-luan'">
        <md-table class="md-scrollbar" v-model="threads" md-sort="updatedAt" md-sort-order="desc">
          <md-table-toolbar>
            <h3 class="md-title md-primary-color" v-if="articles.length > 0">Chủ đề của {{ user.name }}</h3>
            <h3 class="md-title md-primary-color" v-else>Không có chủ đề</h3>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
            <md-table-cell md-label="Chủ đề">
              <nuxt-link :to="`/thao-luan/chu-de-${item.id}`" target="_blank">
                {{ extractDescription(item.content, 80) }}
              </nuxt-link>
            </md-table-cell>
            <md-table-cell md-label="Chuyên mục">{{ item.categories | threadCategoryList }}</md-table-cell>
            <md-table-cell md-label="Phiếu bầu"><md-badge class="md-accent" :md-content="item.vote" /></md-table-cell>
            <md-table-cell md-label="Trả lời"><md-badge class="md-accent" :md-content="item.replyCount" /></md-table-cell>
            <md-table-cell md-label="Cập nhật" md-sort-by="updatedAt">{{ calculateTimeDuration(item.updatedAt) }}</md-table-cell>
          </md-table-row>
        </md-table>
      </template>
    </no-ssr>
  </div>
</template>

<script>
  import FabButton from '~/components/pages/shared/FabButton'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  import extractDescription from '~/assets/mixins/extractDescription.js'
  import { required, minLength, maxLength, email } from 'vuelidate/lib/validators'

  export default {
    name: 'Profile',
    mixins: [calculateTimeDuration, localStorageManager, extractDescription],
    components: {
      FabButton
    },
    async asyncData({$axios, query, error}) {
      let userId = (query.id != undefined ? query.id : '')
      if (userId == '') {
        error({ statusCode: 404, message: 'Trang không tồn tại' })
      }
      
      let user = await $axios.get(`${process.env.apiUrl}/user/single?id=${userId}`)
        .then((res) => {
          return {
            name: res.data.user.name,
            email: res.data.user.email,
            facebookId: res.data.user.facebookId,
            gender: res.data.user.gender,
            avatar: res.data.user.avatar,
            birthday: res.data.user.birthday == 0 ? null : res.data.user.birthday,
            city: res.data.user.city,
            about: res.data.user.about
          }
        }).catch((err) => {
          return error({ statusCode: 404, message: 'Trang không tồn tại' })
        })

      let documents = await $axios.get(`${process.env.apiUrl}/document/list?user=${userId}`)
        .then((res) => {
          return res.data.documents
        }).catch((err) => {
          return error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
      
      let articles = await $axios.get(`${process.env.apiUrl}/article/list?user=${userId}`)
        .then((res) => {
          return res.data.articles
        }).catch((err) => {
          return error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
      let threads = await $axios.get(`${process.env.apiUrl}/discussion/list?user=${userId}`)
        .then((res) => {
          return res.data.threads
        }).catch((err) => {
          return error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
      
      return {
        user: user,
        documents: documents,
        articles: articles,
        threads: threads
      }
    },
    created() {
      if (process.client && this.$store.state.user.id == this.adminId && this.$store.state.user.id == this.id) {
        this.$axios.get(`${process.env.apiUrl}/document/list`)
          .then((res) => {
            this.documents = res.data.documents
          })

        this.$axios.get(`${process.env.apiUrl}/article/list`)
          .then((res) => {
            this.articles = res.data.articles
          })
      }
    },
    data: function() {
      return {
        id: this.$route.query.id,
        profileUpdated: true,
        showChangePassword: false,
        validatePassword: false,
        oldPassword: '',
        newPassword: ''
      }
    },
    validations: {
      oldPassword: {
        required,
        minLength: minLength(6)
      },
      newPassword: {
        required,
        minLength: minLength(6)
      }
    },
    computed: {
      adminId() {
        return process.env.adminUserId
      },
      tab(){
        return (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    },
    methods: {
      updateProfile() {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let formData = new FormData()
          formData.append('name', this.user.name.trim())
          formData.append('avatar', this.user.avatar.trim())
          formData.append('birthday', (this.user.birthday == null ? 0 : new Date(this.user.birthday).getTime()))
          formData.append('gender', this.user.gender)
          formData.append('city', this.user.city.trim())
          formData.append('facebookId', this.user.facebookId.trim())
          formData.append('about', this.user.about.trim())
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: 'put',
            url: `${process.env.apiUrl}/user/update`, 
            data: formData
          })
          .then(function (response) {
            that.profileUpdated = true
            that.saveJwt(response.data.token)
            let payload = that.getPayload()
            payload.likedArticles = that.getLikedArticles()
            payload.likedVocas = that.getLikedVocas()
            payload.likedGrammars = that.getLikedGrammars()
            that.$store.commit('user/setUser', payload)

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
      },
      changeAvatar() {
        if (this.$store.state.user.id == this.id) {
          this.$refs.avatarFile.click()
        }
      },
      onFileSelected({ target, dataTransfer }) {
        const files = target.files || dataTransfer.files
        let formData = new FormData()
        formData.append('file', files[0])
        let that = this
        this.$axios({
          method: 'post',
          url: `${process.env.apiUrl}/image/upload`, 
          data: formData
        })
        .then(function(response) {
          that.user.avatar = `${process.env.apiUrl}/${response.data.location}`
        })
      },
      changePassword() {
        this.validatePassword = true
        if (!(this.$v.oldPassword.$invalid || this.$v.newPassword.$invalid)) {
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let that = this
            let formData = new FormData()
            formData.append('oldPassword', this.oldPassword)
            formData.append('password', this.newPassword)
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: 'put',
              url: `${process.env.apiUrl}/user/changepassword`, 
              data: formData
            })
            .then(function (response) {
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
        this.oldPassword = ''
        this.newPassword = ''
        this.validatePassword = false
        this.showChangePassword = false
      }
    },
    watch: {
      'user.name': function() {
        this.profileUpdated = false
      },
      'user.avatar': function() {
        this.profileUpdated = false
      },
      'user.gender': function() {
        this.profileUpdated = false
      },
      'user.birthday': function() {
        this.profileUpdated = false
      },
      'user.city': function() {
        this.profileUpdated = false
      },
      'user.facebookId': function() {
        this.profileUpdated = false
      },
      'user.about': function() {
        this.profileUpdated = false
      }
    },
    filters: {
      documentCategoryList(categories) {
        if (categories.indexOf('giao-trinh') == -1 || categories.indexOf('bai-tap') == -1) {
          return 'Khác'
        }

        return categories.split(' ').map(function(category) {
          switch(category) {
            case 'giao-trinh':
              return 'Giáo trình'
              break
            case 'bai-tap':
              return 'Bài tập'
              break
          }
        }).join(', ')
      },
      threadCategoryList(categories) {
        if (categories.indexOf('kien-thuc') == -1 || categories.indexOf('bai-tap') == -1) {
          return 'Khác'
        }

        return categories.split(' ').map(function(category) {
          switch(category) {
            case 'kien-thuc':
              return 'Kiến thức'
              break
            case 'bai-tap':
              return 'Bài tập'
              break
          }
        }).join(', ')
      },
      categoryList(categories) {
        if (categories.indexOf('kien-thuc') == -1 || categories.indexOf('chia-se') == -1) {
          return 'Khác'
        }

        return categories.split(' ').map(function(category) {
          switch(category) {
            case 'kien-thuc':
              return 'Kiến thức'
              break
            case 'chia-se':
              return 'Chia sẻ'
              break
          }
        }).join(', ')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .jg-profile-layout {
    justify-content: space-between;
    align-items: flex-start;
    > .md-card {
      margin: 0;
      margin-bottom: 30px;

      &:first-child {
        .md-card-content {
          text-align: center;
          .jg-avatar-changable {
            cursor: pointer;
          }
          .jg-user-avatar {
            min-width: 128px;
            min-height: 128px;
            font-size: 64px;
          }
          .md-title {
            margin: 15px 0 0 0;
          }
          .md-subheading {
            margin: 10px 0;
            word-break: break-all;
          }
        }
      }
    }
  }

  .md-table {
    /deep/ table {
      width: 100% !important;
    }

    .jg-unpublish-article {
      color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
      &:hover {
        color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
      }
    }

    .md-badge {
      position: relative;
    }
  }

  .md-dialog {
    .md-title {
      margin: 0;
    }
    .jg-change-password-wrapper {
      padding: 16px;
    }
  }
</style>
