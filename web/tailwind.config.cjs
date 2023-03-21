/** @type {import('tailwindcss').Config} */
module.exports = {
  corePlugins: {
    preflight: false
  },
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: 'var(--pf-global--primary-color--100)'
      }
    }
  },
  important: true,
  plugins: []
};
