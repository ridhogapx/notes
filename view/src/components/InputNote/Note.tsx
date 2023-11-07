import NoteTitle from "./Title"
import NoteBody from "./Body"
import SaveNote from "./Save"

const NoteInput = () => {
   return (
     <div className="input-container">
      <NoteTitle title="Title" />
      <NoteBody />
      <SaveNote title="Save" />
      <div className="clear"></div>
    </div>
   )
 }

 export default NoteInput
