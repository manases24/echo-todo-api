import { Header } from "./components/Header";
import { TaskCreate } from "./components/TaskCreate";
// import { TaskResult } from "./components/TaskResult";

function App() {
  return (
    <div className="container mx-auto px-4 my-2.5">
      <Header />
      <TaskCreate />
    </div>
  );
}

export default App;
