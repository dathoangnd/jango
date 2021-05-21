<template>
  <div class="md-layout">
    <div class="md-layout-item md-size-10 md-medium-size-10 md-small-size-20 md-xsmall-size-25" v-for="(kanji, index) in (tab == 'yeu-thich' ? likedKanji : kanjiList[tab])" :key="index" @click="showDetails(tab == 'yeu-thich' ? kanji.level : tab, tab == 'yeu-thich' ? kanji.order : index)">
      <Kanji :kanji="kanji" />
    </div>
    <md-dialog :md-active.sync="showDialog">
      <div class="jg-kanji-like">
        <md-button class="md-icon-button" @click="updateLike">
          <md-icon :class="liked ? 'md-accent-color' : ''">favorite</md-icon>
        </md-button>
      </div>
      <div class="jg-kanji-summary">
        <div class="md-layout-item md-size-60">
          <md-table>
            <md-table-row>
              <md-table-cell class="md-primary-color">Hán Việt</md-table-cell>
              <md-table-cell>{{ dialogKanji.read }}</md-table-cell>
            </md-table-row>
            <md-table-row>
              <md-table-cell class="md-primary-color">Nghĩa</md-table-cell>
              <md-table-cell>{{ dialogKanji.meaning }}</md-table-cell>
            </md-table-row>
            <md-table-row>
              <md-table-cell class="md-primary-color">Onyomi</md-table-cell>
              <md-table-cell>{{ dialogKanji.onyomi }}</md-table-cell>
            </md-table-row>
            <md-table-row>
              <md-table-cell class="md-primary-color">Kunyomi</md-table-cell>
              <md-table-cell>{{ dialogKanji.kunyomi }}</md-table-cell>
            </md-table-row>
          </md-table>
        </div>
        <div class="md-layout-item md-size-40">
          <h3 class="md-display-4 md-primary-color">{{ dialogKanji.text }}</h3>
        </div>
      </div>
      <md-divider />

      <md-table>
        <md-table-row v-for="(example, index) in dialogKanji.examples" :key="index" @click="speak(example.text)">
          <md-table-cell>
            <div class="md-subheading md-primary-color">{{ example.text }}</div>
            <div>{{ example.read }}</div>
          </md-table-cell>
          <md-table-cell>{{ example.meaning }}</md-table-cell>
        </md-table-row>
      </md-table>

      <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Đóng</md-button>
      </md-dialog-actions>
    </md-dialog>
  </div>
</template>

<script>
  import Kanji from '~/components/pages/chu-han/Kanji.vue'
  import speak from '~/assets/mixins/speak.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  
  export default {
    mixins: [speak, localStorageManager],
    components: {
      Kanji
    },
    asyncData({ $axios, params, error }) {
      return $axios.get(`${process.env.apiUrl}/static/data/kanji/kanji-list.json`)
        .then((res) => {
          return {
            kanjiList: res.data
          }
        })
        .catch(e => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    data() {
      return {
        tab: (this.$route.query.tab != undefined ? this.$route.query.tab : 'n5'),
        showDialog: false,
        dialogKanji: {},
        dialogKanjiOrder: ''
      }
    },
    computed: {
      liked() {
        return (this.$store.state.user.likedKanji.indexOf(this.dialogKanjiOrder) != -1)
      },
      likedKanji() {
        let kanji = []
        for (let i = 5; i >= 1; i--) {
          for (let j = 0; j < this.kanjiList['n' + i].length; j++) {
            if (this.$store.state.user.likedKanji.indexOf(`n${i}-${j+1}`) != -1) {
              let likedKanji = this.kanjiList[`n${i}`][j]
              likedKanji.level = `n${i}`
              likedKanji.order = j
              kanji.push(likedKanji)
            }
          }
        }
        return kanji
      }
    },
    watch: {
      '$route.query.tab': function() {
        this.tab = (this.$route.query.tab != undefined ? this.$route.query.tab : 'n5')
      }
    },
    methods: {
      showDetails(level, order) {
        this.dialogKanjiOrder = `${level}-${order + 1}`
        this.dialogKanji = this.kanjiList[level][order]
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
          formData.append('kanjiOrder', this.dialogKanjiOrder)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: (that.liked ? 'delete' : 'post'),
            url: `${process.env.apiUrl}/kanji/like`, 
            data: formData
          })
          .then(function (response) {
            if (that.liked) {
              that.removeLikedKanji(that.dialogKanjiOrder)
            } else {
              that.addLikedKanji(that.dialogKanjiOrder)
            }
            that.$store.commit('user/setLikedKanji', that.getLikedKanjiList())
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
  .md-dialog {
    max-width: 100%;
    @media (min-width: 480px) {
      min-width: 480px;
    }

    /deep/ .md-dialog-container {
      display: block;
    }
    
    .jg-kanji-like {
      text-align: right;
      padding: 4px 8px;
    }

    .jg-kanji-summary {
      display: flex;
      padding: 0 16px;
      div:last-child {
        text-align: center;
        h3 {
          margin: 8px 0;
        }
      }
      div:first-child {
        .md-table {
          .md-table-cell {
            height: 36px;
            &:first-child {
              width: 64px;
            }
            /deep/ .md-table-cell-container {
              padding: 6px 0;
            }
          }
        }
      }
    }
  }
</style>