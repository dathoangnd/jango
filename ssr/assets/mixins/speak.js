export default {
  methods: {
    speak(text) {
      let utterance = new SpeechSynthesisUtterance()
      utterance.lang = 'ja-JP'
      utterance.rate = 0.75
      utterance.text = text
      speechSynthesis.cancel()
      speechSynthesis.speak(utterance)
    }
  }
}