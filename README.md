# Todo App

This is a simple Todo application built using Golang and React. Follow the steps below to get the app up and running on your local machine.

## Getting Started

### 1. Clone Repo

```bash
git clone https://github.com/klaemsch/golang-todo.git
cd golang-todo
```

### 2. Set Up Frontend

```bash
cd todo-react
npm install
npm run dev
```

This will install the necessary frontend packages and start the development server.

### 3. Set Up Backend

Open a new terminal in the project directory.

```bash
cd todo
go run main.go
```

This will run the Golang backend server.

### 4. Open the App

Open your browser and navigate to [http://localhost:5173](http://localhost:5173) to access the Todo app.

## Features

### User Authentication
- **Login:** Users can view only their own Todo items.

### Todo Management
- **Create Todo:** Users can create a new Todo task with a title and description and save it in the database.
- **List Todo:** Users can see all their Todo items in a list.
- **Delete Todo:** Users can delete a Todo from the list when it is no longer needed.
- **Mark as Completed:** Users can mark a Todo as completed.
- **Update Todo:** Users can modify an existing Todo.

### Todo Enhancements
- **Categorize Todo:** Users can categorize their Todo tasks.
- **Share Todo:** Users can share their Todos with other users, allowing them to view and mark it as completed.

## Future Development
- **Reorder Todo:** Users will have the ability to rearrange the order of their Todo tasks.
- **Filter and Search:** Users will be able to filter ToDo tasks by category, status (completed, not completed), and search for tasks using keywords.

## Contributing

Feel free to contribute to the development of this Todo app by submitting issues or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).