export const state = () => ({
  id: null,
  email: '',
  name: '',
  avatar: '',
  birthday: 0,
  gender: 0,
  address: '',
  city: '',
  about: '',
  facebookId: '',
  likedArticles: [],
  likedVocas: [],
  likedGrammars: [],
  likedKanji: []
})

export const mutations = {
  setUser (state, user) {
    state.id = user.id
    state.email = user.email
    state.name = user.name
    state.avatar = user.avatar
    state.birthday = user.birthday
    state.gender = user.gender
    state.address = user.address
    state.city = user.city
    state.about = user.about
    state.facebookId = user.facebookId
    state.likedArticles = user.likedArticles
    state.likedVocas = user.likedVocas
    state.likedGrammars = user.likedGrammars
    state.likedKanji = user.likedKanji
  },
  resetUser(state) {
    state.id = null
    state.email = ''
    state.name = ''
    state.avatar = ''
    state.birthday = 0
    state.gender = 0
    state.address = ''
    state.city = ''
    state.about = ''
    state.facebookId = ''
    state.likedArticles = []
    state.likedVocas = []
    state.likedGrammars = []
    state.likedKanji = []
  },
  setId (state, id) {
    state.id  = id
  },
  setEmail (state, email) {
    state.email  = email
  },
  setName (state, name) {
    state.name  = name
  },
  setAvatar (state, name) {
    state.name  = name
  },
  setBirthday (state, name) {
    state.name  = name
  },
  setGender (state, name) {
    state.name  = name
  },
  setAddress (state, name) {
    state.name  = name
  },
  setCity (state, name) {
    state.name  = name
  },
  setAbout (state, about) {
    state.about  = about
  },
  setFacebookId (state, facebookId) {
    state.facebookId = facebookId
  },
  setLikedArticles (state, likedArticles) {
    state.likedArticles = likedArticles
  },
  setLikedVocas (state, likedVocas) {
    state.likedVocas = likedVocas
  },
  setLikedGrammars (state, likedGrammars) {
    state.likedGrammars = likedGrammars
  },
  setLikedKanji (state, likedKanji) {
    state.likedKanji = likedKanji
  }
}