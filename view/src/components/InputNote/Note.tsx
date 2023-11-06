import NoteTitle from "./Title"
import NoteBody from "./Body"
import SaveNote from "./Save"

const NoteInput = () => {
   return (
     <div className="input-container">
      <NoteTitle />
      <NoteBody />
      <SaveNote />
      <div className="clear"></div>
    </div>
   )
 }

 export default NoteInput
