import React, { useEffect, useState } from 'react';
import Axios from 'axios';
import { Button, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField } from '@mui/material';
import apiUrl from "../backend-ip-address.js"

// Define a functional component called PostForm
function PostForm(props) {

    // Define the URL where you will send the POST request
    // When Deployed, Change to VM's IP
    const [url, setUrl] = useState("http://"+ apiUrl + ":8080")
    const [update, setUpdate] = useState(false)
    const [users, setUsers] = useState([])

    useEffect(()=>{
        // Send a Get request to receive all users
        getUsers();
    }, [update])

    useEffect(()=>{
        // Send a Get request to receive all users
        getUsers();
    }, [])

    const getUsers = () => {
        // Send a Get request to receive all users
        Axios.get(url + "/users")
        .then(response => {
            // Log the response from the server to the console
            console.log(response.data)
            setUsers(response.data)
        }).catch(e => {
            console.log(e)
        })
    }

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
    }

    // Function to handle form submission
    function submit(e) {
        e.preventDefault(); // Prevent the default form submission behaviour

        // Send a POST request to the specified URL with the 'data' object
        Axios.post(url + "/createUser", {
            ID: data.ID,
            Name: data.Name,
            Age: parseInt(data.Age, 10), //Parse Age as an integer
            DreamPlaceToLive: data.DreamPlaceToLive
        })
        .then(response => {
            // Log the response from the server to the console
            setUpdate(!update)
            console.log(response)
        })

        document.querySelector(".user-data").innerHTML = `Welcome ${data.Name}!`
    }

    // Function to handle form submission
    function deleteUser(e, data) {
        e.preventDefault(); // Prevent the default form submission behaviour

        // Send a POST request to the specified URL with the 'data' object
        Axios.get(url + "/deleteUser/" + data.ID)
        .then(response => {
            // Log the response from the server to the console
            setUpdate(!update)
            console.log(response)
        })

        document.querySelector(".user-data").innerHTML = `Removed ${data.Name}!`
    }

    // Render the form component
    return(
        <>
        <form className='user-form' onSubmit={(e) => submit(e)}>
            <TextField id="ID" label="ID" variant="outlined" onChange={(e) => handle(e)}/>
            <TextField id="Name" label="Name" variant="outlined" onChange={(e) => handle(e)}/>
            <TextField id="Age" label="Age" variant="outlined" onChange={(e) => handle(e)}/>
            <TextField id="DreamPlaceToLive" label="Dream Place to Live" variant="outlined" onChange={(e) => handle(e)}/>
            <Button type="submit"  variant="outlined">Submit</Button>
            <p class="user-data"></p>
        </form>
        <TableContainer>
            <Table sx={{ minWidth: 650, maxWidth: 800, margin: 'auto' }} size="small" aria-label="Entries">
                <TableHead>
                    <TableRow>
                    <TableCell align="center">ID</TableCell>
                    <TableCell align="center">Name</TableCell>
                    <TableCell align="center">Age</TableCell>
                    <TableCell align="center">Dream Place to Live</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {users != null ? users.sort((a, b) => a.ID - b.ID).map((item, i) => (
                    <TableRow key={i}>
                        <TableCell align="center">{item.ID}</TableCell>
                        <TableCell align="center">{item.Name}</TableCell>
                        <TableCell align="center">{item.Age}</TableCell>
                        <TableCell align="center">{item.DreamPlaceToLive}</TableCell>
                        <Button variant="outlined" color='error' onClick={(e) => deleteUser(e ,item)}>Delete</Button>
                    </TableRow>
                    ))
                    :
                    <></>
                    }
                </TableBody>
            </Table>
        </TableContainer>
    </>
    )
}
export default PostForm;