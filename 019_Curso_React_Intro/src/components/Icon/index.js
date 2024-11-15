import "./Icon.css";

const iconTypes = {
  check: "check",
  delete: "times",
};

function TodoIcon({ type, onClick, typeColor }) {
  return (
    <span className={`Icon Icon-${type} ${typeColor}`} onClick={onClick}>
      <i className={`fas fa-${iconTypes[type]}`}> </i>
    </span>
  );
}

export { TodoIcon };
