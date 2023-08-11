import { defineConfig } from "cypress";

export default defineConfig({
  component: {
    devServer: {
      framework: "react",
      bundler: "vite",
    },
    fixturesFolder: "src",
  },

  env: {
    "cypress-react-unit-test": {
      react: "node_modules/react/umd/react.development.js",
      "react-dom": "node_modules/react-dom/umd/react-dom.development.js",
    },
  },

  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
