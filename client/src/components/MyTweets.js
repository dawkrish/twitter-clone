import React, {useEffect, useState} from 'react';
import Tweets from './Tweets';

export default function MyTweets() {
    const [tweetText, setTweetText] = useState('');
    const [tweets, setTweets] = useState([]);

    useEffect(() => {
        fetchAllTweets()
    }, []);
    async function handleCreateTweet() {
        try {
            const response = await fetch("http://localhost:8080/tweets", {
                method: "post",
                headers: {
                    "Authorization": `Bearer ${localStorage.getItem("token")}`,
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    text: tweetText // Replace 'text' with the actual property name in your API
                })
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }
            let newTweetRecieved = await response.json()

            setTweets([...tweets,newTweetRecieved])
            window.location.reload()
            // Handle the successful creation of the tweet as needed
            console.log('Tweet created successfully');
            // You might want to refetch the tweets here or update the UI in some way
        } catch (error) {
            console.log(error);
        }
    }
    async function fetchAllTweets() {
        try {
            const response = await fetch("http://localhost:8080/mytweets", {
                method: "get",
                headers : {
                    "Authorization" : `Bearer ${localStorage.getItem('token')}`
                }
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }

            const data = await response.json();
            console.log(data)
            setTweets(data);
        } catch (error) {
            console.log(error);
        }
    }

    function handleLogout(){
        localStorage.setItem('token','')
        localStorage.setItem('username','')
    }
    return (
        <div id="home-page">
            <div className={"home-child"} id={"home-1"}>
                <textarea name="" id="" cols="40" rows="5"
                          value={tweetText}
                          onChange={(e) => setTweetText(e.target.value)}
                          placeholder="Enter your tweet"
                ></textarea>
                <br/>
                <button id="create-tweet-btn" onClick={handleCreateTweet}>Create Tweet</button>
            </div>
            <div className={"home-child"} id='home-2'>
                <h1>My Tweets</h1>
                <Tweets tweets={tweets} isUser={true} />
            </div>
            <div className={"home-child"} id={"home-3"}>
                <h3>User : <em>{localStorage.getItem('username')}</em></h3>
                <a href="/" onClick={handleLogout} >Logout</a>
                <a href="/">Home</a>
            </div>
        </div>
    );
}
