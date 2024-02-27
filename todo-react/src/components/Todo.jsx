import Icon from '@mdi/react';
import { mdiCheckboxBlankOutline, mdiCheckboxMarkedOutline } from '@mdi/js';
import { mdiArrowUpBoldOutline, mdiArrowDownBoldOutline } from '@mdi/js';
import { useState } from 'react';
import TodoEdit from "./TodoEdit";


// this component represents a todo and shows the todos data as a card
const Todo = (props) => {

  // stores if this component is to show or edit the todo data
  const [editMode, setEditMode] = useState(false);

  // sets the edit mode
  const handleEditTodoButton = () => {
    setEditMode(!editMode)
  }

  // calls the parent element to mark this todo as done / as undone
  const handleMarkButton = () => {
    props.markTodo(props.id)
  }

  // calls the parent element to delete this todo
  const handleDeleteTodoButton = () => {
    props.removeTodo(props.id)
  }

  // calls the parent element to move this todo up
  const handleMoveUpButton = (event) => {
    props.moveTodo(props.id, 1)
  }

  // calls the parent element to move this todo down
  const handleMoveDownButton = (event) => {
    props.moveTodo(props.id, -1)
  }

  // if the todo has no name show unnamed todo
  const todoName = <p>{(props.name == "" ? "Unnamed Todo" : props.name)}</p>

  // map the categories to tags
  const categoryList = props.category?.map((category) => (
    <span className="tag is-primary" key={category}>{category}</span>
  ));

  // the card header shows the todos title, the done/undone button, the categories and the move-up/down buttons
  const cardHeader = (
    <header className="card-header">
      <button className="card-header-icon" onClick={handleMarkButton}>
        <Icon
          className='icon'
          path={props.done ? mdiCheckboxMarkedOutline : mdiCheckboxBlankOutline}
          size={1}
        />
      </button>
      <div className="card-header-title columns is-vcentered">
        <div className='column'>
          {todoName}
        </div>
        <div className="column tags">
          {categoryList}
        </div>
        <div className='column is-1'>
          <button type="button" className="card-header-icon" onClick={handleMoveUpButton}>
            <Icon className='icon' path={mdiArrowUpBoldOutline} size={1} />
          </button>
        </div>
        <div className='column is-1'>
          <button type="button" className="card-header-icon" onClick={handleMoveDownButton}>
            <Icon className='icon' path={mdiArrowDownBoldOutline} size={1} />
          </button>
        </div>
      </div>

    </header>
  )

  // the card body shows the todos text / content
  const cardBody = (
    <div className="card-content">
      <div className="content" style={{ "whiteSpace": "pre-wrap" }}>
        {props.text}
      </div>
    </div>
  )

  // two buttons, one for switching to edit mode, one to delete the todo
  const cardFooter = (
    <footer className="card-footer">
      <button className="button is-ghost card-footer-item" onClick={handleEditTodoButton}>Edit</button>
      <button className="button is-ghost card-footer-item" onClick={handleDeleteTodoButton}>Delete</button>
    </footer>
  )

  /* if the edit mode is turned on, this component gets "replaced" by the TodoEdit component
   *
   * this switch could be made in the parent element, but this way we dont have to store which element
   * is currently edited and the todo component can decide and switch between edit and view mode
   * 
   * TodoEdit gets all the data of this todo as props plus
   * - edit:        Information that the todo is currently edited and not a new one
   * - collapsed:   Information that the card should not be collapsed to view the form
   * - setEditMode: callback to switch back to view mode
   * - editTodo:    Callback to parent Element (TodoOverview) to call the backend with the new data
  */
  if (editMode) {
    return (
      <TodoEdit
        edit={editMode}
        collapsed={false}
        setEditMode={setEditMode}
        editTodo={props.editTodo}
        id={props.id}
        done={props.done}
        name={props.name}
        text={props.text}
        category={props.category}
      />
    )
  } else {
    // in view mode the data is displayed in the card
    return (
      <div className="card block">
        {cardHeader}
        {cardBody}
        {cardFooter}
      </div>
    );
  }

}

export default Todo;
