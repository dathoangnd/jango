<template>
  <div class="page-container md-layout-column">
    <LoginForm />

    <md-app md-mode="fixed">
      <md-app-drawer md-persistent="mini" :md-active.sync="showNavigation" md-swipeable class="md-scrollbar">
        <md-toolbar class="md-transparent" md-elevation="0">
          <nuxt-link to="/">
            <span class="md-title md-primary-color">Jango</span>
          </nuxt-link>
        </md-toolbar>

        <md-divider></md-divider>

        <md-list> 
          <md-list-item v-for="(item, index) in upperMenuItems" :key="index" :to="item.route">
            <md-icon :class="menuItemClass(item.route)">{{ item.icon }}</md-icon>
            <span class="md-list-item-text" :class="menuItemClass(item.route)">{{ item.text }}</span>
            <md-tooltip md-direction="right" md-delay="200" v-show="!showNavigation">{{ item.text }}</md-tooltip>
          </md-list-item>

          <md-divider></md-divider>

          <md-list-item v-for="(item, index) in lowerMenuItems" :key="index + upperMenuItems.length" :to="item.route">
            <md-icon :class="menuItemClass(item.route)">{{ item.icon }}</md-icon>
            <span class="md-list-item-text" :class="menuItemClass(item.route)">{{ item.text }}</span>
            <md-tooltip md-direction="right" md-delay="200" v-show="!showNavigation">{{ item.text }}</md-tooltip>
          </md-list-item>

          <md-divider></md-divider>

          <a class="jg-social-menu-item" v-for="(item, index) in socialMenuItems" :key="index + upperMenuItems.length + lowerMenuItems.length" :href="item.route" target="_blank">
            <md-list-item>
                <md-avatar class="md-avatar-icon md-small md-primary">{{ item.icon }}</md-avatar>
                <span class="md-list-item-text">{{ item.text }}</span>
                <md-tooltip md-direction="right" md-delay="200" v-show="!showNavigation">{{ item.text }}</md-tooltip>
            </md-list-item>
          </a>
        </md-list>
      </md-app-drawer>

      <md-app-toolbar class="md-primary">
        <div class="md-toolbar-row jg-top-bar">
          <div class="md-toolbar-section-start">
            <md-button class="md-icon-button" @click="showNavigation = !showNavigation">
              <md-icon>{{ showNavigation ? 'arrow_back_ios' : 'menu' }}</md-icon>
            </md-button>
            <span class="md-title md-xsmall-hide">{{ currentPage }}</span>
          </div>

          <md-autocomplete
            class="search"
            v-model="searchTerm"
            :md-options="searchResults"
            @md-opened="getResults"
            @md-changed="getResults"
            @md-selected="val=>chooseResult(val.path, val.group)"
            md-layout="box">
            <label><md-icon>search</md-icon> Tìm kiếm...</label>
            <template slot="md-autocomplete-item" slot-scope="{ item, term }">
              <md-icon>{{ item.group == 'article' ? 'create' : (item.group == 'document' ? 'chrome_reader_mode' : 'forum') }}</md-icon>
              <md-highlight-text :md-term="term">{{ extractDescription(item.title, 40) }}</md-highlight-text>
            </template>
            <template slot="md-autocomplete-empty" slot-scope="{ term }">
              Không tìm thấy kết quả nào cho từ khóa "{{ term }}"
            </template>
          </md-autocomplete>
          
          <div class="md-toolbar-section-end">
            <no-ssr>
              <md-button @click="navbarButtonClick">
                <span class="jg-navbar-button md-xsmall-hide">{{ $store.state.loginStatus ? $store.state.user.name : 'Đăng nhập' }}</span>
                <md-avatar class="md-avatar-icon" v-if="$store.state.loginStatus">
                  <img :src="$store.state.user.avatar" alt="$store.state.user.name" v-if="$store.state.user.avatar != ''" />
                  <span v-else>{{ $store.state.user.name[0] }}</span>
                </md-avatar>
                <md-icon v-else>person</md-icon>
              </md-button>
            </no-ssr>
          </div>
        </div>

        <div class="md-toolbar-row jg-tab-menu" v-if="subMenu.length > 0">
          <md-tabs md-sync-route>
            <md-tab v-for="(subMenuItem, index) in subMenu" :key="index" :md-label="subMenuItem.text" :to="tabLink(subMenuItem.tab)" exact></md-tab>
          </md-tabs>
        </div>
      </md-app-toolbar>
    
      <md-app-content>
        <nuxt />
      </md-app-content>
    </md-app>

    <md-drawer class="md-right" :v-if="$store.state.loginStatus" :md-active.sync="showSidepanel">
      <div class="jg-account-settings-wrapper">
        <p class="md-title md-primary-color">Tài khoản</p>
        <md-list>
          <md-list-item :to="`/ho-so?id=${$store.state.user.id}`" @click="showSidepanel = false">
            <span class="md-list-item-text">Hồ sơ</span>
          </md-list-item>

          <md-list-item :to="`/ho-so?id=${$store.state.user.id}&tab=tai-lieu`" @click="showSidepanel = false">
            <span class="md-list-item-text">Tài liệu</span>
          </md-list-item>

          <md-list-item :to="`/ho-so?id=${$store.state.user.id}&tab=bai-viet`" @click="showSidepanel = false">
            <span class="md-list-item-text">Bài viết</span>
          </md-list-item>

          <md-list-item :to="`/ho-so?id=${$store.state.user.id}&tab=thao-luan`" @click="showSidepanel = false">
            <span class="md-list-item-text">Thảo luận</span>
          </md-list-item>

          <md-list-item>
            <span class="md-list-item-text">Đăng xuất</span>
            <md-button class="md-icon-button md-list-action" @click="logOut">
              <md-icon class="md-accent">exit_to_app</md-icon>
            </md-button>
          </md-list-item>
        </md-list>
      </div>

      <md-divider></md-divider>

      <div class="jg-notifications-wrapper">
        <p class="md-title md-primary-color">Thông báo</p>
        <md-list class="jg-notifications md-scrollbar">
          <div class="jg-notifications-empty" v-if="notifications.length == 0">
            <md-icon>airplanemode_active</md-icon>
            <p>Không có thông báo</p>
          </div>
          <nuxt-link v-else v-for="(notification, index) in notifications" :key="index" :to="notification.route">
            <md-list-item  v-html="notification.text" />
          </nuxt-link>
        </md-list>
      </div>
    </md-drawer>

    <md-snackbar md-position="center" :md-duration="4000" :md-active.sync="showSnackbar" md-persistent>
      <span>{{ $store.state.snackbarMessage }}</span>
      <md-button class="md-primary" @click="showSnackbar = false">Đóng</md-button>
    </md-snackbar>
  </div>
