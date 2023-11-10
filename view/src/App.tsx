import { useState, useEffect } from 'react'
import axios from 'axios'
import './App.css'
import Navbar from "./components/Navbar/Navbar"
import NoteInput from "./components/InputNote/Note"

function App() {
  useEffect(() => {
      axios.get("/api/v1/notes").then((res) => {
          console.log(res.data.message)
        })
    }, [])

  return (
    <>
     <Navbar /> 
     <NoteInput />
    </>
  )
}

export default App
