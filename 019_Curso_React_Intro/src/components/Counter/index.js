import { useContext } from "react";
import "./Counter.css";
import { TodoContext } from "../Context";

function TodoCounter() {
  const { completedTodos, totalTodos } = useContext(TodoContext);

  return (
    <h1>
      You've completed {completedTodos} of {totalTodos} TODOS
    </h1>
  );
}

export { TodoCounter };
