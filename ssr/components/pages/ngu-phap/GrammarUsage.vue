<template>
  <md-table-row>
    <md-table-cell>
      {{ usage.desc }}
    </md-table-cell>
    <md-table-cell>
      <div>
        <div class="jg-example-wrapper" v-for="(example, index) in usage.examples" :key="index">
          <div class="jg-voca-speaker" @click="speak(example['example_text'])"><md-icon >volume_up</md-icon></div>
          <div>
            <div class="md-primary-color md-subheading">{{example['example_text']}}</div>
            <div>{{example['example_meaning']}}</div>
          </div>
        </div>
      </div>
    </md-table-cell>
  </md-table-row>
</template>

<script>
  import speak from '~/assets/mixins/speak.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'

  export default {
    mixins: [speak],
    props: {
      usage: Object
    },
    data: function(){
      return {
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-table-row {
    .md-table-cell {
      &:first-child {
        width: 30%;
      }
      /deep/ .md-table-cell-container {
        padding: 0px 16px;
        .jg-example-wrapper {
          padding: 6px 0;
        }
      }
    }
  }
</style>


