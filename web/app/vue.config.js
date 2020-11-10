module.exports = {
    devServer: {
        proxy: {
          '^/backend/*': {
            target: 'http://localhost:4444',
            changeOrigin: true,
            secure: false,
            pathRewrite: {'^/backend' : ''}
          }
        }
      }
    
}