export default {
  methods: {
    calculateTimeDuration: function(time) {
      let lastTime = (new Date(time)).getTime()
      let duration = parseInt((Date.now() - lastTime) / 1000);
      if (duration >= (7 * 24 * 60 * 60)) {
        let date = new Date(lastTime);
        return `${date.getDate()} tháng ${date.getMonth() + 1}, ${date.getFullYear()}`;
      } else if (duration >= (24 * 60 * 60)) {
        return parseInt(duration / (24 * 60 * 60)) + ' ngày trước';
      } else if (duration >= (60 * 60)) {
        return parseInt(duration / (60 * 60)) + ' giờ trước';
      } else if (duration >= 60) {
        return parseInt(duration / 60) + ' phút trước';
      } else {
        return 'Vừa xong';
      }
    }
  }
}