</template>

<script>
  import LoginForm from '~/components/layouts/default/LoginForm'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  import extractDescription from '~/assets/mixins/extractDescription.js'

  export default {
    components: {
      LoginForm
    },

    mixins: [localStorageManager, extractDescription],

    created() {
      if (process.client) {
        if (localStorage.jwt_token != undefined) {
          let payload = this.getPayload()
          if (payload != null) {
            payload.likedArticles = this.getLikedArticles()
            payload.likedVocas = this.getLikedVocas()
            payload.likedGrammars = this.getLikedGrammars()
            payload.likedKanji = this.getLikedKanjiList()
        
            this.$store.commit('setLoginStatus', true)
            this.$store.commit('user/setUser', payload)
          }
        }
      } 
    },
    data: () => ({
      showNavigation: false,
      showSidepanel: false,
      upperMenuItems: [
        {
          text: "Từ vựng",
          icon: "translate",
          route: "/tu-vung"
        },
        {
          text: "Ngữ pháp",
          icon: "color_lens",
          route: "/ngu-phap"
        },
        {
          text: "Chữ Hán",
          icon: "style",
          route: "/chu-han"
        },
        {
          text: "Thi thử",
          icon: "fitness_center",
          route: "/thi-thu"
        }
      ],
      lowerMenuItems: [
        {
          text: "Tài liệu",
          icon: "chrome_reader_mode",
          route: "/tai-lieu"
        },
        {
          text: "Bài viết",
          icon: "create",
          route: "/bai-viet"
        },
        {
          text: "Thảo luận",
          icon: "forum",
          route: "/thao-luan"
        },
        {
          text: "Giới thiệu",
          icon: "info",
          route: "/gioi-thieu"
        }
      ],
      socialMenuItems: [
        {
          text: "Jango trên Facebook",
          icon: "f",
          route: "https://www.facebook.com/jangodotvn"
        }
      ],
      subMenuItems: [
        {
          name: "bai-viet",
          items: [
            {
              text: "Tất cả",
              tab: ""
            },
            {
              text: "Kiến thức",
              tab: "kien-thuc"
            },
            {
              text: "Chia sẻ",
              tab: "chia-se"
            },
            {
              text: "Khác",
              tab: "khac"
            }
          ]
        },
        {
          name: "tai-lieu",
          items: [
            {
              text: "Tất cả",
              tab: ""
            },
            {
              text: "Giáo trình",
              tab: "giao-trinh"
            },
            {
              text: "Bài tập",
              tab: "bai-tap"
            },
            {
              text: "Khác",
              tab: "khac"
            }
          ]
        },
        {
          name: "ho-so",
          items: [
            {
              text: "Hồ sơ",
              tab: ""
            },
            {
              text: "Tài liệu",
              tab: "tai-lieu"
            },
            {
              text: "Bài viết",
              tab: "bai-viet"
            },
            {
              text: "Thảo luận",
              tab: "thao-luan"
            }
          ]
        },
        {
          name: "ngu-phap",
          items: [
            {
              text: "N5",
              tab: ""
            },
            {
              text: "N4",
              tab: "n4"
            },
            {
              text: "N3",
              tab: "n3"
            },
            {
              text: "N2",
              tab: "n2"
            },
            {
              text: "N1",
              tab: "n1"
            },
            {
              text: "Yêu thích",
              tab: "yeu-thich"
            }
          ]
        },
        {
          name: "tu-vung-slug",
          items: [
            {
              text: "Tất cả",
              tab: ""
            },
            {
              text: "Yêu thích",
              tab: "yeu-thich"
            }
          ]
        },
        {
          name: "tu-vung-kana",
          items: [
            {
              text: "Hiragana",
              tab: ""
            },
            {
              text: "Katakana",
              tab: "katakana"
            }
          ]
        },
        {
          name: "chu-han",
          items: [
            {
              text: "N5",
              tab: ""
            },
            {
              text: "N4",
              tab: "n4"
            },
            {
              text: "N3",
              tab: "n3"
            },
            {
              text: "N2",
              tab: "n2"
            },
            {
              text: "N1",
              tab: "n1"
            },
            {
              text: "Yêu thích",
              tab: "yeu-thich"
            }
          ]
        },
        {
          name: "gioi-thieu",
          items: [
            {
              text: "FAQ",
              tab: ""
            },
            {
              text: "Về Jango",
              tab: "ve-jango"
            }
          ]
        },
        {
          name: "thao-luan",
          items: [
            {
              text: "Tất cả",
              tab: ""
            },
            {
              text: "Kiến thức",
              tab: "kien-thuc"
            },
            {
              text: "Bài tập",
              tab: "bai-tap"
            },
            {
              text: "Khác",
              tab: "khac"
            }
          ]
        }
      ],
      otherRoutes: [
        {
          text: 'Sửa bài viết',
          route: '/chinh-sua/bai-viet'
        }
      ],
      notifications: [],
      searchTerm: '',
      searchResults: []
    }),

    computed: {
      currentPage: function() {
        for (var i = 0; i < this.upperMenuItems.length; i++) {
          if (this.$nuxt.$route.path == this.upperMenuItems[i].route || this.$nuxt.$route.path.indexOf(this.upperMenuItems[i].route + '/') == 0) {
            return this.upperMenuItems[i].text
          }
        }

        for (var i = 0; i < this.lowerMenuItems.length; i++) {
          if (this.$nuxt.$route.path == this.lowerMenuItems[i].route || this.$nuxt.$route.path.indexOf(this.lowerMenuItems[i].route + '/') == 0) {
            return this.lowerMenuItems[i].text
          }
        }

        for (var i = 0; i < this.otherRoutes.length; i++) {
          if (this.$nuxt.$route.path == this.otherRoutes[i].route || this.$nuxt.$route.path.indexOf(this.otherRoutes[i].route + '/') == 0) {
            return this.otherRoutes[i].text
          }
        }

        switch(this.$nuxt.$route.name) {
          case 'ho-so':
            return 'Hồ sơ'
            break
          default:
            return 'Jango'
        }
      },

      subMenu: function() {
        for (var i = 0; i < this.subMenuItems.length; i++) {
          if (this.$nuxt.$route.name == this.subMenuItems[i].name) {
            return this.subMenuItems[i].items
          }
        }
  
        return []
      },

      showSnackbar: {
        get() {
          return this.$store.state.showSnackbar
        },
        set(show) {
          this.$store.commit("toggleSnackbar", show)
        }
      }
    },

    methods: {
      menuItemClass: function(route) {
        return ((this.$nuxt.$route.path == route || this.$nuxt.$route.path.indexOf(route + '/') == 0) ? 'md-primary-color' : '')
      },

      triggerResize: function() {
        setTimeout(function() {
          var event = new Event('resize');
          window.dispatchEvent(event);
        }, 0);
      },

      navbarButtonClick: function() {
        if (this.$store.state.loginStatus) {
          this.showSidepanel = true
        } else {
          this.$store.commit('toggleLoginForm', true)
        }
      },

      tabLink(tab) {
        let query = {}
        
        for (let param in this.$route.query) {
          query[param] = this.$route.query[param]
        }
       
        if (tab == '') {
          delete query.tab
        } else {
          query.tab = tab
        }
        let queryString = ''
        for (let param in query) {
          queryString += `&${param}=${query[param]}`
        }
        if (queryString != '') {
          queryString = `?${queryString.slice(1)}`
        }
    
        return this.$route.path + queryString
      },

      logOut: function() {
        this.$store.commit('setLoginStatus', false)
        this.removeJwt()
        this.removeLikedArticles()
        this.removeLikedVocas()
        this.removeLikedGrammars()
        this.removeLikedKanjiList()
        this.removeLikedThreads()
        this.removeLikedReplies()

        this.$store.commit('user/resetUser')
        this.$store.commit('setSnackbarMessage', "Đã đăng xuất")
        this.$store.commit('toggleSnackbar', true)
        this.showSidepanel = false
      },

      getResults() {
        this.triggerResize()
        let that = this
        this.searchResults = new Promise(resolve => {
          that.$axios.get(`${process.env.apiUrl}/search/list?s=${that.searchTerm}`)
            .then((res) => {
              resolve(res.data.results)
            })
            .catch((err) => {
            })
        })
      },

      chooseResult(path, group) {
        this.searchTerm = ''
        this.$router.push(`/${group == 'article' ? 'bai-viet' : (group == 'document' ? 'tai-lieu' : 'thao-luan')}/${path}`)
      }
    }
  }
