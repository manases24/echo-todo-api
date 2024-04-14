import { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

export const TaskList = () => {
  const [tasks, setTasks] = useState<{
    [key: string]: {
      id: number;
      name: string;
      description: string;
      done: boolean;
    };
  }>({});

  const handleTask = async () => {
    try {
      const res = await axios.get("http://localhost:2024/api/v1/tasks");
      console.log(res.data);
      setTasks(res.data.data);
    } catch (error) {
      console.error("Error al obtener las tareas:", error);
    }
  };

  useEffect(() => {
    handleTask();
  }, []); // Se ejecuta solo una vez al cargar el componente

  const renderedTasks = Object.values(tasks).map(
    ({ id, name, description }) => {
      if (!id) {
        return null;
      }
      return (
        <div
          className="flex-shrink-0 container mx-auto bg-white shadow-md rounded-lg w-1/3 p-4 mb-4 border border-pink-500"
          style={{ width: "30%", marginBottom: "20px" }}
          key={id}
        >
          <div className="border-b border-pink-500 mb-4 pb-2">
            <h3 className="text-xl font-bold">{name}</h3>
            <p className="text-xl">{description}</p>
          </div>
        </div>
      );
    }
  );

  return (
    <>
      <header className="my-2.5">
        <nav className="bg-gray-50 dark:bg-gray-700">
          <div className="max-w-screen-xl px-4 py-3 mx-auto">
            <div className="flex items-center">
              <ul className="flex flex-row font-medium mt-0 space-x-8 rtl:space-x-reverse text-sm">
                <li className="mr-6">
                  <Link className="text-blue-500 hover:text-blue-800" to="/">
                    Home
                  </Link>
                </li>
              </ul>
            </div>
          </div>
        </nav>
      </header>
      <div className="flex flex-wrap justify-between">{renderedTasks}</div>
    </>
  );
};
