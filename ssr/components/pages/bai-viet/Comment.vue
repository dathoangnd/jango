<template>
  <div>
    <div class="jg-comment-container">
      <div class="jg-avatar-wrapper">
        <md-avatar class="md-avatar-icon md-accent">
          <img :src="comment.userAvatar" :alt="comment.userName" v-if="comment.userAvatar != ''" />
          <span v-else>{{ comment.userName[0] }}</span>
        </md-avatar>
      </div>
      <div class="jg-comment-wrapper">
        <div class="md-subheading"><nuxt-link :to="`/ho-so?id=${comment.userId}`">{{ comment.userName }}</nuxt-link></div>
        <div class="md-caption">{{ calculateTimeDuration(comment.updatedAt) }}</div>
        <div class="md-body-1">{{ comment.content }}</div>
        <div class="jg-comment-actions">
          <span class="md-primary-color" v-if="$store.state.user.id != null && ($store.state.user.id == comment.userId || $store.state.user.id == adminId)" @click="$emit('deletecomment', comment.id)">Xóa</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import calculateTimeDuration from '~/assets/mixins/calculateTimeDuration.js'
  import localStorageManager from '~/assets/mixins/localStorageManager.js'
  export default {
    mixins: [calculateTimeDuration, localStorageManager],
    props: {
      comment: Object
    },
    computed: {
      adminId() {
        return process.env.adminUserId
      }
    }
  }
</script>


<style lang="scss" scoped>
  .jg-comment-container {
    display: flex;
    margin-bottom: 10px;

    .jg-avatar-wrapper {
      flex-basis: 60px;
    }

    .jg-comment-wrapper {
      flex-grow: 1;

      .jg-comment-actions {
        font-size: 12px;
        > span {
          display: inline-block;
          margin-right: 5px;
          cursor: pointer;
        }
      }
    }
  }
</style>

