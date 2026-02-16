import { useState, useEffect } from 'react'; // 1. Add useEffect import
import reactLogo from '../assets/react.svg';
import viteLogo from '/vite.svg';
import '../App.css';
import { Button } from "../components/ui/button"
import { Link, useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/authstore"



function App() {
    const [goMessage, setGoMessage] = useState("Loading message from Go backend..."); // 2. New state for backend message
    const logout = useAuthStore((s) => s.logout)
    const user = useAuthStore((s) => s.user)
    const navigate = useNavigate()

    console.log("API URL:", import.meta.env.VITE_URL);

    function handleLogout() {
        logout()
        navigate("/")
    }
    // 3. Add data fetching logic with useEffect
    useEffect(() => {
        fetch(`${import.meta.env.VITE_URL}/api/hello`)
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


    // ----------------------------------------------------
    return (
        <div className="flex items-center justify-center h-screen">

            {/* Optional: A sub-container to group content together if needed */}
            <div className="text-center">
                <div>
                    <h1 className="text-4xl font-bold text-cyan-400 mb-4">
                        Spencer's Map App
                    </h1>
                    {/* Wrap logos in a flex container to put them side-by-side */}
                    <div className="flex justify-center space-x-6 mb-4">
                        <a href="https://vite.dev" target="_blank" rel="noopener noreferrer">
                            <img src={viteLogo} className="logo w-24" alt="Vite logo" />
                        </a>
                        <a href="https://react.dev" target="_blank" rel="noopener noreferrer">
                            <img src={reactLogo} className="logo react w-24" alt="React logo" />
                        </a>
                    </div>
                </div>

                <h1 className="text-2xl font-semibold">Vite + React</h1>
                <h2 className="text-lg text-gray-600 mb-4">Spencer's Learning</h2>
                <p className=" text-black-600 mb-6">
                    Hello, {user?.username ? (user?.username) : (<>Guest</>)} welcome to my Map App.
                    <br /> This is my first solo project that I've developed in order to improve my knowledge,
                    skills and employability.
                    <br /> As this is a learning project I cannot gurantee full data protection so please do not use personal information.
                    <br />Should you find any issues from a security or design perspective I would love the feedback on how I can further improve this application

                </p>
                {/* 4. Display the message from the Go backend */}
                <div className="card mb-4 p-4 border rounded shadow-sm bg-white">
                    <p className="font-bold"> Due to this being a solo project with a limited budget the backend and database will not be active at
                        all times of the day
                    </p>
                    <p>Go Backend Status: **{goMessage}**</p>
                    {/*<p>Hello, {user?.username ? (user?.username) : (<>Guest</>)}</p>*/}
                </div>

                {/* 5. Keep the existing counter functionality */}
                <div className="card mb-4 p-4 border rounded shadow-sm bg-white">

                    <p className="mt-2 text-sm text-gray-500">
                    </p>
                    <Button> <Link to="/Map">Map</Link> </Button>
                </div>
                <p className="read-the-docs text-sm text-gray-500">
                    Click on the Vite and React logos to learn more about them
                </p>

                {user?.email ? (
                    <Button onClick={handleLogout}>
                        Logout
                    </Button>) : (

                    <ul>
                        {
                            <li>
                                <Button> <Link to="/Login">Login</Link> </Button>
                                <Button> <Link to="/Signup">Sign up</Link> </Button>
                            </li>

                        }

                    </ul>)}

            </div>
        </div >
    );
}


export default App;