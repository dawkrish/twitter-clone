export default function Tweets(props) {
    async function handleDelete(id){
        fetch(`http://localhost:8080/tweets/${id}`, {
            method: 'DELETE',
            headers : {
                "Authorization" : `Bearer ${localStorage.getItem('token')}`
            }
        })
            .then(res => res.json())
            .then(data => window.location.reload())
            .catch(error => alert(error));
    }
    return (
        <div id="tweet-box">
            {props.tweets.map((tweet, index) => (
                <div className="tweet" key={tweet.id}>
                    <p><strong><i>{tweet.username}</i></strong></p>
                    <p id={"twee-text"}>{tweet.text}</p>
                    <p id={"tweet-time"}>{tweet.timestamp}</p>
                    {props.isUser ? (
                        <>
                            <button onClick={()=> handleDelete(tweet.id)}>Delete</button>
                        </>
                    ) : null}

                </div>
            ))}
        </div>
    );
}
