const NoteTitle = (props) => {
  return ( 
      <input className="note-title" type="text" name="title" placeholder={props.title} value={ props.val} onChange={ props.handler} />
  )
}

export default NoteTitle
