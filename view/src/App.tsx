import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/public/vite.svg'
import './App.css'
import Navbar from "./components/Navbar/Navbar"
import NoteInput from "./components/InputNote/Note"

function App() { 
  return (
    <>
     <Navbar /> 
     <NoteInput />
    </>
  )
}

export default App
