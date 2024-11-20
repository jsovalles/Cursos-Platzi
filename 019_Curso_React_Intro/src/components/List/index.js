import './List.css';

function TodoList({ children }) {
  return (
    <ul>
      {children}
    </ul>
  );
}

export { TodoList };
