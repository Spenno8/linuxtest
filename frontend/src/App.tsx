//import { useState, useEffect } from 'react'; // 1. Add useEffect import
//import reactLogo from './assets/react.svg';
//import viteLogo from '/vite.svg';
import './App.css';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import Home from "./pages/Home"
import Login from "./pages/Login"
import Signup from "./pages/Signup"
import Map from "./pages/Map"

//import { Button } from "@/components/ui/button"


function App() {
  return (
    // https://www.geeksforgeeks.org/reactjs/how-to-redirect-to-another-page-in-reactjs/
    <>
      <Router>
        <Routes>
          <Route
            //exact
            path="/"
            element={<Home />}
          />
          <Route
            path="/Login"
            element={<Login />}
          />
          <Route
            path="/Signup"
            element={<Signup />}
          />
          <Route
            path="/Map"
            element={<Map />}
          />
          <Route
            path="*"
            element={<Navigate to="/" />}
          />
        </Routes>
      </Router>
    </>
  );
}

export default App;

