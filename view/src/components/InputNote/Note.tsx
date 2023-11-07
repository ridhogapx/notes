import { useState } from "react"
import axios from "axios"

import NoteTitle from "./Title"
import NoteBody from "./Body"
import SaveNote from "./Save"

interface noteData {
  title: string, 
  body: string,
}

const NoteInput = () => {
  const [note, setNote] = useState<noteData>({
      title: "",
      body: "",
    }) 

  const handleOnChange = (e: any) => {
    setNote({
      ...note,
      [e.target.name]: e.target.value
      })
  }

   return (
     <div className="input-container">
      <NoteTitle title="Title" val={note.title} handler={ handleOnChange}/>
      <NoteBody handler={ handleOnChange } val={ note.body} />
      <SaveNote title="Save" handler={() => { console.log("foo")}} />
      <div className="clear"></div>
    </div>
   )
 }

 export default NoteInput
