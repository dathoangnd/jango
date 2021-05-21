<template>
  <div class="md-layout">
    <div class="md-layout-item md-xlarge-size-100 md-large-size-100 md-medium-size-100 md-small-size-100 md-xsmall-size-100" v-for="(grammar, index) in (tab == 'yeu-thich' ? likedGrammars : grammars[tab])" :key="index">
        <Grammar :data="grammar" :level="tab == 'yeu-thich' ? grammar.level : tab" @open="showDetails(tab == 'yeu-thich' ? grammar.level : tab, tab == 'yeu-thich' ? grammar.order : index)" />
    </div>

    <md-dialog :md-active.sync="showDialog">
      <md-table>
        <md-table-toolbar>
          <div>
            <h3 class="md-title md-primary-color">{{ dialogGrammar.text }}</h3>
            <p class="md-subheading">{{ dialogGrammar.meaning }}</p>
          </div>
          <md-button class="md-icon-button" @click="updateLike">
            <md-icon :class="liked ? 'md-accent-color' : ''">favorite</md-icon>
          </md-button>
        </md-table-toolbar>

        <GrammarUsage v-for="(usage, index) in dialogGrammar.usage" :key="index" :usage="usage" />
      </md-table>

      <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Đóng</md-button>
      </md-dialog-actions>
    </md-dialog>
  </div>
</template>

<script>
  import Grammar from '~/components/pages/ngu-phap/Grammar'
  import GrammarUsage from '~/components/pages/ngu-phap/GrammarUsage'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [localStorageManager],
    components: {
      Grammar,
      GrammarUsage
    },
    asyncData({$axios}) {
      return $axios.get(`${process.env.apiUrl}/static/data/grammars/grammar-list.json`)
        .then((res) => {
          return {
            grammars: res.data
          }
        })
    },
    data: function() {
      return {
        tab: (this.$route.query.tab != undefined ? this.$route.query.tab : 'n5'),
        showDialog: false,
        dialogGrammar: {},
        dialogGrammarOrder: ''
      }
    },
    computed: {
      liked() {
        return (this.$store.state.user.likedGrammars.indexOf(this.dialogGrammarOrder) != -1)
      },
      likedGrammars() {
        let grammars = []
        for (let i = 5; i >= 1; i--) {
          for (let j = 0; j < this.grammars['n' + i].length; j++) {
            if (this.$store.state.user.likedGrammars.indexOf(`n${i}-${j+1}`) != -1) {
              let likedGrammar = this.grammars[`n${i}`][j]
              likedGrammar.level = `n${i}`
              likedGrammar.order = j
              grammars.push(likedGrammar)
            }
          }
        }
        return grammars
      }
    },
    methods: {
      showDetails(level, order) {
        this.dialogGrammarOrder = `${level}-${order + 1}`
        this.dialogGrammar = this.grammars[level][order]
        this.showDialog = true
      },
      updateLike: function() {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let formData = new FormData()
          formData.append('grammarOrder', this.dialogGrammarOrder)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: (that.liked ? 'delete' : 'post'),
            url: `${process.env.apiUrl}/grammar/like`, 
            data: formData
          })
          .then(function (response) {
            if (that.liked) {
              that.removeLikedGrammar(that.dialogGrammarOrder)
            } else {
              that.addLikedGrammar(that.dialogGrammarOrder)
            }
            that.$store.commit('user/setLikedGrammars', that.getLikedGrammars())
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
        this.tab = (this.$route.query.tab != undefined ? this.$route.query.tab : 'n5')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-layout {
    .md-layout-item {
      margin-bottom: 30px;
    }
  }

  .md-table {
    .md-table-toolbar {
      justify-content: space-between;
      .md-title {
        font-size: 24px;
        line-height: 36px;
        margin: 10px 0 0 0;
      }
      .md-subheading {
        margin-top: 0;
      }
    }

    /deep/ table {
      table-layout: fixed;
    }
  }
</style>