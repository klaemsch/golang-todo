import { useEffect, useState } from 'react';

import TodoOverview from "./components/TodoOverview";
import StartPage from "./components/StartPage";
import Loader from "./components/Loader";

import reqClient from "./reqClient";

const App = () => {

  // keeps track wether the app is currently checking wether the session is correct or not
  const [isLoading, setIsLoading] = useState(true);
  // keeps track of the users session / token / id of the todo list
  const [session, setSession] = useState("");

  useEffect(() => {

    // get possible session string from location
    const possibleSessionString = window.location.pathname.substring(1)

    // if the url ends with / -> show normal start page
    if (possibleSessionString.length === 0) {
      setIsLoading(false)
      return
    }

    // if the url has an ending like /abc -> redirect to / and start page
    if (possibleSessionString.length !== 32) {
      location.href = "/"
      return
    }

    // fetch the get all todos endpoint with the possible session string to check wether its valid or not
    reqClient.get(
      '/todo',
      { headers: { 'Authorization': 'Bearer ' + possibleSessionString } }
    ).then((response) => {
      if (response.status === 200) {
        // the session string is valid -> set header and show todo page
        reqClient.defaults.headers.common.Authorization = 'Bearer ' + possibleSessionString
        setSession(possibleSessionString)
        setIsLoading(false)
      }
    }).catch((error) => {
      // an error occured while validating the session string -> redirect to start page
      // maybe its better to show some kind of error to the user
      location.href = "/"
      return
    })
  }, []);

  // show loader if session isnt validated yet
  if (isLoading) {
    return <Loader />
  }

  if (session == "") {
    // no session found -> show start page
    return <StartPage setSession={setSession} />
  } else {
    // session found -> show todo overview
    return <TodoOverview />
  }
}

export default App;
