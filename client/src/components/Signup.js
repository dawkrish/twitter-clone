import React, {useState} from "react";

export default function Signup(){
    const [name, setName] = useState('');
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('');

    async function handleSubmit(e) {
        e.preventDefault();
        const formData = {
            name: name,
            email:email,
            password: password
        };

        try {
            const response = await fetch("http://localhost:8080/signup", {
                method: "post",
                body: JSON.stringify(formData)
                // You can include headers, body, etc. here
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }

            const data = await response.json();
            console.dir(data)
            window.location.href = "/login"

            console.dir(data);
        } catch (error) {
            alert(error.message)
            console.dir(error);
        }
    }


    return (
        <div className={"forms"}>
            <form>
                <div id="form-grid">
                    <div className="form-col">
                        <label htmlFor="name">Name</label>
                        <label htmlFor="email">Email</label>
                        <label htmlFor="password">Password</label>
                    </div>
                    <div className="form-col">
                        <input
                            type="text"
                            name="name"
                            id="name"
                            value={name}
                            required={true}
                            onChange={(e) => setName(e.target.value)}
                        />
                        <input
                            type="email"
                            name="email"
                            id="email"
                            value={email}
                            required={true}
                            onChange={(e) => setEmail(e.target.value)}
                        />
                        <input
                            type="password"
                            name="password"
                            value={password}
                            required={true}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                </div>



                <button type="submit" onClick={handleSubmit}>Submit</button>

            </form>
            <p>Already have an account ? <a href="/login">Log In</a></p>
        </div>
    );
}