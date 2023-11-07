const NoteBody = (props) => {
  return(
    <input className="note-body" name="body" onChange={ props.handler} value={ props.val} type="text" placeholder="Add your note here..." />
  )
}

export default NoteBody
