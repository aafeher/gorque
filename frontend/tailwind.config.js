/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        dark: {
          primary: '#1E293B',
          secondary: '#334155',
          accent: '#4F46E5',
        },
      },
    },
  },
  plugins: [],
};
