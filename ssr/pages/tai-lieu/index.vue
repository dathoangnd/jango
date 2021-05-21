<template>
  <div class="md-layout">
    <div class="md-layout-item md-xlarge-size-25 md-large-size-25 md-medium-size-33 md-small-size-50 md-xsmall-size-100" v-for="(document, index) in filteredDocuments" :key="index">
      <Document :document="document" />
    </div>
    <FabButton icon="cloud_upload" to="/chinh-sua/tai-lieu" />
  </div>
</template>

<script>
  import Document from '~/components/pages/tai-lieu/Document'
  import FabButton from '~/components/pages/shared/FabButton'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    components: {
      Document,
      FabButton
    },
    mixins: [localStorageManager],
    asyncData({ $axios }) {
      return $axios.get(`${process.env.apiUrl}/document/list?publish=1`)
        .then((res) => {
          return {
            documents: res.data.documents
          }
        })
    },
    data: function() {
      return {
        category: (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    },
    computed: {
      filteredDocuments() {
        let that = this
        return this.documents.filter(function(document) {
          if (that.category == 'khac')
            return document.categories.indexOf('giao-trinh') == -1 && document.categories.indexOf('bai-tap') == -1
          return document.categories.indexOf(that.category) != -1
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