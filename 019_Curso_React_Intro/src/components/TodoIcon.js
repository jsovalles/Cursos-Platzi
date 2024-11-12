import "../styles/TodoIcon.css";

const iconTypes = {
  check: "check",
  delete: "times",
};

function TodoIcon({ type, onClick, typeColor }) {
  return (
    <span className={`Icon Icon-${type} ${typeColor}`} onClick={onClick}>
      {console.log(iconTypes[type])}
      <i className={`fas fa-${iconTypes[type]}`}> </i>
    </span>
  );
}

export { TodoIcon };
