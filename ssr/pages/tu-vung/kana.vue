<template>
<div class="md-layout md-gutter">
      <div class="md-layout-item md-xlarge-size-50 md-large-size-50 md-medium-size-100 md-small-size-100 md-xsmall-size-100">
        <div class="md-layout">
          <div class="md-layout-item md-size-20" v-for="(voca, index) in (tab == 'katakana' ? katakana[0] : hiragana[0])" :key="index" @click="speak(voca.text)">
            <Letter :text="voca.text" :read="voca.read" /> 
          </div>
        </div>
      </div>
      <div class="md-layout-item md-xlarge-size-50 md-large-size-50 md-medium-size-100 md-small-size-100 md-xsmall-size-100">
        <div class="md-layout">
          <div class="md-layout-item md-size-20" v-for="(voca, index) in (tab == 'katakana' ? katakana[1] : hiragana[1])" :key="index" @click="speak(voca.text)">
            <Letter :text="voca.text" :read="voca.read" /> 
          </div>
        </div>
      </div>
    </div>
</template>

<script>
  import Letter from '~/components/pages/tu-vung/Letter.vue'
  import speak from '~/assets/mixins/speak.js'

  export default {
    mixins: [speak],
    components: {
      Letter
    },
    asyncData({ $axios, params, error }) {
      return $axios.get(`${process.env.apiUrl}/static/data/vocabularies/kana.json`)
        .then((res) => {
          return {
            hiragana: res.data.hiragana,
            katakana: res.data.katakana
          }
        })
        .catch(e => {
          error({ statusCode: 404, message: 'Trang không tồn tại' })
        })
    },
    data() {
      return {
        tab: (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    },
    watch: {
      '$route.query.tab': function() {
        this.tab = (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    }
  }
</script>

<style lang="scss" scoped>

</style>


