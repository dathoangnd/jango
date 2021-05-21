<template>
  <div class="md-layout">
    <div class="md-layout-item md-xlarge-size-100 md-large-size-100 md-medium-size-100 md-small-size-100 md-xsmall-size-100" v-for="(level, index) in levels" :key="index">
      <Level :data="level" @open="openModal" />
    </div>
    <md-dialog :md-active.sync="showDialog">
      <md-dialog-title>Đề thi thử {{ dialogLevel.toUpperCase() }} số {{ dialogOrder + 1 }}</md-dialog-title>

      <md-tabs md-dynamic-height  v-if="dialogOrder >= 0">
        <md-tab md-label="Tổng quan">
          <p><b>Thời gian thi:</b> {{ dialogLevel == 'n5' ? 105 : dialogLevel == 'n4' ? 125 : dialogLevel == 'n3' ? 140 : dialogLevel == 'n2' ? 155 : 170 }} phút</p>
          <p><b>Điểm tối đa:</b> 180</p>
          <p><b>Điểm của bạn:</b> {{ levels[5-parseInt(dialogLevel[1])].tests[dialogOrder].score < 0 ? 'Chưa có' : levels[5-parseInt(dialogLevel[1])].tests[dialogOrder].score }}</p>
        </md-tab>

        <md-tab md-label="Xếp hạng">
          <md-table>
            <md-table-row>
              <md-table-head>Hạng</md-table-head>
              <md-table-head>Tên</md-table-head>
              <md-table-head>Điểm</md-table-head>
            </md-table-row>

            <md-table-row v-for="(rank, index) in levels[5-parseInt(dialogLevel[1])].tests[dialogOrder].ranking" :key="index">
              <md-table-cell>{{ index + 1 }}</md-table-cell>
              <md-table-cell><a :href="'/ho-so?id=' + rank.userId" target="_blank">{{ rank.userName }}</a></md-table-cell>
              <md-table-cell>{{ rank.score }}</md-table-cell>
            </md-table-row>
          </md-table>
        </md-tab>
      </md-tabs>

      <md-dialog-actions>
        <md-button @click="showDialog = false">Đóng</md-button>
        <md-button class="md-primary md-raised" @click="showDialog = false">Vào thi</md-button>
      </md-dialog-actions>
    </md-dialog>
  </div>
</template>

<script>
  import Level from '~/components/pages/thi-thu/Level'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [localStorageManager],
    components: {
      Level
    },
    data: function() {
      return {
        levels: [
          {
            title: "n5",
            description: "Trình độ cơ bản",
            tests: [
              {
                score: 120,
                ranking: [
                  {
                    userName: "Dat Hoang",
                    userId: 2,
                    score: 180
                  },
                  {
                    userName: "Tuấn Thành",
                    userId: 3,
                    score: 175
                  },
                  {
                    userName: "Ving Đỗ",
                    userId: 4,
                    score: 160
                  }
                ]
              }
            ]
          }
        ],
        showDialog: false,
        dialogLevel: '',
        dialogOrder: -1
      }
    },
    methods: {
      openModal(level, order) {
        let token = this.getJwt()
        if (token == null) {
          this.$store.commit('setLoginStatus', false)
          this.$store.commit('user/resetUser')
        }

        if (this.$store.state.loginStatus) {
          this.dialogLevel = level
          this.dialogOrder = order
          this.showDialog = true
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
  .md-layout {
    .md-layout-item {
      margin-bottom: 30px;
    }
  }

  .md-dialog {
    /deep/ .md-dialog-container {
      max-width: 100%;
      .md-dialog-actions {
        padding: 16px;
      }
    }
  }
</style>