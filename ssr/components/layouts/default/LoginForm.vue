<template>
  <md-dialog :md-active="showLoginForm" @md-closed="closeDialog">
    <md-dialog-title v-show="!showForgotPassword">Tham gia cộng đồng Jango</md-dialog-title>

    <md-tabs md-dynamic-height v-show="!showForgotPassword">
      <md-tab md-label="Đăng nhập">
        <div class="md-layout md-gutter jg-login-wrapper">
          <div class="md-layout-item md-size-100 jg-account-login">
            <p>Nhập tài khoản Jango của bạn</p>
            <form class="jg-account-login-form">
              <md-field :class="validateLogin && $v.email.$invalid ? 'md-invalid' : ''">
                <label>Email</label>
                <md-input v-model.trim="email" type="email"></md-input>
                <span class="md-error">{{ $v.email.required ? 'Email không hợp lệ' : 'Email là bắt buộc' }}</span>
              </md-field>
              <md-field :class="validateLogin && $v.password.$invalid ? 'md-invalid' : ''">
                <label>Mật khẩu</label>
                <md-input v-model.trim="password" type="password"></md-input>
                <span class="md-error">{{ $v.password.required ? 'Mật khẩu phải chứa ít nhất 6 kí tự' : 'Mật khẩu là bắt buộc' }}</span>
              </md-field>
              <p class="md-caption jg-forgot-password" title="Chưa khả dụng">Quên mật khẩu?</p>
            </form>
          </div>
        </div>

        <md-dialog-actions>
          <md-button @click="closeDialog">Hủy</md-button>
          <md-button class="md-raised md-primary" @click="login">Đăng nhập</md-button>
        </md-dialog-actions>
      </md-tab>

      <md-tab md-label="Đăng ký">
        <div class="md-layout md-gutter jg-register-wrapper">
          <div class="md-layout-item md-size-100">
            <form class="jg-account-register-form">
              <md-field :class="validateRegister && $v.newEmail.$invalid ? 'md-invalid' : ''">
                <label>Email</label>
                <md-input v-model.trim="newEmail"></md-input>
                <span class="md-error">{{ $v.newEmail.required ? ($v.newEmail.email ? 'Email quá dài' : 'Email không hợp lệ') : 'Email là bắt buộc' }}</span>
              </md-field>
              <md-field :class="validateRegister && $v.newPassword.$invalid ? 'md-invalid' : ''">
                <label>Mật khẩu</label>
                <md-input v-model.trim="newPassword" type="password"></md-input>
                <span class="md-error">{{ $v.newPassword.required ? 'Mật khẩu phải chứa ít nhất 6 kí tự' : 'Mật khẩu là bắt buộc' }}</span>
              </md-field>
              <md-field :class="validateRegister && $v.name.$invalid ? 'md-invalid' : ''">
                <label>Tên bạn</label>
                <md-input v-model="name"></md-input>
                <span class="md-error">{{ $v.name.required ? ($v.name.maxLength ? '' : "Tên quá dài") : 'Vui lòng nhập tên của bạn' }}</span>
              </md-field>
            </form>
          </div>
        </div>

        <md-dialog-actions>
          <md-button @click="closeDialog">Hủy</md-button>
          <md-button class="md-raised md-primary" @click="register">Đăng ký</md-button>
        </md-dialog-actions>
      </md-tab>
    </md-tabs>

    <div class="jg-forgot-password-wrapper" v-show="showForgotPassword">
      <div class="md-layout md-gutter">
        <div class="md-layout-item md-size-100">
          <p>Nhập email đăng ký tài khoản của bạn</p>
          <form class="jg-account-register-form">
            <md-field :class="validateForgotPassword && $v.forgotPasswordEmail.$invalid ? 'md-invalid' : ''">
              <label>Email</label>
              <md-input v-model.trim="forgotPasswordEmail"></md-input>
              <span class="md-error">{{ $v.forgotPasswordEmail.required ? ($v.forgotPasswordEmail.email ? 'Email quá dài' : 'Email không hợp lệ') : 'Email là bắt buộc' }}</span>
            </md-field>
          </form>
        </div>
      </div>

      <md-dialog-actions>
        <md-button @click="showForgotPassword = false">Hủy</md-button>
        <md-button class="md-raised md-primary" @click="resetPassword">Đặt lại mật khẩu</md-button>
      </md-dialog-actions>
    </div>
  </md-dialog>
</template>

