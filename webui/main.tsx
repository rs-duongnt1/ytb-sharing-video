import React from "react";
import ReactDOM from "react-dom/client";
import { Provider } from "react-redux";
import { store } from "./store";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./index.css";
import "./reset.css";
import AppContainer from "./containers/AppContainer.js";
import { WebSocketProvider } from "./providers/WebSocketProvider";
console.log("loaded");
ReactDOM.createRoot(document.getElementById("root")!).render(
  <WebSocketProvider>
    <React.StrictMode>
      <Provider store={store}>
        <AppContainer />
      </Provider>
    </React.StrictMode>
  </WebSocketProvider>
);
