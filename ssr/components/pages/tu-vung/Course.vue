<template>
  <md-card :class="accent ? 'md-accent' : 'md-primary'">
    <md-card-header>
      <md-card-header-text>
        <div class="md-title">{{ data.title }}</div>
        <div class="md-subhead">{{ data.description }}</div>
      </md-card-header-text>

      <md-card-media md-small>
        <img :src="data.icon" :alt="data.title">
      </md-card-media>
    </md-card-header>
    
    <md-list :md-expand-single="false">
      <md-list-item md-expand :md-expanded.sync="expandLessons">
        <md-icon class="md-xsmall-hide">assignment</md-icon>
        <span class="md-list-item-text">Danh sách bài học</span>

        <md-list slot="md-expand">
          <nuxt-link class="md-list-item-container" v-for="(lesson, index) in data.lessons" :key="index" :to="`/tu-vung/${lesson.route}`">
            <md-list-item class="md-inset">
              {{ lesson.title }}
              <div class="md-caption" v-if="lesson.favorite == undefined">
                {{ lesson.wordCount }} từ vựng
              </div>
              <div class="md-caption" v-else>
                <span v-if="lesson.wordCount != undefined" class="md-xsmall-hide">{{ lesson.wordCount }} từ vựng - </span>{{ lesson.favorite }} từ yêu thích
              </div>
            </md-list-item>
          </nuxt-link>
        </md-list>
      </md-list-item>
    </md-list>
  </md-card>
  
</template>

<script>
  export default {
    props: {
      data: Object,
      accent: Boolean
    },
    data() {
      return {
        expandLessons: false
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
      /deep/ .md-list-item-content {
        @media (max-width: 599px) {
          padding: 4px 16px;
        }
      }
    }
  }
</style>