interface saveProps {
  title: string
}

const SaveNote = (props: saveProps) => {
  return (
    <button className="save-note">{ props.title }</button>
  )
}

export default SaveNote
