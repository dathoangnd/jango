export const state = () => ({
  loginStatus: false,
  showLoginForm: false,
  showSnackbar: false,
  snackbarMessage: ''
})

export const mutations = {
  toggleLoginForm (state, show) {
    state.showLoginForm  = show
  },
  toggleSnackbar (state, show) {
    state.showSnackbar = show
  },
  setSnackbarMessage(state, message) {
    state.snackbarMessage = message
  },
  setLoginStatus(state, status) {
    state.loginStatus = status
  }
}