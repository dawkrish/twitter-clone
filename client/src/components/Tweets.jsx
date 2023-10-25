export default function Tweets(props) {
  async function handleDelete(id) {
    fetch(`http://localhost:8080/tweets/${id}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    })
      .then((res) => res.json())
      .then((data) => window.location.reload())
      .catch((error) => alert(error));
  }
  async function handleUpdate(tweet) {
    try{
        const response = await fetch(`http://localhost:8080/tweets/${tweet.id}`,{
                method : "PUT",
                headers: {
                    "Authorization": `Bearer ${localStorage.getItem("token")}`,
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    text: document.getElementById('update_tweet_textarea').value // Replace 'text' with the actual property name in your API
                })
            }
        )
        if(!response.ok){
            throw new Error(await response.text());
        }
        window.location.reload()
    }catch(e){
        console.log(e);
    }
  }

  function showUpdateBox(tweet) {
    let overlay = document.createElement("div");
    overlay.innerHTML = `<div>
      <h4>Current tweet : ${tweet.text} </h4>
      <br />
      <textarea id="update_tweet_textarea" name="" id="" cols="40" rows="5"
                          placeholder="Enter updated tweet"
                ></textarea>
      <button id="update-btn">submit & update </button>
    </div>`;
    overlay.classList.add("overlay");
    document.getElementById("home-2").appendChild(overlay);
    document.getElementById("tweet-box").style.opacity = "0.2";
    document.getElementById('update-btn').addEventListener('click',()=>{handleUpdate(tweet)})
  }
  return (
    <div id="tweet-box">
      {props.tweets.map((tweet, index) => (
        <div className="tweet" key={tweet.id}>
          <p>
            <strong>
              <i>{tweet.username}</i>
            </strong>
          </p>
          <p id={"twee-text"}>{tweet.text}</p>
          <p id={"tweet-time"}>{tweet.timestamp}</p>
          {props.isUser ? (
            <>
              <button onClick={() => handleDelete(tweet.id)}>Delete</button>
              <button onClick={() => showUpdateBox(tweet)}>Update</button>
            </>
          ) : null}
        </div>
      ))}
    </div>
  );
}
