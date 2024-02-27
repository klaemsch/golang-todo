import { useState, useEffect } from 'react';
import Icon from '@mdi/react';
import { mdiLinkVariant } from '@mdi/js';

import Todo from "./Todo";
import TodoEdit from "./TodoEdit";
import reqClient from '../reqClient'

// import bulma-tooltip, an bulma-add on to show nice tooltips
import "../../node_modules/@creativebulma/bulma-tooltip/dist/bulma-tooltip.css"

// component that lists all todos as individual cards and a card for adding new todos
// also renders a share button
const TodoOverview = () => {

  // state where todos will be saved
  const [todos, setTodosWrapped] = useState([]);

  // wraps the setter of the todos state to sort them every time the todos are updated
  // maybe better to sort them backend-side
  const setTodos = (todos) => {
    const newTodos = todos.sort((a, b) => {
      return a.rank > b.rank ? -1 : 1
    })
    setTodosWrapped(newTodos)
  }

  // on component load: fetch all todos from backend and save the response in the todo state
  useEffect(() => {
    const fetchTodos = async () => {
      let response = await reqClient.get('/todo')
      if (response.data !== null) setTodos(response.data);
    }
    fetchTodos()
  }, []);

  // uses given data to call the backend to create a new todo, adds the response to the todo state
  const addTodo = async (done, name, text, category) => {
    let response = await reqClient.post('/todo', { done: done, name: name, text: text, category: category })
    if (response.status == 200 && response.data !== null) setTodos([...todos, response.data]);
  }

  // uses given id to call the backend to remove the todo, removesit from the todo state
  const removeTodo = async (id) => {
    await reqClient.delete('/todo', {
      params: { id: id }
    })
    setTodos(
      todos.filter((todo) => {
        return todo.id !== id;
      })
    );
  };

  // uses given data to call the backend to edit a todo, updates the todo state
  const editTodo = async (id, done, name, text, category) => {
    
    // find todo with id locally
    let todoToEdit = todos.find((todo) => {
      return todo.id === id
    })
    if (todoToEdit === undefined) return

    // update data to local todo
    todoToEdit.done = done
    todoToEdit.name = name
    todoToEdit.text = text
    todoToEdit.category = category

    // send updated data to backend
    let response = await reqClient.put('/todo', todoToEdit)
    if (response.status !== 200 && response.data !== null) return

    // replace todo with updated response
    setTodos(
      todos.map((todo) => {
        if (todoToEdit.id === todo.id) {
          return response.data
        } else {
          return todo
        }
      })
    );
  };

  // uses given id to call the backend to mark a todo as done/undone, updates the todo state
  const markTodo = async (id) => {

    // find todo locally
    let markedTodo = todos.find((todo) => {
      return todo.id === id
    })
    if (markedTodo === null) return

    // update done state of todo locally
    markedTodo.done = !markedTodo.done

    // send updated todo to backend
    let response = await reqClient.put('/todo', markedTodo)
    if (response.status !== 200 && response.data !== null) return

    // replace todo with updated response
    setTodos(
      todos.map((todo) => {
        if (markedTodo.id === todo.id) {
          return response.data
        } else {
          return todo
        }
      })
    );
  };

  // uses given data to call the backend to move a todo, updates the todo state
  const moveTodo = async (id, rankDif) => {

    // find todo locally
    let markedTodo = todos.find((todo) => {
      return todo.id === id
    })
    if (markedTodo === null) return

    // edit todo
    markedTodo.rank += rankDif

    // send todo data to backend
    let response = await reqClient.put('/todo', markedTodo)
    if (response.status !== 200 && response.data !== null) return

    // replace todo with updated response
    setTodos(
      todos.map((todo) => {
        if (markedTodo.id === todo.id) {
          return response.data
        } else {
          return todo
        }
      })
    );
  };

  // writes share url to clipboard
  // adds tooltip css classes to button to show tooltip
  // removes classes after 2sec
  const handleShareButton = (event) => {
    navigator.clipboard.writeText(window.location.href);
    const target = event.currentTarget
    target.classList.add('has-tooltip-active')
    target.classList.add('has-tooltip-left')
    target.setAttribute("data-tooltip", "Copied to Clipboard")
    window.setTimeout(() => {
      target.classList.remove('has-tooltip-active')
      target.classList.remove('has-tooltip-left')
      target.removeAttribute("data-tooltip")
    }, 2000)
  }

  // maps todos to components; pass data, mark, move, edit, remove callbacks
  const todoComponentList = todos.map((todo) => (
    <Todo
      key={todo.id}
      id={todo.id}
      name={todo.name}
      text={todo.text}
      category={todo.category}
      done={todo.done}
      markTodo={markTodo}
      moveTodo={moveTodo}
      editTodo={editTodo}
      removeTodo={removeTodo}
    />
  ));


  // main html structure of the app, with heading, share button, todos and the add-Todo-card
  return (
    <section className="section">
      <div className="container">
        <nav className="level">
          <div className="level-left">
            <h1 className="level-item title">Todo List</h1>
          </div>
          <div className="level-right">
            <button
              type="button"
              className='level-item button is-primary'
              onClick={handleShareButton}
            >
              <span>Share your Todo-List</span>
              <Icon className="icon" path={mdiLinkVariant} size={1} />
            </button>
          </div>
        </nav>
        <TodoEdit addTodo={addTodo} />
        {todoComponentList}
      </div>
    </section>
  );
}

export default TodoOverview;