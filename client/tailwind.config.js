/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        backgroundColor: "rgba(var(--backgroundColor))",
        myMessage: "rgba(var(--myMessage))",
        theirMessage: "rgba(var(--theirMessage))",
        fontColor: "rgba(var(--fontColor))",
        chatBackground: "rgba(var(--chatBackground))",
      }
    },
  },
  plugins: [],
}

