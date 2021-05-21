<template>
  <div class="md-layout md-gutter">
    <div class="md-layout-item md-xlarge-size-70 md-large-size-70 md-medium-size-60 md-small-size-100 md-xsmall-size-100">
      <md-field :class="validateForm && $v.title.$invalid ? 'jg-new-document-title-wrapper md-invalid' : 'jg-new-article-title-wrapper'">
        <label>Tiêu đề</label>
        <md-input v-model="title"></md-input>
        <span class="md-error">{{ $v.title.required ? 'Tiêu đề không được quá 255 kí tự' : 'Tiêu đề là bắt buộc' }}</span>
      </md-field>
      <md-field :class="validateForm && $v.file.$invalid ? 'jg-new-document-title-wrapper md-invalid' : 'jg-new-document-title-wrapper'">
        <label>Tài liệu</label>
        <MdFile v-model.trim="file" type="pdf" accept="application/pdf" />
        <span class="md-error">{{ $v.title.required ? 'Đường dẫn tài liệu không được quá 255 kí tự' : 'Tài liệu là bắt buộc' }}</span>
      </md-field>
      <md-field :class="validateForm && $v.content.$invalid ? 'jg-new-document-content-wrapper md-invalid' : 'jg-new-document-content-wrapper'">
          <editor v-model="content" :api-key="tinyMceKey" :init="{
            statusbar: false, menubar: false, contextmenu: false,
            width: '100%', height : 'calc(100vh - 300px)',
            image_dimensions: false,
            image_class_list: [{ title:'None', value:'jg-document-image' },{ title:'Box', value:'jg-document-image jg-document-image-box' }, { title:'Full Width', value:'jg-document-image jg-document-image-full-width' }],
            toolbar: 'formatselect | bold italic forecolor image link table numlist bullist alignleft aligncenter alignjustify code preview',
            images_upload_url: imageUploadUrl, images_upload_base_path: apiUrl,
            language_url: '/tinymce/langs/vi_VN.js', language: 'vi_VN',
            content_css: '/tinymce/css/styles.css',
            plugins: 'image link table lists imagetools preview code'}" />
        <span class="md-error">{{ $v.content.required ? 'Nội dung quá dài' : 'Nội dung là bắt buộc' }}</span>
      </md-field>
    </div>

    <div class="md-layout-item md-xlarge-size-30 md-large-size-30 md-medium-size-40 md-small-size-100 md-xsmall-size-100">
      <md-card class="jg-new-document-meta-wrapper">
        <md-card-header>
          <span class="md-subheading">Xuất bản</span>
          <md-menu md-direction="bottom-end" :md-offset-x="20" md-align-trigger v-show="$store.state.user.id == adminId || publishStatus">
            <md-button md-menu-trigger>
              <md-icon>more_vert</md-icon>
            </md-button>

            <md-menu-content>
              <nuxt-link class="md-list-item-container" target="_blank" :to="`/tai-lieu/${slug}-${pageId}`" v-if="publishStatus">
                <md-menu-item>Tới trang</md-menu-item>
              </nuxt-link>
              <md-menu-item @click="changePublicStatus(!publishStatus)" v-if="$store.state.user.id == adminId && publishedContent != ''">{{ publishStatus ? 'Thu hồi' : 'Xuất bản' }}</md-menu-item>
              <md-menu-item @click="remove" v-if="$store.state.user.id == userId">Xóa</md-menu-item>
            </md-menu-content>
          </md-menu>
        </md-card-header>
        <md-card-content>
          <md-list>
            <md-list-item md-expand>
              <md-icon>assignment_ind</md-icon>
              <span class="md-list-item-text">Tác giả</span>
              <div class="jg-new-document-meta-slot" slot="md-expand">
                <md-field>
                  <label>Tên tác giả</label>
                  <md-input v-model="author"></md-input>
                </md-field>
              </div>
            </md-list-item>

            <md-list-item md-expand>
              <md-icon>rss_feed</md-icon>
              <span class="md-list-item-text">Nguồn</span>
              <div class="jg-new-document-meta-slot" slot="md-expand">
                <md-field>
                  <label>Nguồn tài liệu</label>
                  <md-input v-model="source"></md-input>
                </md-field>
              </div>
            </md-list-item>

            <md-list-item md-expand>
              <md-icon>category</md-icon>
              <span class="md-list-item-text">Loại</span>
              <div class="jg-new-document-meta-slot" slot="md-expand">
                <div v-for="(category, index) in categoryList" :key="index">
                  <md-checkbox v-model="categories" :value="category.slug">{{category.text}}</md-checkbox>
                </div>
              </div>
            </md-list-item>

            <md-list-item md-expand>
              <md-icon>art_track</md-icon>
              <span class="md-list-item-text">Ảnh đại diện</span>
              <div class="jg-new-document-meta-slot" slot="md-expand">
                <md-field>
                  <label>Đường dẫn ảnh</label>
                  <MdFile v-model.trim="avatar" accept="image/*" />
                </md-field>
              </div>
            </md-list-item>
          </md-list>
        </md-card-content>

        <md-card-actions md-alignment="space-between">
          <div class="md-caption">{{ postStatus ? (pageId==''?'':(publishStatus?'Đã xuất bản':'Chờ xét duyệt')): 'Chưa lưu' }}</div>
          <div>
            <md-button class="md-raised md-primary" :disabled="postStatus || ($store.state.user.id != userId && userId != null)" @click="publish()">Đăng</md-button>
          </div>
        </md-card-actions>
      </md-card>
    </div>
  </div>
