import React, {useState} from 'react';

export default function Login() {
    const [name, setName] = useState('');
    const [password, setPassword] = useState('');

    async function handleSubmit(e) {
        e.preventDefault();
        const formData = {
            name: name,
            password: password
        };

        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "post",
                body: JSON.stringify(formData)
                // You can include headers, body, etc. here
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }

            const data = await response.json();
            localStorage.setItem("token", data.token);
            localStorage.setItem('username',name)
            window.location.href = "/"
            console.dir(data);
        } catch (error) {
            alert(error.message)
            console.log(error);
        }
    }


    return (
        <div className={"forms"}>
            <form>
                <div id="form-grid">
                    <div className="form-col">
                        <label htmlFor="name">Name</label>
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
            <p>Don't have an account ? <a href="/signup">Sign Up</a></p>
        </div>
    );
}
