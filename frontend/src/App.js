import React from "react";
import "./css/style.css";
import MainPage from "./pages/MainPage";
import ArticlePage from "./pages/Articles";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  Link
} from "react-router-dom";

function App() {
  return (
    <Router>
      <NavigateBar />
        <Routes>
          <Route path="/" element={<MainPage />} />
          <Route path="/articles" element={<ArticlePage />} />
        </Routes>
      <Footer />
    </Router>
  );
}

function NavigateBar() {
  return (
    <div className="main-page">
      <div className="div">
        <div className="group">
            <div className="navbar">
              <div className="blog">
                <Link to="/">Home</Link>
              </div>
              <div className="about">
                <Link to="/articles">Blogs</Link>
              </div>
              <a className="text-wrapper" href="https://github.com" rel="noopener noreferrer" target="_blank">
                Projects
              </a>
              <div className="text-wrapper-2">
                <Link to="/resume">Resume</Link>
              </div>
              <div className="text-wrapper-3">
                <Link to="/links">Links</Link>
              </div>
            </div>
        </div>
      </div>
    </div>
  );
}

function Footer() {
  return(
    <div className="main-page">
      <div className="div">
        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-4">© 2024 Safa Bayar</div>
              <div className="group-2">
                <div className="text-wrapper-5">LinkedIn</div>
                <div className="text-wrapper-6">LinkedIn</div>
                <div className="text-wrapper-7">LinkedIn</div>
                <div className="text-wrapper-8">LinkedIn</div>
                <div className="text-wrapper-9">LinkedIn</div>
              </div>
          </div>
        </div>
  </div>
</div>
  );
}

export default App;