</template>

<script>
  import MdFile from '~/components/pages/shared/MdFile'
  import { required, maxLength } from 'vuelidate/lib/validators'
  import tinymce from '@tinymce/tinymce-vue'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  import convertToSlug from '~/assets/mixins/convertToSlug.js'
  export default {
    components: {
      MdFile
    },
    mixins: [localStorageManager, convertToSlug],
    beforeRouteLeave (to, from, next) {
      if (this.postStatus || confirm('Các thay đổi bạn đã thực hiện sẽ không được lưu.')) {
        next(true);
      } else {
        next(false);
      }
    },
    asyncData({$axios, query, error}) {
      if (query.id == undefined || query.id == '') {
        return {
          publishStatus: false,
          title: '',
          slug: '',
          content: '',
          publishedContent: '',
          avatar: '',
          categories: [],
          userId: null,
          file: '',
          author: '',
          source: ''
        }
      }

      return $axios.get(`${process.env.apiUrl}/document/single?id=${query.id}`)
        .then((res) => {
          return {
            publishStatus: (res.data.document.publish == '1'),
            title: res.data.document.title,
            slug: res.data.document.slug,
            content: res.data.document.content,
            publishedContent: res.data.document.content,
            avatar: res.data.document.avatar,
            categories: res.data.document.categories.split(' '),
            userId: res.data.document.userId,
            file: res.data.document.file,
            author: res.data.document.author,
            source: res.data.document.source,
          }
        })
        .catch((e) => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    data: function() {
      return {
        validateForm: false,
        
        postStatus: true,
        categoryList: [
          {
            text: "Giáo trình",
            slug: "giao-trinh"
          },
          {
            text: "Bài tập",
            slug: "bai-tap"
          }
        ]
      }
    },
    computed: {
      tinyMceKey: function() {
        return process.env.tinyMceKey
      },
      apiUrl: function() {
        return process.env.apiUrl
      },
      adminId() {
        return process.env.adminUserId
      },
      imageUploadUrl: function() {
        return process.env.apiUrl + '/image/upload'
      },
      pageId: function() {
        return (this.$route.query.id == undefined ? '' : this.$route.query.id)
      }
    },
    validations: {
      title: {
        required,
        maxLength: maxLength(255)
      },
      content: {
        required,
        maxLength: maxLength(5000)
      },
      file: {
        required,
        maxLength: maxLength(255)
      }
    },
    methods: {
      publish() {
        this.validateForm = true;
        if (!(this.$v.title.$invalid || this.$v.content.$invalid || this.$v.file.$invalid)) {
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let formData = new FormData()
            formData.append('title', this.title)
            formData.append('file', this.file)
            formData.append('content', this.content)
            formData.append('author', this.author)
            formData.append('source', this.source)
            formData.append('categories', this.categories.join(' '))
            formData.append('avatar', this.avatar)
            formData.append('id', this.pageId)
            formData.append('slug', this.convertToSlug(this.title))
            let requestMethod = (this.pageId == '' ? 'post' : 'put')
            let that = this
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: requestMethod,
              url: `${process.env.apiUrl}/document/publish`, 
              data: formData
            })
            .then(function (response) {
              that.setChangeStatus(true)
              if (requestMethod == 'post')
                that.$router.push({path: that.$route.path, query: { id: response.data.id }})
              that.$store.commit('setSnackbarMessage', response.data.status)
              that.$store.commit('toggleSnackbar', true)
            })
            .catch(function (error) {
              if (error.response.status == 401) {
                that.$store.commit('setLoginStatus', false)
                that.removeJwt()
                that.$store.commit('user/resetUser')
                that.$store.commit('toggleLoginForm', true)
                that.$store.commit('setSnackbarMessage', 'Phiên đăng nhập đã hết hạn')
                that.$store.commit('toggleSnackbar', true)
              } else {
                that.$store.commit('setSnackbarMessage', error.response.data.status)
                that.$store.commit('toggleSnackbar', true)
              }
            });
          } else {
            this.$store.commit('toggleLoginForm', true)
            this.$store.commit('setSnackbarMessage', 'Vui lòng đăng nhập để tiếp tục')
            this.$store.commit('toggleSnackbar', true)
          }
        }
      },
      setChangeStatus(postStatus) {
        this.postStatus = postStatus
      },
      changePublicStatus(status) {
        if (confirm(`${status ? 'Xuất bản' : 'Thu hồi'} tài liệu?`)) {
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let formData = new FormData()
            formData.append('id', this.pageId)
            formData.append('status', status ? '1' : '0')
            let that = this
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: 'put',
              url: `${process.env.apiUrl}/document/status`, 
              data: formData
            })
            .then(function (response) {
              that.$store.commit('setSnackbarMessage', response.data.status)
              that.$store.commit('toggleSnackbar', true)
              that.publishStatus = status
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
      remove() {
        if (confirm('Bài viết đã xóa sẽ không thể khôi phục. Bạn có thực sự muốn xóa?')) {
          let token = this.getJwt()
          if (token == null) {
            this.$store.commit('setLoginStatus', false)
            this.$store.commit('user/resetUser')
          }

          if (this.$store.state.loginStatus) {
            let formData = new FormData()
            formData.append('id', this.pageId)
            let that = this
            this.$axios.setToken(token, 'Bearer')
            this.$axios({
              method: 'delete',
              url: `${process.env.apiUrl}/document/single`, 
              data: formData
            })
            .then(function (response) {
              that.setChangeStatus(true)
              that.$store.commit('setSnackbarMessage', response.data.status)
              that.$store.commit('toggleSnackbar', true)
              that.$router.push(`/tai-lieu`)
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
      }
    },
    watch: {
      title() {
        this.setChangeStatus(false)
      },
      content() {
        this.setChangeStatus(false)
      },
      file() {
        this.setChangeStatus(false)
      },
      avatar() {
        this.setChangeStatus(false)
      },
      author() {
        this.setChangeStatus(false)
      },
      source() {
        this.setChangeStatus(false)
      },
      categories() {
        this.setChangeStatus(false)
      }
    }
  }
</script>

<style lang="scss" scoped>
  .jg-new-document-title-wrapper {
    margin-bottom: 16px;
  }
  .jg-new-document-content-wrapper {
    margin-bottom: 36px;
  }

  .jg-new-document-meta-wrapper {
    margin: 0 0 36px 0;
    .md-card-header {
      display: flex;
      padding-right: 28px;
      /deep/ a {
        margin-right: 10px;
      }
      .md-subheading {
        font-weight: bold;
        line-height: 36px;
        flex-grow: 1;
      }
      .md-button {
        min-width: 36px;
      }
    }
    .md-card-content {
      .jg-new-document-meta-slot {
        padding: 0 16px;
        .md-checkbox {
          margin: 8px 0;
        }
      }
    }
  }
</style>