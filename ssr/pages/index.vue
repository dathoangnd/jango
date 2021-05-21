<template>
  <div class="md-layout md-gutter">
    <div class="md-layout-item md-size-40 md-small-hide md-xsmall-hide">
      <img class="jg-homepage-image" src="/images/pages/gopher-ninja.png" alt="Jango" />
    </div>
    <div class="md-layout-item md-size-60 md-small-size-100 md-xsmall-size-100">
      <h3 class="md-display-1 md-primary-color">Học tiếng Nhật miễn phí</h3>
      <div class="md-layout md-gutter jg-modules">
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/tu-vung">
            <img src="/images/pages/vocabulary.svg" alt="Từ vựng" />
          </nuxt-link>
          <p class="md-subheading">Từ vựng</p>
        </div>
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/ngu-phap">
            <img src="/images/pages/grammar.svg" alt="Ngữ pháp" />
          </nuxt-link>
          <p class="md-subheading">Ngữ pháp</p>
        </div>
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/chu-han">
            <img src="/images/pages/kanji.svg" alt="Chữ Hán" />
          </nuxt-link>
          <p class="md-subheading">Chữ Hán</p>
        </div>
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/thi-thu">
            <img src="/images/pages/exam.svg" alt="Thi thử" />
          </nuxt-link>
          <p class="md-subheading">Thi thử</p>
        </div>
      </div>

      <h3 class="md-display-1 md-primary-color">Tham gia cộng đồng Jango</h3>
      <div class="md-layout md-gutter jg-modules">
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/tai-lieu">
            <img src="/images/pages/document.svg" alt="Tài liệu" />
          </nuxt-link>
          <p class="md-subheading">Tài liệu</p>
        </div>
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/bai-viet">
            <img src="/images/pages/article.svg" alt="Bài viết" />
          </nuxt-link>
          <p class="md-subheading">Bài viết</p>
        </div>
        <div class="md-layout-item md-size-25">
          <nuxt-link to="/thao-luan">
            <img src="/images/pages/discussion.svg" alt="Thảo luận" />
          </nuxt-link>
          <p class="md-subheading">Thảo luận</p>
        </div>
        <div class="md-layout-item md-size-25">
          <a href="https://www.facebook.com/jangodotvn" target="_blank">
            <img src="/images/pages/fanpage.svg" alt="Fanpage" />
          </a>
          <p class="md-subheading">Fanpage</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {

  }
</script>

<style lang="scss" scoped>
  .jg-homepage-image {
    width: 100%;
  }
  
  .jg-modules {
    .md-layout-item {
      text-align: center;
      img {
        width: 60%;
      }
    }
  }
</style>
