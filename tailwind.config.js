module.exports = {
  content: [
    "./app/views/**/*.{html,js,go}",  // covers .html, .js, .go templates
    "./public/js/**/*.js",            // includes client-side JS in public
    "./lib/**/*.go",                  // if you render templates via Go files
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
