import { TodoCounter } from "./TodoCounter";
import { TodoSearch } from "./TodoSearch";
import { TodoList } from "./TodoList";
import { TodoItem } from "./TodoItem";
import { CreateTodoButton } from "./CreateTodoButton";
import "../styles/App.css";
import React from "react";
import "@fortawesome/fontawesome-free/css/all.min.css";

function App() {
  const [searchValue, setSearchValue] = React.useState("");

  const [todos, setTodos] = React.useState([
    { text: "first", completed: true },
    { text: "second", completed: false },
    { text: "third", completed: false },
  ]);

  const completedTodos = todos.filter((todo) => todo.completed).length;
  const totalTodos = todos.length;

  const searchedTodos = todos.filter((todo) =>
    todo.text.toLowerCase().includes(searchValue.toLowerCase())
  );

  const onComplete = (text) => {
    const todoIndex = todos.findIndex((todo) => todo.text === text);
    const newTodos = [...todos];
    newTodos[todoIndex].completed = !newTodos[todoIndex].completed;
    setTodos(newTodos);
  };

  const onDelete = (text) => {
    const newTodos = todos.filter((todo) => todo.text !== text);
    setTodos(newTodos);
  };

  return (
    <>
      <div className="todo-container">
        <TodoCounter
          completed={completedTodos}
          total={totalTodos}
          className="todo-counter"
        />
        <div className="search-container">
          <TodoSearch
            className="todo-search"
            searchValue={searchValue}
            setSearchValue={setSearchValue}
          />
        </div>

        <TodoList>
          {searchedTodos.map((todo) => (
            <TodoItem
              key={todo.text}
              text={todo.text}
              completed={todo.completed}
              onComplete={() => onComplete(todo.text)}
              onDelete={() => onDelete(todo.text)}
            />
          ))}
        </TodoList>
      </div>
      <CreateTodoButton />
    </>
  );
}

export default App;
