module.exports = {
  devServer: {
    proxy: {
      '^/api': {
        target: 'http://localhost:1323',
        ws: true,
        changeOrigin: true,
      },
    },
  },
};
