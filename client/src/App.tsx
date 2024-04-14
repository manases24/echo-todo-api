import { Header } from "./components/Header";
import { TaskCreate } from "./components/TaskCreate";
import { TaskList } from "./components/TaskList";

function App() {
  return (
    <div className="container mx-auto px-4 my-2.5">
      <Header />
      <TaskCreate />
      <h1>Tasks</h1>
      <TaskList/>
    </div>
  );
}

export default App;
