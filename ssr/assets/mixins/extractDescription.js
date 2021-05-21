export default {
  methods: {
    extractDescription: function(content, maxLength) {
      if (process.client) {
        let span = document.createElement('span');
        span.innerHTML = content;
        let text = span.textContent || span.innerText;
        if (text.length <= maxLength)
          return text
        return text.slice(0, maxLength) + '...'
      }

      return ''
    }
  }
}