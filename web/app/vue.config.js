module.exports = {
  configureWebpack: {
    devtool: 'source-map'
  },
  devServer: {
    proxy: {
      '^/backend/*': {
        target: 'http://localhost:4444',
        changeOrigin: true,
        secure: false,
        logLevel: "debug",
        pathRewrite: {
          '^/backend': ''
        }
      }
    }
  }

}