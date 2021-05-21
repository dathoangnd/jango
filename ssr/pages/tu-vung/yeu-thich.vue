<template>
  <div>
    <div class="md-subhead jg-breadcrumb breadcrumb">
      <ol itemscope itemtype="http://schema.org/BreadcrumbList">
        <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
          <nuxt-link to="/" itemprop="item"><span itemprop="name">Jango</span></nuxt-link>
          <meta itemprop="position" content="1" />
        </li>

        <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
          <nuxt-link to="/tu-vung" itemprop="item"><span itemprop="name">Từ vựng</span></nuxt-link>
          <meta itemprop="position" content="2" />
        </li>
      </ol>
    </div>

    <md-table>
      <md-table-toolbar>
        <h1 class="md-title md-primary-color">{{course == 0 ? 'Tất cả từ vựng yêu thích' : course + '000 từ vựng thiết yếu - Yêu thích'}}</h1>
      </md-table-toolbar>

      <Voca v-for="(voca, index) in vocabularies" :key="index" :voca="voca" :order="voca.id" />
    </md-table>
  </div>
</template>

<script>
  import Voca from '~/components/pages/tu-vung/Voca.vue'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [localStorageManager],
    components: {
      Voca
    },
    created() {
      if (process.client) {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios.get(`${process.env.apiUrl}/vocabulary/like?course=${this.course}`)
            .then((res) => {
              that.vocabularies = res.data.vocas
            })
            .catch(error => {
              that.$store.commit('setSnackbarMessage', error.response.data.status)
              that.$store.commit('toggleSnackbar', true)
            })
        }
      }
    },
    data() {
      return {
        vocabularies: [],
        course: (this.$route.query.course == undefined || this.$route.query.course == '' || parseInt(this.$route.query.course[0]) == NaN) ? 0 : parseInt(this.$route.query.course[0])
      }
    },
    watch: {
      '$route.query.course': function() {
        this.course = (this.$route.query.course == undefined || this.$route.query.course == '' || parseInt(this.$route.query.course[0]) == NaN) ? 0 : parseInt(this.$route.query.course[0])
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios.get(`${process.env.apiUrl}/vocabulary/like?course=${this.course}`)
            .then((res) => {
              that.vocabularies = res.data.vocas
            })
            .catch(error => {
              that.$store.commit('setSnackbarMessage', error.response.data.status)
              that.$store.commit('toggleSnackbar', true)
            })
        }
      }
    }
  }
</script>


<style lang="scss" scoped>
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

  .md-table {
    .md-table-toolbar {
      padding: 0;
      .md-title {
        font-size: 28px;
        line-height: 36px;
        margin: 10px 0 30px 0;
      }
    }

    /deep/ table {
      table-layout: fixed;
    }
  }
</style>
