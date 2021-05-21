<template>
  <md-card class="md-primary">
    <md-card-header>
      <md-card-header-text>
        <div class="md-title">Đề thi thử {{ data.title.toUpperCase() }}</div>
        <div class="md-subhead">{{ data.description }}</div>
      </md-card-header-text>
    </md-card-header>
    
    <md-list :md-expand-single="false">
      <md-list-item md-expand :md-expanded.sync="expandTests">
        <md-icon class="md-xsmall-hide">all_inbox</md-icon>
        <span class="md-list-item-text">Danh sách đề thi</span>

        <md-list slot="md-expand">
          <div class="md-list-item-container" v-for="(test, index) in data.tests" :key="index" @click="$emit('open', data.title, index)">
            <md-list-item class="md-inset">
              Đề {{ index + 1 }}
              <div class="md-caption">
                {{ test.score < 0 ? 'Chưa thi' : 'Điểm của bạn: ' + test.score }}
              </div>
            </md-list-item>
          </div>
        </md-list>
      </md-list-item>
    </md-list>
  </md-card>
  
</template>

<script>
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'

  export default {
    mixins: [calculateTimeDuration],
    props: {
      data: Object
    },
    data() {
      return {
        expandTests: false
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-card-header {
    display: flex;
  }

  .md-list {
    padding: 0;
    .md-list-item.md-inset {
      cursor: pointer;
      /deep/ .md-list-item-content {
        @media (max-width: 599px) {
          padding: 4px 16px;
        }
      }
    }
  }
</style>