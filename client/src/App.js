import './App.css';
import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Routes,Route } from 'react-router-dom'; // Import Router and Route
import Home from "./components/Home";
import StartMenu from "./components/StartMenu";
import Login from "./components/Login"; // Import your Login component
import Signup from "./components/Signup"; // Import your Signup component
import MyTweets from "./components/MyTweets";
function App() {
    const [loggedIn, setLoggedIn] = useState(false);

    useEffect(() => {
        checkLogin();
    }, []);

    async function checkLogin() {
        try {
            const response = await fetch("http://localhost:8080/loginstate", {
                headers: {
                    "Authorization": `Bearer ${localStorage.getItem('token')}`
                }
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }

            setLoggedIn(true);
            console.dir(await response.json());
        } catch (e) {
            localStorage.removeItem('token')
            console.log(e);
        }
    }

    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route path="/login" element={<Login/>} />
                    <Route path="/signup" element={<Signup/>}/>
                    <Route exact path="/" element={loggedIn ? <Home /> : <StartMenu />}/>
                    <Route path="/mytweets" element={<MyTweets/>} />
                </Routes>

            </div>
        </Router>
    );
}

export default App;
