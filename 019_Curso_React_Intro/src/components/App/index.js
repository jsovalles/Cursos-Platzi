import "./App.css";
import "@fortawesome/fontawesome-free/css/all.min.css";
import { AppUI } from "./AppUI";
import { TodoProvider } from "../Context";

function App() {
  return (
    <TodoProvider>
      <AppUI></AppUI>
    </TodoProvider>
  );
}

export default App;
