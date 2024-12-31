const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: "./src/index.tsx",
  output: {
    publicPath: "/",
    path: path.resolve(__dirname, "dist"),
    filename: "bundle.js",
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        exclude: /node_modules/,
        use: "ts-loader",
      },
      {
        test: /\.scss$/,
        use: ["style-loader", "css-loader", "sass-loader"],
      },
      {
        test: /\.svg$/,
        use: [
          {
            loader: "@svgr/webpack",
            options: {
              svgo: false,
            },
          },
          "file-loader",
        ],
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js", ".jsx", ".json", ".scss"],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, "public", "index.html"),
      filename: "index.html",
    }),
    new webpack.DefinePlugin({
      "process.env.IGI_API_DEV": JSON.stringify("http://localhost:8080"),
    }),
  ],
  devServer: {
    static: path.join(__dirname, "dist"),
    compress: true,
    port: 3000,
    historyApiFallback: true,
  },
  performance: {
    hints: false,
  },
};