<script>
  import { required, minLength, maxLength, email } from 'vuelidate/lib/validators'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  import convertToSlug from '~/assets/mixins/convertToSlug.js'

  export default {
    mixins: [localStorageManager, convertToSlug],
    data: () => ({
      validateLogin: false,
      validateRegister: false,
      validateForgotPassword: false,
      showForgotPassword: false,
      email: '',
      password: '',
      newEmail: '',
      newPassword: '',
      name: '',
      forgotPasswordEmail: ''
    }),
    validations: {
      email: {
        required,
        email
      },
      password: {
        required,
        minLength: minLength(6)
      },
      newEmail: {
        required,
        email,
        maxLength: maxLength(255)
      },
      newPassword: {
        required,
        minLength: minLength(6)
      },
      name: {
        required,
        maxLength: maxLength(50)
      },
      forgotPasswordEmail: {
        required,
        email
      }
    },
    computed: {
      showLoginForm() {
        return this.$store.state.showLoginForm && !this.$store.state.loginStatus
      },
      slug() {
        return this.convertToSlug(this.name)
      }
    },
    methods: {
      login() {
        this.validateLogin = true
        if (!(this.$v.email.$invalid || this.$v.password.$invalid)) {
          let that = this
          let formData = new FormData()
          formData.append('email', this.email)
          formData.append('password', this.password)

          this.$axios({
            method: 'post',
            url: `${process.env.apiUrl}/user/login`, 
            data: formData
          })
          .then(function (response) {
            that.$store.commit('setLoginStatus', true)
            that.saveJwt(response.data.token)
            that.saveLikedArticles(response.data.likedArticles)
            that.saveLikedVocas(response.data.likedVocas)
            that.saveLikedGrammars(response.data.likedGrammars)
            that.saveLikedKanjiList(response.data.likedKanji)
            that.saveLikedThreads('')
            that.saveLikedReplies('')

            let payload = that.getPayload()
            payload.likedArticles = that.getLikedArticles()
            payload.likedVocas = that.getLikedVocas()
            payload.likedGrammars = that.getLikedGrammars()
            payload.likedKanji = that.getLikedKanjiList()
            
            that.$store.commit('user/setUser', payload)
            that.$store.commit('setSnackbarMessage', response.data.status)
            that.$store.commit('toggleSnackbar', true)
          })
          .catch(function (error) {
            that.$store.commit('setSnackbarMessage', error.response.data.status)
            that.$store.commit('toggleSnackbar', true)
          });
        }
      },
      register() {
        this.validateRegister = true
        if (!(this.$v.newEmail.$invalid || this.$v.newPassword.$invalid || this.$v.name.$invalid)) {
          let that = this
          let formData = new FormData()
          formData.append('email', this.newEmail)
          formData.append('password', this.newPassword)
          formData.append('name', this.name)
          formData.append('slug', this.slug)

          this.$axios({
            method: 'post',
            url: `${process.env.apiUrl}/user/register`, 
            data: formData
          })
          .then(function (response) {
            that.$store.commit('setLoginStatus', true)
            that.saveJwt(response.data.token)
            that.saveLikedArticles('')
            that.saveLikedVocas('')
            that.saveLikedGrammars('')
            that.saveLikedKanjiList('')
            that.saveLikedThreads('')
            that.saveLikedReplies('')

            let payload = that.getPayload()
            payload.likedArticles = that.getLikedArticles()
            payload.likedVocas = that.getLikedVocas()
            payload.likedGrammars = that.getLikedGrammars()
            payload.likedKanji = that.getLikedKanjiList()
  
            that.$store.commit('user/setUser', payload)
            that.$store.commit('setSnackbarMessage', response.data.status)
            that.$store.commit('toggleSnackbar', true)
          })
          .catch(function (error) {
            that.$store.commit('setSnackbarMessage', error.response.data.status)
            that.$store.commit('toggleSnackbar', true)
          });
        }
      },
      resetPassword() {
        this.validateForgotPassword = true
      },
      closeDialog() {
        this.validateLogin = false
        this.validateRegister = false
        this.validateForgotPassword = false
        this.email = ''
        this.password = ''
        this.newEmail = ''
        this.newPassword = ''
        this.name = ''
        this.forgotPasswordEmail = ''
        this.$store.commit('toggleLoginForm', false)
      }
    },
  }
</script>

<style lang="scss" scoped>
  .md-dialog {
    z-index: 9999;
    /deep/ .md-dialog-container {
      max-width: 100%;
      .jg-login-wrapper {
        .jg-social-login {
          .md-button {
            width: 100%;
            margin: 10px 0;
          }
        }

        .jg-account-login {
          > p {
            margin: 15px 0 0 0;
          }
        }

        .jg-forgot-password {
          margin: 0;
          &:hover {
            cursor: no-drop;
          }
        }
      }

      .jg-forgot-password-wrapper {
        padding: 16px;
      }
    } 
  }
</style>