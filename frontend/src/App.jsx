import { useState, useEffect } from "react";
import { BrowserRouter as Router, Route, Routes, Navigate, BrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import LoginForm from "./components/LoginForm";
import SignupForm from "./components/SignUpForm";
import Navbar from "./components/Navbar";

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem("token"));

  useEffect(() => {
    const checkAuth = () => {
      setIsLoggedIn(!!localStorage.getItem("token"));
    };

    window.addEventListener("storage", checkAuth);
    return () => {
      window.removeEventListener("storage", checkAuth);
    };
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setIsLoggedIn(false);
    window.location.href = "/auth/login"; // Redirect user on logout
  };

  return (
    <>
    <Navbar isLoggedIn={isLoggedIn} onLogout={handleLogout} />
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home isLoggedIn={isLoggedIn} />} />
        <Route path="/auth/login" element={isLoggedIn ? <Navigate to="/" /> : <LoginForm setIsLoggedIn={setIsLoggedIn} />} />
        <Route path="/auth/signup" element={isLoggedIn ? <Navigate to="/" /> : <SignupForm />} />
      </Routes>
    </BrowserRouter>
    </>
  );
};

export default App;
