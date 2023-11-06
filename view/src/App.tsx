import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/public/vite.svg'
import './App.css'
import Navbar from "./components/Navbar/Navbar"
import NoteTitle from "./components/Note/NoteTitle"

function App() { 
  return (
    <>
     <Navbar /> 
     <NoteTitle />
    </>
  )
}

export default App
