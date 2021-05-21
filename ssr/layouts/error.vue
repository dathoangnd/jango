<template>
  <div>
    <md-empty-state
      md-icon="sentiment_very_dissatisfied"
      :md-label="error.message"
      md-description="Tiếp tục xem các nội dung khác trên Jango">
      <nuxt-link to="/"><md-button class="md-primary md-raised">Về trang chủ</md-button></nuxt-link>
    </md-empty-state>
  </div>
</template>

<script>
export default {
  props: ['error']
}
</script>