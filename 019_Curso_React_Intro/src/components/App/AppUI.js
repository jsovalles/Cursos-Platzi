import { TodoCounter } from "../Counter";
import { TodoSearch } from "../Search";
import { TodoList } from "../List";
import { TodoItem } from "../Item";
import { CreateTodoButton } from "../Button";
import { Loading } from "../Loading";
import { Error } from "../Error";
import { EmptyTodo } from "../EmptyTodo";
import { TodoContext } from "../Context";
import { useContext } from "react";
import { Modal } from "../Modal";
import { Form } from "../Form";

function AppUI() {
  const { searchedTodos, onComplete, onDelete, loading, error, openModal } =
    useContext(TodoContext);

  return (
    <>
      <div className="todo-container">
        <TodoCounter />
        <TodoSearch />

        <TodoList>
          {loading && <Loading />}
          {error && <Error />}
          {!loading && !searchedTodos.length && <EmptyTodo />}

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

      {openModal && (
        <Modal>
          <Form></Form>
        </Modal>
      )}
    </>
  );
}

export { AppUI };
