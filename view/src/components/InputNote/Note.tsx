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

  const handleOnChange = (e) => {
    setNote({
      ...note,
      [e.target.name]: e.target.value
      })
  }

  const handleSave = async(e) => {
    const res = await axios.post("http://localhost:8080/api/v1/note", {
      title: note.title,
      body: note.body,
    })

    console.log(res.data)
  }

   return (
     <div className="input-container">
      <NoteTitle title="Title" val={note.title} handler={ handleOnChange}/>
      <NoteBody handler={ handleOnChange } val={ note.body} />
      <SaveNote title="Save" handler={ handleSave } />
      <div className="clear"></div>
    </div>
   )
 }

 export default NoteInput
