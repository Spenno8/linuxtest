import { useState, useEffect } from 'react'; // 1. Add useEffect import
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';

function App() {
  const [count, setCount] = useState(0); // Existing counter state
  const [goMessage, setGoMessage] = useState("Loading message from Go backend..."); // 2. New state for backend message

  // 3. Add data fetching logic with useEffect
  useEffect(() => {
    fetch('http://localhost:8080/api/hello')
      .then(async (response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setGoMessage(data.message);
      })
      .catch((error) => {
        console.error("Fetch error:", error);
        setGoMessage("Failed to connect to the Go backend.");
      });
  }, []);

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <h2>Spencer's Learning</h2>

      {/* 4. Display the message from the Go backend */}
      <div className="card">
        <p>Go Backend Status: **{goMessage}**</p>
      </div>

      {/* 5. Keep the existing counter functionality */}
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App;
