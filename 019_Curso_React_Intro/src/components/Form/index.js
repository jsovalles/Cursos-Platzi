import { useContext, useState } from "react";
import "./Form.css";
import { TodoContext } from "../Context";

function Form() {
  const { addTodo, setOpenModal } = useContext(TodoContext);

  const [newTodoValue, setNewTodoValue] = useState("");

  const onSubmit = (event) => {
    event.preventDefault();
    addTodo(newTodoValue);
    setOpenModal(false);
  };

  const onCancel = () => {
    setOpenModal(false);
  };

  const onChange = (event) => {
    setNewTodoValue(event.target.value);
  };

  return (
    <form onSubmit={onSubmit}>
      <label>Write your next TODO</label>
      <textarea
        placeholder="Meeting with SCRUM team at 2PM"
        value={newTodoValue}
        onChange={onChange}
      ></textarea>
      <div className="button-container">
        <button
          type="button"
          className="form-button form-button--cancel"
          onClick={onCancel}
        >
          Cancel
        </button>
        <button type="submit" className="form-button form-button--add">
          Add
        </button>
      </div>
    </form>
  );
}

export { Form };
