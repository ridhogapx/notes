interface saveProps {
  title: string,
  handler: (e: any) => void 
}

const SaveNote = (props: saveProps) => {
  return (
    <button className="save-note" onClick={ props.handler }>{ props.title }</button>
  )
}

export default SaveNote
