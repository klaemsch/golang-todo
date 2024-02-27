import { useState } from 'react';

import reqClient from "../reqClient";

const StartPage = (props) => {

  // stores the id of the todo list, at start empty
  // can be filled either by url or by input
  const [listIdState, setListIdState] = useState("");

  // every time the input value for the listId input field changes update the state
  const handleListIdValueChange = (event) => {
    setListIdState(event.target.value)
  }

  // redirects the user to the todo list with id stored in list id state
  const loadTodoList = (event) => {
    event.preventDefault();
    location.href = '/' + listIdState
  }

  // calls the backend to create a new todo list and redirects the user to the list
  const newTodoList = async () => {
    let response = await reqClient.get('/list')
    if (response.data !== null) location.href = '/' + response.data;
  }

  // start page with two colums
  // -> left opens an existing list after the Todo-List-Identifier was given
  // -> right creates a new todo list
  return (
    <section className="section">
      <div className="container">
        <h1 className="title">Welcome to your Todo App</h1>
        <h5 className="subtitle is-6">Open an existing list or create a new one</h5>

        <div className="columns">
          <div className="column">
            <h5 className="subtitle is-5"><strong>Option 1: </strong>Open Existing Todo List</h5>
            <p>
              You have used the app before? Enter your Todo-List password to access your existing todo list. You can share the password with friends or coworkers to collaborate seamlessly.
            </p>
            <br />
            <form onSubmit={loadTodoList}>
              <div className="field">
                <div className="control">
                  <input
                    required
                    className="input"
                    type="text"
                    placeholder="Your Todo-List password"
                    value={listIdState}
                    onChange={handleListIdValueChange}
                  />
                </div>
              </div>
              <div className="field">
                <div className="control">
                  <button type="submit" className="button is-primary">Load Todo-List</button>
                </div>
              </div>

            </form>
          </div>
          <div className="column">
            <h5 className="subtitle is-5"><strong>Option 2: </strong>Start a New Todo List</h5>
            <p>
              Ready for a fresh start? Create a new todo list and begin organizing your tasks.
            </p>
            <br />
            <button className="button is-primary" onClick={newTodoList}>Create New Todo-List</button>
          </div>
        </div>

      </div>
    </section>
  )
}

export default StartPage;
