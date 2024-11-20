import { useContext } from "react";
import "./Search.css";
import { TodoContext } from "../Context";

function TodoSearch() {
  const { searchValue, setSearchValue } = useContext(TodoContext);

  return (
    <div className="search-container">
      <input
        className="todo-search"
        placeholder="Your task"
        value={searchValue}
        onChange={(event) => {
          setSearchValue(event.target.value);
        }}
      />
    </div>
  );
}

export { TodoSearch };
