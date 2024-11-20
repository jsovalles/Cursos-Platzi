import "./Item.css";
import { CompleteIcon } from "../Icon/CompleteIcon";
import { DeleteIcon } from "../Icon/DeleteIcon";

function TodoItem({ text, completed, onComplete, onDelete }) {
  return (
    <li className="TodoItem">
      <CompleteIcon
        completed={completed}
        onComplete={onComplete}
      />
      <p className={`TodoItem-p ${completed && "TodoItem-p--complete"}`}>
        {text}
      </p>
      <DeleteIcon onDelete={onDelete} />
    </li>
  );
}

export { TodoItem };
