export default {
  methods: {
    saveJwt: function(token) {
      localStorage.jwt_token = token
    },

    removeJwt: function() {
      localStorage.removeItem('jwt_token');
    },

    getJwt: function() {
      if (localStorage.jwt_token != undefined) {
        let token = localStorage.jwt_token
        let tokenParts = token.split('.')
        if (tokenParts.length == 3) {
          let payload = JSON.parse(decodeURIComponent(escape(atob(tokenParts[1]))))
          let timeExpr = payload.exp * 1000
          if (timeExpr > Date.now()) {
            return token
          }
        }
      }

      return null
    },

    getPayload: function() {
      if (localStorage.jwt_token != undefined) {
        let token = localStorage.jwt_token
        let tokenParts = token.split('.')
        if (tokenParts.length == 3) {
          let payload = JSON.parse(decodeURIComponent(escape(atob(tokenParts[1]))))
          let timeExpr = payload.exp * 1000
          if (timeExpr > Date.now()) {
            return payload
          }
        }
      }

      return null
    },

    saveLikedArticles: function(likedArticles) {
      localStorage.liked_articles = likedArticles.trim()
    },

    removeLikedArticles: function() {
      localStorage.removeItem('liked_articles');
    },

    getLikedArticles: function() {
      if (localStorage.liked_articles != undefined && localStorage.liked_articles.trim() != '') {
        let likedArticles = localStorage.liked_articles.split(' ')
        likedArticles.forEach(function(id, index) {
          likedArticles[index] = Number(id)
        })

        return likedArticles
      }

      return []
    },

    addLikedArticle: function(articleId) {
      articleId = parseInt(articleId)
      let likedArticles = []
      if (localStorage.liked_articles != undefined && localStorage.liked_articles != '') {
        likedArticles = localStorage.liked_articles.split(' ')
        likedArticles.forEach(function(id, index) {
          likedArticles[index] = Number(id)
        })
      }

      if (likedArticles.indexOf(articleId) == -1) {
        likedArticles.push(articleId)
      }

      localStorage.liked_articles = likedArticles.join(' ')
    },

    removeLikedArticle: function(articleId) {
      articleId = parseInt(articleId)
      let likedArticles = []
      if (localStorage.liked_articles != undefined) {
        likedArticles = localStorage.liked_articles.split(' ')
        likedArticles.forEach(function(id, index) {
          likedArticles[index] = Number(id)
        })
      }
    
      if (likedArticles.indexOf(articleId) != -1) {
        likedArticles.splice(likedArticles.indexOf(articleId), 1)
      }

      localStorage.liked_articles = likedArticles.join(' ')
    },

    saveLikedVocas: function(likedVocas) {
      localStorage.liked_vocas = likedVocas.trim()
    },
    
    removeLikedVocas: function() {
      localStorage.removeItem('liked_vocas');
    },

    getLikedVocas: function() {
      if (localStorage.liked_vocas != undefined && localStorage.liked_vocas.trim() != '') {
        let likedVocas = localStorage.liked_vocas.split(' ')
        likedVocas.forEach(function(order, index) {
          likedVocas[index] = Number(order)
        })

        return likedVocas
      }

      return []
    },

    addLikedVoca: function(vocaOrder) {
      let likedVocas = []
      if (localStorage.liked_vocas != undefined && localStorage.liked_vocas != '') {
        likedVocas = localStorage.liked_vocas.split(' ')
        likedVocas.forEach(function(order, index) {
          likedVocas[index] = Number(order)
        })
      }

      if (likedVocas.indexOf(vocaOrder) == -1) {
        likedVocas.push(vocaOrder)
      }

      localStorage.liked_vocas = likedVocas.join(' ')
    },

    removeLikedVoca: function(vocaOrder) {
      let likedVocas = []
      if (localStorage.liked_vocas != undefined && localStorage.liked_vocas != '') {
        likedVocas = localStorage.liked_vocas.split(' ')
        likedVocas.forEach(function(order, index) {
          likedVocas[index] = Number(order)
        })
      }
    
      if (likedVocas.indexOf(vocaOrder) != -1) {
        likedVocas.splice(likedVocas.indexOf(vocaOrder), 1)
      }

      localStorage.liked_vocas = likedVocas.join(' ')
    },

    saveLikedGrammars: function(likedGrammars) {
      localStorage.liked_grammars = likedGrammars.trim()
    },
    
    removeLikedGrammars: function() {
      localStorage.removeItem('liked_grammars');
    },

    getLikedGrammars: function() {
      if (localStorage.liked_grammars != undefined && localStorage.liked_grammars.trim() != '') {
        let likedGrammars = localStorage.liked_grammars.split(' ')

        return likedGrammars
      }

      return []
    },

    addLikedGrammar: function(grammarOrder) {
      let likedGrammars = []
      if (localStorage.liked_grammars != undefined && localStorage.liked_grammars != '') {
        likedGrammars = localStorage.liked_grammars.split(' ')
      }

      if (likedGrammars.indexOf(grammarOrder) == -1) {
        likedGrammars.push(grammarOrder)
      }

      localStorage.liked_grammars = likedGrammars.join(' ')
    },

    removeLikedGrammar: function(grammarOrder) {
      let likedGrammars = []
      if (localStorage.liked_grammars != undefined && localStorage.liked_grammars != '') {
        likedGrammars = localStorage.liked_grammars.split(' ')
      }
    
      if (likedGrammars.indexOf(grammarOrder) != -1) {
        likedGrammars.splice(likedGrammars.indexOf(grammarOrder), 1)
      }

      localStorage.liked_grammars = likedGrammars.join(' ')
    },

    ///
    saveLikedKanjiList: function(likedKanji) {
      localStorage.liked_kanji = likedKanji.trim()
    },
    
    removeLikedKanjiList: function() {
      localStorage.removeItem('liked_kanji');
    },

    getLikedKanjiList: function() {
      if (localStorage.liked_kanji != undefined && localStorage.liked_kanji.trim() != '') {
        let likedKanji = localStorage.liked_kanji.split(' ')

        return likedKanji
      }

      return []
    },

    addLikedKanji: function(kanjiOrder) {
      let likedKanji = []
      if (localStorage.liked_kanji != undefined && localStorage.liked_kanji != '') {
        likedKanji = localStorage.liked_kanji.split(' ')
      }

      if (likedKanji.indexOf(kanjiOrder) == -1) {
        likedKanji.push(kanjiOrder)
      }

      localStorage.liked_kanji = likedKanji.join(' ')
    },

    removeLikedKanji: function(kanjiOrder) {
      let likedKanji = []
      if (localStorage.liked_kanji != undefined && localStorage.liked_kanji != '') {
        likedKanji = localStorage.liked_kanji.split(' ')
      }
    
      if (likedKanji.indexOf(kanjiOrder) != -1) {
        likedKanji.splice(likedKanji.indexOf(kanjiOrder), 1)
      }

      localStorage.liked_kanji = likedKanji.join(' ')
    },

    saveLikedThreads: function(likedThreads) {
      localStorage.liked_threads = likedThreads.trim()
    },

    removeLikedThreads: function() {
      localStorage.removeItem('liked_threads');
    },

    getLikedThreads: function() {
      if (localStorage.liked_threads != undefined && localStorage.liked_threads.trim() != '') {
        let likedThreads = localStorage.liked_threads.split(' ')
        return likedThreads
      } else {
        localStorage.liked_threads = ''
      }

      return []
    },

    addLikedThread: function(id) {
      let likedThreads = []
      if (localStorage.liked_threads != undefined && localStorage.liked_threads != '') {
        likedThreads = localStorage.liked_threads.split(' ')
      }

      if (likedThreads.indexOf(id) == -1) {
        likedThreads.push(id)
      }

      localStorage.liked_threads = likedThreads.join(' ')
    },

    saveLikedReplies: function(likedReplies) {
      localStorage.liked_replies = likedReplies.trim()
    },

    removeLikedReplies: function() {
      localStorage.removeItem('liked_replies');
    },

    getLikedReplies: function() {
      if (localStorage.liked_replies != undefined && localStorage.liked_replies.trim() != '') {
        let likedReplies = localStorage.liked_replies.split(' ')
        return likedReplies
      } else {
        localStorage.liked_replies = ''
      }

      return []
    },

    addLikedReply: function(id) {
      let likedReplies = []
      if (localStorage.liked_replies != undefined && localStorage.liked_replies != '') {
        likedReplies = localStorage.liked_replies.split(' ')
      }

      if (likedReplies.indexOf(id) == -1) {
        likedReplies.push(id)
      }

      localStorage.liked_replies = likedReplies.join(' ')
    }
  }
}