<template>
  <div class="md-layout">
    <div class="md-layout-item md-size-70 md-small-size-100">
      <div class="md-subhead jg-breadcrumb breadcrumb">
        <ol itemscope itemtype="http://schema.org/BreadcrumbList">
          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <nuxt-link to="/" itemprop="item"><span itemprop="name">Jango</span></nuxt-link>
            <meta itemprop="position" content="1" />
          </li>

          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <nuxt-link to="/tai-lieu" itemprop="item"><span itemprop="name">Tài liệu</span></nuxt-link>
            <meta itemprop="position" content="2" />
          </li>

          <li itemprop="itemListElement" itemscope itemtype="http://schema.org/ListItem">
            <template v-if="document.categories != ''">
              <nuxt-link v-for="(category, index) in document.categories.trim().split(' ')" :key="index" :to="`/tai-lieu?tab=${category}`" itemprop="item"><span itemprop="name">{{ category == 'giao-trinh' ? (document.categories.indexOf('bai-tap') != -1 && index == 0 ? 'Giáo trình, ' : 'Giáo trình') : (document.categories.indexOf('giao-trinh') != -1 && index == 0 ? 'Bài tập, ' : 'Bài tập') }}</span></nuxt-link>
            </template>
            <nuxt-link v-else to="/tai-lieu?tab=khac" itemprop="item"><span itemprop="name">Khác</span></nuxt-link>
            <meta itemprop="position" content="3" />
          </li>
        </ol>
      </div>
      <h1 class="md-title md-primary-color" itemprop="name">{{ document.title }}</h1>
      <div class="md-subhead"><nuxt-link :to="`/ho-so?id=${document.userId}`" itemprop="author">{{ document.userName }}</nuxt-link> &bull; <time itemprop="dateModified" :datetime="updatedAtSchemaData">{{ calculateTimeDuration(document.updatedAt) }}</time></div>
      <br />

      <div class="md-layout jg-document-info">
        <div class="md-layout-item md-size-40 md-small-size-100 md-xsmall-size-100">
          <img :src="document.avatar != '' ? document.avatar : defaultAvatar" :alt="document.title" />
        </div>
        <div class="md-layout-item md-size-55 md-small-size-100 md-xsmall-size-100">
          <md-table class="md-scrollbar">
            <md-table-row>
              <md-table-cell><b>Tác giả</b></md-table-cell>
              <md-table-cell>{{ document.author != '' ? document.author : 'Chưa xác định' }}</md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell><b>Nguồn</b></md-table-cell>
              <md-table-cell>{{ document.source != '' ? document.source : 'Chưa xác định' }}</md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell><b>Lượt tải</b></md-table-cell>
              <md-table-cell>{{ document.download }}</md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell>
                <a :href="document.file" target="_blank" download>
                  <md-button class="md-raised md-primary">Tải xuống</md-button>
                </a>
              </md-table-cell>
              <md-table-cell>
                <md-button class="md-raised md-accent">Chia sẻ</md-button>
              </md-table-cell>
            </md-table-row>
          </md-table>
        </div>
      </div>
      <md-divider></md-divider>
      <br />
      <div class="jg-document-content" v-html="document.content" />
      <br /><br /><br />
    </div>

    <div class="md-layout-item md-size-30 md-small-size-100 jg-newest-documents-wrapper">
      <h4>Tài liệu mới nhất</h4>
      <div class="jg-newest-documents">
        <div class="jg-newest-document" v-for="(document, index) in newestDocuments" :key="index">
          <div class="jg-newest-document-image">
            <img :src="document.avatar != '' ? document.avatar : defaultAvatar" :alt="document.title" />
          </div>
          <div class="jg-newest-document-summary">
            <nuxt-link :to="`/tai-lieu/${document.slug}-${document.id}`"><h3>{{ document.title }}</h3></nuxt-link>
            <p>{{document.userName}}</p>
          </div> 
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import DownloadStats from '~/components/pages/tai-lieu/DownloadStats'
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'

  export default {
    mixins: [calculateTimeDuration],
    components: {
      DownloadStats
    },
    asyncData({ $axios, params, error }) {
      let id = parseInt(params.slug.slice(params.slug.lastIndexOf('-') + 1))
      if (id == undefined || id == '' || id == NaN) {
        error({ statusCode: 404, message: 'Trang không tồn tại' })
        return
      }

      return $axios.get(`${process.env.apiUrl}/document/single?id=${id}&newest=5`)
        .then((res) => {
          return res.data
        })
        .catch(e => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    computed: {
      updatedAtSchemaData: function() {
        let updatedDate = new Date(this.document.updatedAt);
        return `${updatedDate.getFullYear()}-${updatedDate.getMonth() + 1}-${updatedDate.getDate()}`
      },
      defaultAvatar() {
        return process.env.defaultDocumentAvatar
      }
    }
  }
</script>


<style lang="scss" scoped>
  .md-layout {
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

    h1.md-title {
      font-size: 28px;
      margin: 10px 0;
    }

    .md-subhead {
      a {
        color: var(--md-theme-default-text-primary-on-background, rgba(0, 0, 0, 0.87));
      }
    }

    .jg-document-info {
      justify-content: space-between;
      .md-button {
        width: 100%;
        margin: 6px 0;
      }
    }

    .jg-reactions {
      padding-left: 16px;
    }

    .jg-newest-documents-wrapper {
      padding-left: 50px;
      h4 {
        margin-top: 0;
        font-size: 18px;
      }

      .jg-newest-document {
        display: flex;
        margin-bottom: 25px;
        .jg-newest-document-image {
          flex-basis: 25%;
          img {
            width: 100%;
          }
        }
        .jg-newest-document-summary {
          flex-basis: 75%;
          padding-left: 10px;
          h3 {
            margin: 0;
            font-size: 14px;
          }
          p {
            margin: 0;
            font-size: 12px;
          }
        }
      }
    }
  }
  
</style>
