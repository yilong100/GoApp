import React, { useEffect, useState } from 'react';
import Axios from 'axios';

// Define a functional component called PostForm
function PostForm(props) {
    // Define the URL where you will send the POST request
    // When Deployed, Change to VM's IP
    const [url, setUrl] = useState("http://34.31.23.240:8080/createUser")

    // Initialize a state variable 'data' using the useState hook
    const [data, setData] = useState({
        ID:"",
        Name:"",
        Age:"",
        DreamPlaceToLive:""
    })

    // Function to update the 'data' state when input fields change
    function handle(e) {
        // Create a copy of 'data' using the spread operator
        const newData = {...data}
        // Update the value of the property corresponding to the input field
        newData[e.target.id] = e.target.value
        //Update the 'data' state with the new data
        setData(newData)
        // Log the updated data to the console
        console.log(newData) 
    }

    // Function to handle form submission
    function submit(e) {
        e.preventDefault(); // Prevent the default form submission behaviour

        // Send a POST request to the specified URL with the 'data' object
        Axios.post(url, {
            ID: data.ID,
            Name: data.Name,
            Age: parseInt(data.Age, 10), //Parse Age as an integer
            DreamPlaceToLive: data.DreamPlaceToLive
        })
        .then(response => {
            // Log the response from the server to the console
            console.log(response)
        })

        document.querySelector(".user-data").innerHTML = `Welcome ${data.Name}!`
    }

    // Render the form component
    return(
        <div>
            <form onSubmit={(e) => submit(e)}>
            <label>ID:</label><br></br>
            <input id="ID" onChange={(e) => handle(e)} ></input><br></br>
            <label>Name:</label><br></br>
            <input id="Name" onChange={(e) => handle(e)}></input><br></br>
            <label>Age:</label><br></br>
            <input id="Age" onChange={(e) => handle(e)}></input><br></br>
            <label>Dream Place to Live:</label><br></br>
            <input id="DreamPlaceToLive" onChange={(e) => handle(e)}></input><br></br>
            <p></p>
            <button type="submit">Submit</button>
            <p class="user-data"></p>
            </form>
        </div>
    )

}

export default PostForm;