import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import About from "./pages/About";

const App = () => {
  return (
    <Routes>
      <Route index element={<Home />} />
      <Route path="/about" element={<About />} />
    </Routes>
  );
};

const root = ReactDOM.createRoot(document.querySelector("#App")!);
root.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>
);
