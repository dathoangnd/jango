<template>
  <div class="md-layout">
    <div class="md-layout-item md-xlarge-size-25 md-large-size-25 md-medium-size-33 md-small-size-50 md-xsmall-size-100" v-for="(article, index) in filteredArticles" :key="index">
      <Article :article="article" />
    </div>
    <FabButton icon="edit" to="/chinh-sua/bai-viet" />
  </div>
</template>

<script>
  import Article from '~/components/pages/bai-viet/Article'
  import FabButton from '~/components/pages/shared/FabButton'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    components: {
      Article,
      FabButton
    },
    mixins: [localStorageManager],
    asyncData({ $axios }) {
      return $axios.get(`${process.env.apiUrl}/article/list?publish=1`)
        .then((res) => {
          return {
            articles: res.data.articles
          }
        })
    },
    data: function() {
      return {
        category: (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    },
    computed: {
      filteredArticles() {
        let that = this
        return this.articles.filter(function(article) {
          if (that.category == 'khac')
            return article.categories.indexOf('kien-thuc') == -1 && article.categories.indexOf('chia-se') == -1
          return article.categories.indexOf(that.category) != -1
        })
      }
    },
    watch: {
      '$route.query.tab': function() {
        this.category = (this.$route.query.tab != undefined ? this.$route.query.tab : '')
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
</style>
