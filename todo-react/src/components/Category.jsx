import { mdiClose, mdiPlus } from '@mdi/js';
import Icon from '@mdi/react';
import { useState } from 'react';

const Category = (props) => {

  // if true render input for new category, if false render category
  const isAdd = props.isAdd

  // saves the current user input for the category name
  const [newCategoryName, setNewCategoryName] = useState("");

  // removes the category from the parent todo
  const handleRemove = () => {
    props.removeCategory(props.name)
  }

  // adds the category with the current input to the parent todo
  const handleAdd = (event) => {
    event.preventDefault();
    // only non empty todos are valid
    if (newCategoryName === '') return
    // add category to parent todo
    props.addCategory(newCategoryName)
    // clear input
    setNewCategoryName("")
  }

  // every time the category name gets changed update the value state
  const handleCategoryNameChange = (event) => {
    setNewCategoryName(event.target.value)
  }

  // input field for category name
  const inputField = (
    <input
      type="text"
      className='input is-small is-primary is-rounded'
      placeholder="Name of your category"
      autoComplete="off"
      value={newCategoryName}
      onChange={handleCategoryNameChange}
    />
  )

  // button adds the current name to the categories
  const addButton = (
    <button className='button is-small is-primary' onClick={handleAdd}>
      <Icon className='icon' path={mdiPlus} size={1} />
    </button>
  )

  // button removes the category from the todo
  const removeButton = (
    <button type="button" onClick={handleRemove} className='button is-small is-primary'>
      <Icon className='icon' path={mdiClose} size={1} />
    </button>
  )

  return (
    <span className="tag is-large is-primary is-rounded">
      {isAdd ? inputField : props.name}
      {isAdd ? addButton : removeButton}
    </span>
  )
}

export default Category