var webpack = require('webpack');
var ExtractTextPlugin = require('extract-text-webpack-plugin');
var production = minimize = process.argv.indexOf('--production') !== -1;

const entry = './js/app.js';

/*
 * If you want to minify assets provide the --production flag
 * to the webpack binary
 * if you don't provide this everything remains normal
 */

if(!production){
        var plugins = [new ExtractTextPlugin("css/out/build.css")]
}
else{
        var plugins = [
                new ExtractTextPlugin("css/out/build.css"),
                new webpack.optimize.UglifyJsPlugin({
                        compress: { warnings: false }
                })
        ]
}

module.exports = {
        entry: entry,
        
        resolve: {
                alias:{
                        vue: 'vue/dist/vue.js'
                }
        },
        
        output: {
                path: __dirname,
                filename: 'js/out/build.js'
        },
        
        /* Loaders
         * .vue -> view-loader
         * .js -> babel-loader
         */
        module: {
              loaders: [
                {
                        test: /\.vue$/,
                        loader: 'vue'
                },
                {
                        test: /\.js$/,
                        loader: 'babel',
                        exclude: /node_modules/
                }
              ]
        },
        
        /* Vue loader settings */
        vue: {
                autoprefixer: {
                        browsers: ['last 2 versions', 'IE >= 10', 'IOS 7']
                },
                /* Save inline css to file */
                loaders:{
                        sass: ExtractTextPlugin.extract("css!sass")
                }
        },
        
        /* Load babel with presets */
        babel: {
                presets: ['es2015'],
                plugins: ['transform-runtime']
        },
        
        plugins: plugins,
        
        /* Default to watching with some extras */
        watch: true,
        progress: true,
        colors: true
}