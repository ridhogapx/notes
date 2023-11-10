import { useState } from "react"
import axios from "axios"
import toastr from "toastr"

import NoteTitle from "./Title"
import NoteBody from "./Body"
import SaveNote from "./Save"
import ClearNote from "./Clear"

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
    const res = await axios.post("/api/v1/note", {
      title: note.title,
      body: note.body,
    })

    toastr.success(res.data.message)
  }

  const handlerClear = () => {
    setNote({
        title: "",
        body: "",
      })
  }

   return (
     <div className="input-container">
      <NoteTitle title="Title" val={note.title} handler={ handleOnChange}/>
      <NoteBody handler={ handleOnChange } val={ note.body} />
      <ClearNote handler={ handlerClear } />
      <SaveNote title="Save" handler={ handleSave } />
      <div className="clear"></div>
    </div>
   )
 }

 export default NoteInput
