/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['**/*.{html,templ}','!./node_modules', ],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
}
