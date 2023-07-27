/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",                     //出現下面這個錯誤，在tsx後面加上vue 
    "./src/**/*.{js,ts,jsx,tsx,vue}",   //warn - No utility classes were detected in your source files. If this is unexpected, double-check the `content` option in your Tailwind CSS configuration.
    'node_modules/preline/dist/*.js',   // warn - https://tailwindcss.com/docs/content-configuration
  ],                                    //https://blog.csdn.net/ccyu86/article/details/130125700
  theme: {
    extend: {},
  },
  plugins: [ 
    require('preline/plugin'),
  ],
  darkMode: 'class',
}
