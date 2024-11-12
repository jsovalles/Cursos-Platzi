import React from "react";
import { TodoIcon } from "./TodoIcon.js";
import "../styles/TodoIcon.css";

function CompleteIcon({ completed, onComplete }) {
  return (
    <TodoIcon
      typeColor={completed && "Icon-check--active"}
      type="check"
      onClick={onComplete}
    />
  );
}

export { CompleteIcon };
