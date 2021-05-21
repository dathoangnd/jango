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
        <h1 class="md-title md-primary-color">{{course}}000 từ vựng thiết yếu - Bài {{lesson}}</h1>
      </md-table-toolbar>

      <Voca v-for="(voca, index) in vocabularies" :key="index" :voca="voca" :order="(course - 1) * 1000 + (lesson - 1) * 100 + index + 1" />
    </md-table>
  </div>
</template>

<script>
  import Voca from '~/components/pages/tu-vung/Voca.vue'

  export default {
    components: {
      Voca
    },
    validate({ params }) {
      let slugParts = params.slug.split('-')
      return (slugParts.length == 3 && slugParts[1] != '' && slugParts[2] != '')
    },
    asyncData({ $axios, params, error }) {
      let slugParts = params.slug.split('-')
      return $axios.get(`${process.env.apiUrl}/vocabulary/list?course=${slugParts[1]}&lesson=${slugParts[2]}`)
        .then((res) => {
          return {
            vocabularies: res.data.vocas
          }
        })
        .catch(e => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    data() {
      return {
        course: parseInt(this.$route.params.slug.split('-')[1]),
        lesson: parseInt(this.$route.params.slug.split('-')[2])
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
