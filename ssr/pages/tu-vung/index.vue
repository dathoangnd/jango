<template>
  <div class="md-layout">
    <div class="md-layout-item md-xlarge-size-100 md-large-size-100 md-medium-size-100 md-small-size-100 md-xsmall-size-100" v-for="(course, index) in courses" :key="index">
      <Course :data="course" />
    </div>
    <div class="md-layout-item md-xlarge-size-100 md-large-size-100 md-medium-size-100 md-small-size-100 md-xsmall-size-100">
      <Course :data="favoriteList" :accent="true" />
    </div>
  </div>
</template>

<script>
  import Course from '~/components/pages/tu-vung/Course'

  export default {
    components: {
      Course
    },

    created() {
      for (let i = 0; i < this.$store.state.user.likedVocas.length; i++) {
        let order = this.$store.state.user.likedVocas[i] - 1
        this.courses[parseInt(order / 1000) + 1].lessons[parseInt((order % 1000) / 100)].favorite++
      }
    },
    
    data: function() {
      return {
        courses: [
          {
            title: "Hiragana và Katakana",
            icon: "/images/courses/course-0.svg",
            description: "Bước khởi đầu trên hành trình học tiếng Nhật",
            lessons: [
              {
                title: "Hiragana",
                route: "kana",
                wordCount: 104
              },
              {
                title: "Katakana",
                route: "kana?tab=katakana",
                wordCount: 104
              }
            ]
          },
          {
            title: "1000 từ vựng tiếng Nhật thiết yếu",
            icon: "/images/courses/course-1.svg",
            description: "Từ vựng cơ bản trong giao tiếp hàng ngày",
            lessons: [
              {
                title: "Bài 1",
                route: "bai-1-1",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 2",
                route: "bai-1-2",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 3",
                route: "bai-1-3",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 4",
                route: "bai-1-4",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 5",
                route: "bai-1-5",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 6",
                route: "bai-1-6",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 7",
                route: "bai-1-7",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 8",
                route: "bai-1-8",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 9",
                route: "bai-1-9",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 10",
                route: "bai-1-10",
                wordCount: 100,
                favorite: 0
              }
            ]
          },
          {
            title: "2000 từ vựng tiếng Nhật thiết yếu",
            icon: "/images/courses/course-2.svg",
            description: "Bao phủ hầu hết các cuộc hội thoại thông thường và phát triển kỹ năng đọc văn bản đơn giản",
            lessons: [
              {
                title: "Bài 1",
                route: "bai-2-1",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 2",
                route: "bai-2-2",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 3",
                route: "bai-2-3",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 4",
                route: "bai-2-4",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 5",
                route: "bai-2-5",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 6",
                route: "bai-2-6",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 7",
                route: "bai-2-7",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 8",
                route: "bai-2-8",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 9",
                route: "bai-2-9",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 10",
                route: "bai-2-10",
                wordCount: 100,
                favorite: 0
              }
            ]
          },
          {
            title: "3000 từ vựng tiếng Nhật thiết yếu",
            icon: "/images/courses/course-3.svg",
            description: "Khả năng đọc hiểu văn bản thông thường mà không cần từ điển",
            lessons: [
              {
                title: "Bài 1",
                route: "bai-3-1",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 2",
                route: "bai-3-2",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 3",
                route: "bai-3-3",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 4",
                route: "bai-3-4",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 5",
                route: "bai-3-5",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 6",
                route: "bai-3-6",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 7",
                route: "bai-3-7",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 8",
                route: "bai-3-8",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 9",
                route: "bai-3-9",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 10",
                route: "bai-3-10",
                wordCount: 100,
                favorite: 0
              }
            ]
          },
          {
            title: "4000 từ vựng tiếng Nhật thiết yếu",
            icon: "/images/courses/course-4.svg",
            description: "Giao tiếp độc lập trong hầu hết ngữ cảnh phức tạp",
            lessons: [
              {
                title: "Bài 1",
                route: "bai-4-1",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 2",
                route: "bai-4-2",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 3",
                route: "bai-4-3",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 4",
                route: "bai-4-4",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 5",
                route: "bai-4-5",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 6",
                route: "bai-4-6",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 7",
                route: "bai-4-7",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 8",
                route: "bai-4-8",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 9",
                route: "bai-4-9",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 10",
                route: "bai-4-10",
                wordCount: 100,
                favorite: 0
              }
            ]
          },
          {
            title: "5000 từ vựng tiếng Nhật thiết yếu",
            icon: "/images/courses/course-5.svg",
            description: "Đọc hiểu các tiểu thuyết tiếng Nhật",
            lessons: [
              {
                title: "Bài 1",
                route: "bai-5-1",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 2",
                route: "bai-5-2",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 3",
                route: "bai-5-3",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 4",
                route: "bai-5-4",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 5",
                route: "bai-5-5",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 6",
                route: "bai-5-6",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 7",
                route: "bai-5-7",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 8",
                route: "bai-5-8",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 9",
                route: "bai-5-9",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 10",
                route: "bai-5-10",
                wordCount: 100,
                favorite: 0
              }
            ]
          },
          {
            title: "6000 từ vựng tiếng Nhật thiết yếu",
            icon: "/images/courses/course-6.svg",
            description: "Xin chào tiền bối",
            lessons: [
              {
                title: "Bài 1",
                route: "bai-6-1",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 2",
                route: "bai-6-2",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 3",
                route: "bai-6-3",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 4",
                route: "bai-6-4",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 5",
                route: "bai-6-5",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 6",
                route: "bai-6-6",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 7",
                route: "bai-6-7",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 8",
                route: "bai-6-8",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 9",
                route: "bai-6-9",
                wordCount: 100,
                favorite: 0
              },
              {
                title: "Bài 10",
                route: "bai-6-10",
                wordCount: 100,
                favorite: 0
              }
            ]
          }
        ]
      }
    },
    computed: {
      favoriteList() {
        let favorites = [0, 0, 0, 0, 0, 0]
        let totalFavorite = 0
        
        for (let i = 1; i < this.courses.length; i++) {
          for (let j = 0; j < this.courses[i].lessons.length; j++) {
            favorites[i-1] += this.courses[i].lessons[j].favorite
            totalFavorite += this.courses[i].lessons[j].favorite
          }
        }
        return (
          {
            title: "Yêu thích",
            icon: "/images/courses/favorite.svg",
            description: "Xem lại các từ vựng mà bạn yêu thích",
            lessons: [
              {
                title: "Tất cả",
                route: "yeu-thich",
                favorite: totalFavorite
              },
              {
                title: "1000 từ vựng tiếng Nhật thiết yếu",
                route: "yeu-thich?course=1000-tu-vung",
                favorite: favorites[0]
              },
              {
                title: "2000 từ vựng tiếng Nhật thiết yếu",
                route: "yeu-thich?course=2000-tu-vung",
                favorite: favorites[1]
              },
              {
                title: "3000 từ vựng tiếng Nhật thiết yếu",
                route: "yeu-thich?course=3000-tu-vung",
                favorite: favorites[2]
              },
              {
                title: "4000 từ vựng tiếng Nhật thiết yếu",
                route: "yeu-thich?course=4000-tu-vung",
                favorite: favorites[3]
              },
              {
                title: "5000 từ vựng tiếng Nhật thiết yếu",
                route: "yeu-thich?course=5000-tu-vung",
                favorite: favorites[4]
              },
              {
                title: "6000 từ vựng tiếng Nhật thiết yếu",
                route: "yeu-thich?course=6000-tu-vung",
                favorite: favorites[5]
              }
            ]
          }
        )
      }
    },
    watch: {
      
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
