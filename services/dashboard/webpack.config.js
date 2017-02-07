const webpack = require('webpack');

const extractCommons = new webpack.optimize.CommonsChunkPlugin({
  name: 'commons',
  filename: 'commons.js'
});

const ExtractTextPlugin = require('extract-text-webpack-plugin');
const extractCSS = new ExtractTextPlugin('[name].bundle.css');

const config = {
  context: __dirname + '/app',
  entry: {
    app: './app.js',
    commons: ['angular']
  },
  watch: true,
  output: {
    path: __dirname + '/dist',
    publicPath: '/dist/',
    filename: '[name].bundle.js'
  },
  devServer: {
    inline: true,
    port: 8000
  },
  module: {
    rules: [{
      test: /\.(png|jpg)$/,
      use: [{
        loader: 'url-loader',
        options: { limit: 10000 } // Convert images < 10k to base64 strings
      }]
    }, {
      test: /\.scss$/,
      loader: ['style-loader', 'css-loader','sass-loader']
    }, {
      test: /\.js$/,
      include: __dirname + '/app',
      use: [{
        loader: 'babel-loader',
        options: {
          presets: [
            ['es2015', { modules: false }]
          ]
        }
      }]
    },
    {
      test: /\.html$/,
      loader: 'raw-loader'
    }]
  },
  plugins: [
    new webpack.NamedModulesPlugin(),
    extractCSS,
    extractCommons
  ]
};

module.exports = config;
