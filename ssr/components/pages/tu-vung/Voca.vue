<template>
  <md-table-row v-show="filter == '' || liked">
    <md-table-cell>
      <md-button @click="updateLike">
        <md-icon :class="liked ? 'md-accent-color' : ''">favorite</md-icon>
      </md-button>
    </md-table-cell>
    <md-table-cell>
      <div class="jg-voca-speaker" @click="speak(voca.text)"><md-icon>volume_up</md-icon></div>
      <div>
        <div class="md-primary-color md-subheading" >{{voca.text}}</div>
        <div class="md-accent-color md-caption">{{voca.read != '' ? voca.read : voca.text}}</div>
        <div>{{voca.meaning}}</div>
      </div>
    </md-table-cell>
    <md-table-cell>
      <div class="jg-voca-speaker" @click="speak(voca.exampleText)"><md-icon >volume_up</md-icon></div>
      <div>
        <div class="md-primary-color md-subheading">{{voca.exampleText}}</div>
        <div class="md-accent-color md-caption">{{voca.exampleRead != '' ? voca.exampleRead : voca.exampleText}}</div>
        <div>{{voca.exampleMeaning}}</div>
      </div>
    </md-table-cell>
  </md-table-row>
</template>

<script>
  import speak from '~/assets/mixins/speak.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [speak, localStorageManager],
    props: {
      voca: Object,
      order: Number
    },
    data: function(){
      return {
        filter: (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    },
    computed: {
      liked() {
        return (this.$store.state.user.likedVocas.indexOf(this.order) != -1)
      }
    },
    methods: {
      updateLike: function() {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          let formData = new FormData()
          formData.append('vocaOrder', this.order)
          let that = this
          this.$axios.setToken(token, 'Bearer')
          this.$axios({
            method: (that.liked ? 'delete' : 'post'),
            url: `${process.env.apiUrl}/vocabulary/like`, 
            data: formData
          })
          .then(function (response) {
            if (that.liked) {
              that.removeLikedVoca(that.order)
            } else {
              that.addLikedVoca(that.order)
            }
            that.$store.commit('user/setLikedVocas', that.getLikedVocas())
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
        this.filter = (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    },
  }
</script>

<style lang="scss" scoped>
  .md-table-cell {
    &:first-child {
      width: 50px;
      /deep/ .md-table-cell-container {
        padding: 0;
        .md-button {
          min-width: auto;
        }
      }
    }
    &:nth-child(2) {
      width: 30%;
    }
  
    /deep/ .md-table-cell-container {
      display: flex;
      padding: 6px 10px;
      @media (max-width: 599px) {
        & {
          display: block;
        }
      }
      .jg-voca-speaker {
        padding-right: 10px;
        &:hover {
          cursor: pointer;
        }
        align-self: center; 
      }
      .md-subheading {
        margin-bottom: 5px;
        font-weight: bold;
      }
    }
  }
</style>