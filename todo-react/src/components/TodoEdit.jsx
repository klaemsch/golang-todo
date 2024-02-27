import { useState, useEffect } from 'react';
import Icon from '@mdi/react';
import { mdiArrowDown, mdiArrowUp, mdiCheckboxBlankOutline, mdiCheckboxMarkedOutline } from '@mdi/js';

import Category from './Category';

/* card component
 * can collapse
 * has a form to edit already existing todos
 * can use the same form to create new todos
 */
const TodoEdit = (props) => {

  // states for storing wether or not the component is collapsed
  const [collapsedState, setCollapsedState] = useState(true);

  // states for storing the input data
  const [doneState, setDoneState] = useState(false);
  const [nameState, setNameState] = useState("");
  const [textState, setTextState] = useState("");
  const [categoryState, setCategoryState] = useState([]);

  // if set, load todo values from props (important for editing)
  useEffect(() => {
    if (props.collapsed != undefined) setCollapsedState(props.collapsed)
    if (props.done != undefined) setDoneState(props.done)
    if (props.name != undefined) setNameState(props.name)
    if (props.text != undefined) setTextState(props.text)
    if (props.category != undefined) setCategoryState(props.category)
  }, []);

  // reads the input states and calls the parent to call the backend to create or update the todo
  const handleAddButton = (event) => {
    event.preventDefault();
    if (props.edit) {
      // if edit mode -> call edit todo with updated data and close the edit mode
      console.log(categoryState)
      props.editTodo(props.id, doneState, nameState, textState, categoryState)
      props.setEditMode(false)
    } else {
      // if add mode -> call add todo with new data and empty the input fields
      props.addTodo(doneState, nameState, textState, categoryState)
      setDoneState(false)
      setNameState("")
      setTextState("")
      setCategoryState([])
    }
  }

  // every click flips the done / undone state
  const handleDoneValueChange = (event) => {
    setDoneState(!doneState)
  }

  // update the state values with current input data
  const handleNameValueChange = (event) => {
    setNameState(event.target.value)
  }

  // update the state values with current input data
  const handleTextValueChange = (event) => {
    setTextState(event.target.value)
  }

  // every call flips the collapsed state
  const handleCollapsedChange = () => {
    setCollapsedState(!collapsedState)
  }

  // adds a new category string to the local category store
  // does not change the todo yet
  const addCategory = async (name) => {
    setCategoryState([...categoryState, name]);
  }

  // removes a category from the store
  // does not change the todo yet
  const removeCategory = async (name) => {
    setCategoryState(
      categoryState.filter((category) => {
        return category !== name;
      })
    );
  };

  // cancel edit -> set editMode in upper level todo to false -> will render normal todo view
  const handleCancelButton = () => {
    props.setEditMode(false)
  }

  // map category components for each category in store
  const categorieComponents = (
    categoryState.map((category) => (
      <Category
        key={category}
        name={category}
        removeCategory={removeCategory}
      />
    ))
  )

  // shows mode (Edit or Create) and a collapse icon-button
  const cardHeader = (
    <header className="card-header is-clickable" onClick={handleCollapsedChange}>
      <div className="card-header-title">
        {props.edit ? "Edit Todo" : "Create Todo"}
      </div>
      <button className="card-header-icon">
        <Icon className='icon' path={collapsedState ? mdiArrowDown : mdiArrowUp} size={1} />
      </button>
    </header>
  )

  // shows the "form" with status, name, text and categories
  const cardBody = (
    <div className="card-content">
      <div className="content">

        <div className="field">
          <label className="label">Status</label>
          <button type="button" className="card-header-icon" onClick={handleDoneValueChange}>
            <Icon
              className='icon'
              path={doneState ? mdiCheckboxMarkedOutline : mdiCheckboxBlankOutline}
              size={1}
            />
          </button>
        </div>

        <div className="field">
          <label className="label">Name</label>
          <div className="control">
            <input
              required
              className="input"
              type="text"
              placeholder="Name of your todo"
              autoComplete="off"
              value={nameState}
              onChange={handleNameValueChange}
            />
          </div>
        </div>

        <div className="field">
          <label className="label">Todo Content</label>
          <div className="control">
            <textarea
              className="textarea"
              placeholder="Describe your todo"
              value={textState}
              onChange={handleTextValueChange}
            />
          </div>
        </div>

        <div className="tags">
          {categorieComponents}
          <Category
            isAdd={true}
            addCategory={addCategory}
          />
        </div>
      </div>
    </div>
  )

  // shows buttons for saving/adding and canceling
  const cardFooter = (
    <footer className="card-footer">
      <button className="button is-ghost card-footer-item" type="submit" >{props.edit ? "Save" : "Add"}</button>
      <button className="button is-ghost card-footer-item" type="button" onClick={props.edit ? handleCancelButton : handleCollapsedChange}>Cancel</button>
    </footer>
  )

  // the double <></<> should be replaced with a nice if else condition block using vars
  return (
    <div className="card block">
      {cardHeader}
      <form action="" onSubmit={handleAddButton}>
        {collapsedState ? <></> : cardBody}
        {collapsedState ? <></> : cardFooter}
      </form>
    </div>
  );
}

export default TodoEdit;