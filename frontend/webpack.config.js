var path = require('path');
module.exports = {
  entry:  {
    main: './src/main.js'
  },
  output: {
    path: path.join(__dirname, '../www/js'),
    filename: '[name].bundle.js'
  },
  module: {
    loaders: [
      {
        test: path.join(__dirname, 'src'),
        loader: 'babel-loader' ,
        query: {
            presets: ['es2015']
        }
      }
    ]
  }
};
