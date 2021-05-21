<template>
  <div class="md-layout">
    <div class="md-layout-item md-size-100" v-show="tab == ''">
      <md-list>
        <md-list-item md-expand v-for="(faq, index) in faqs" :key="index">
          <md-icon>help</md-icon>
          <span class="md-list-item-text">{{ faq.question }}</span>
          <div class="jg-new-article-meta-slot" slot="md-expand">
            <div class="jg-faq-answer" v-html="faq.answer" />
          </div>
        </md-list-item>
      </md-list>
    </div>
    <div class="md-layout-item md-size-100" v-show="tab == 've-jango'">
      <md-empty-state
        class="md-primary"
        md-icon="devices_other"
        md-label="Jango"
        md-description="Nền tảng học tiếng Nhật miễn phí phát triển vì cộng đồng">
        <a href="https://www.facebook.com/jangodotvn" target="_blank">
          <md-button class="md-primary md-raised">Chúng tôi trên Facebook</md-button>
        </a>
      </md-empty-state>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        faqs: [
          {
            question: "Jango là gì?",
            answer: "Jango là nền tảng học tiếng Nhật miễn phí và là nơi cộng đồng Jango tham gia viết bài, thảo luận về việc học tiếng Nhật."
          },
          {
            question: "Ai tạo ra Jango?",
            answer: "Jango được tạo ra bởi đội ngũ Jango. Sự phát triển của Jango phụ thuộc rất lớn vào đóng góp của cộng đồng thành viên mà bạn là một trong số đó."
          },
          {
            question: "Jango miễn phí?",
            answer: "Jango là hoàn toàn miễn phí. Bạn không cần trả bất cứ khoản phí nào để tiếp cận mọi nội dung trên Jango."
          },
          {
            question: "Tôi học được gì từ Jango?",
            answer: "Bạn có thể học từ vựng, ngữ pháp ở mọi cấp độ. Ngoài ra bạn cũng có thể tìm tài liệu và làm bài thi thử trên Jango."
          },
          {
            question: "Làm thế nào để tham gia cộng đồng Jango?",
            answer: "Việc đăng ký tài khoản chỉ trong vòng 1 phút và bạn có thể tham gia học, viết bài, thảo luận với cộng đồng Jango. Bạn cũng có thể kết nối với mọi người trên trang Facebook."
          },
          {
            question: "Quyền riêng tư của tôi trên Jango?",
            answer: "Thông tin bạn chia sẻ sẽ được hiển thị công khai với mọi người ghé thăm hồ sơ của bạn. Hãy đảm bảo không chia sẻ điều gì mà bạn cho là riêng tư. Chúng tôi cũng đảm bảo không bao giờ sử dụng thông tin của bạn cho mục đích lợi nhuận."
          },
          {
            question: "Jango cho thiết bị di động?",
            answer: "Jango được thiết kế tối ưu cho các thiết bị di động. Bạn có thể thêm Jango vào màn hình chính để truy cập Jango nhanh hơn từ bất cứ đâu."
          },
          {
            question: "Làm thế nào để liên lạc với nhà phát triển Jango?",
            answer: "Nếu bạn có góp ý, thắc mắc hoặc muốn giúp đỡ chúng tôi thì có thể liên lạc với đội ngũ phát triển thông qua trang Facebook của Jango."
          }
        ]
      }
    },
    computed: {
      tab() {
        return (this.$route.query.tab != undefined ? this.$route.query.tab : '')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .md-list-item {
    .md-list-item-expand {
      border-bottom: 1px solid var(--md-theme-default-divider-on-background, rgba(0, 0, 0, 0.12));
    }
    .jg-faq-answer {
      padding: 16px 40px;
      white-space: initial;
      line-height: 150%;
    }
  }

  .md-empty-state {
    max-width: 320px;
  }
</style>