</script>

<style lang="scss" scoped>
  .page-container {
    overflow: hidden;
    position: relative;

    .md-app {
      height: 100vh;
      .md-app-drawer {
        .md-list {
          .jg-social-menu-item {
            .md-avatar {
              margin-right: 32px;
            }
          }
        }
      }

      .md-app-toolbar {
        padding: 0;
        .jg-top-bar {
          margin: 0 16px;
          .md-title {
            @media (max-width: 960px) {
              & {
                margin-left: 6px;
                margin-right: 20px;
              }
            }
          }
          .search {
            max-width: 500px;
            /deep/ .md-input {
              @media (max-width: 599px) {
                & {
                  padding-right: 16px;
                }
              }
            }
          }
          .md-toolbar-section-end {
            position: relative;
            .jg-navbar-button {
              @media (min-width: 601px) {
                & {
                  white-space: nowrap;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  max-width: 120px;
                  display: inline-block;
                  position: relative;
                  top: 6px;
                }
              }

              @media (min-width: 961px) {
                & {
                  max-width: 200px;
                }
              }
            }

            .md-avatar {
              width: 36px;
              min-width: 36px;
              height: 36px;
              span {
                font-size: 20px;
              }
            } 

            .md-button {
              min-width: auto;
              margin-left: 8px;
              /deep/ .md-ripple {
                padding: 0;
              }
            }
          }
        }

        .jg-tab-menu {
          min-height: auto;
          background: white;
          .md-tabs {
            justify-content: center;
            flex-direction: row;
            width: 100%;
            overflow-x: scroll;
            -ms-overflow-style: none;  // IE 10+
            scrollbar-width: none;  // Firefox
            &::-webkit-scrollbar { 
              display: none;  // Safari and Chrome
            }
            @media (max-width: 450px) {
              & {
                justify-content: initial;
              }
            }

            /deep/ .md-button-content {
              color: var(--md-theme-default-primary-on-background, #448aff) !important;
            }
          }

        }
      }

      /deep/ .md-app-scroller {
        @media (max-width: 599px) {
          & {
            padding-left: 0 !important;
          }
        }
        @media (min-width: 600px) {
          & {
            padding-left: 70px !important;
          }
        }

        .md-app-content {
        padding: 48px;
        @media (max-width: 960px) {
          padding: 36px 18px;
        }
      }
      }  
    }

    .md-right {
      height: 100vh;
      overflow: hidden;
      flex-direction: column;
      p {
        padding: 0 16px;
        margin-bottom: 0;
      }

      .jg-notifications {
        max-height: calc(100vh - 290px);
        overflow-y: auto;
        a {
          &:hover {
            background-color: #E0E0E0;
          }
          .md-list-item {
            margin: 8px 8px 8px 16px;
            color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
          }
        }

        .jg-notifications-empty {
          text-align: center;
          .md-icon {
            margin-top: 16px;
            font-size: 64px !important;
            width: 72px;
            height: 72px;
            transform: rotate(45deg);
          }
        }
      }

      .md-list:not(.jg-notifications) {
        /deep/ .md-list-item-content {
          height: 36px;
          min-height: auto;
        }
      }
    }

    .md-drawer {
      width: 230px;
      max-width: calc(100vw - 125px);
      a:hover {
        text-decoration: none !important;
      }

      &.md-persistent-mini {
        @media (max-width: 599px) {
          position: absolute;
          &:not(.md-active) {
            display: none;
          }
        }
      }
    }
  }

  .md-autocomplete-items {
    .color {
      width: 16px;
      height: 16px;
      display: inline-block;
      margin-right: 16px;
      border: 1px solid rgba(#000, .12);
    }

  }

  .md-snackbar {
    z-index: 99999;
  }
</style>
