import pkg from './package'

export default {
  mode: 'universal',

  server: {
    host: '0',
    port: 8000
  },

  env: {
    apiUrl: process.env.API_URL || 'http://localhost:3000',
    tinyMceKey: 'vpk5tyazspp1ek1y2pn6h33vds99yxa8mj6rmqcgn7qiwdn7',
    defaultArticleAvatar: 'https://i.stack.imgur.com/y9DpT.jpg',
    defaultDocumentAvatar: 'https://i.stack.imgur.com/y9DpT.jpg',
    adminUserId: 1
  },

  /*
  ** Headers of the page
  */
  head: {
    title: pkg.name,
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: pkg.description }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: '//fonts.googleapis.com/css?family=Roboto:400,500,700,400italic|Material+Icons' }
    ]
  },

  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff', throttle: 0 },

  /*
  ** Global CSS
  */
  css: [
  ],

  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    '~/plugins/index',
    '~/plugins/error'
  ],

  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios'
  ],
  /*
  ** Axios module configuration
  */
  axios: {
    // See https://github.com/nuxt-community/axios-module#options
  },

  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
    }
  }
}
