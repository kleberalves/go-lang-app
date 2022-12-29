const nextRuntimeDotenv = require('next-runtime-dotenv')
const path = require('path');

require('dotenv').config({
  path: path.resolve(
    __dirname,
    `.env.${process.env.NODE_ENV}`,
  ),
});

const withConfig = nextRuntimeDotenv({
  public: [
    'BASE_URL'
  ],
  server: [
    'BASE_URL_SERVER'
  ]
})

module.exports = withConfig()