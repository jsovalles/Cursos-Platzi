import { useContext } from "react";
import "./Button.css";
import { TodoContext } from "../Context";

function CreateTodoButton() {
  const { setOpenModal } = useContext(TodoContext);

  return (
    <button
      className="floating-button"
      onClick={() => {
        setOpenModal((value) => !value);
      }}
    >
      +
    </button>
  );
}

export { CreateTodoButton };
