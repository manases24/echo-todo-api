import { useEffect, useState } from "react";
import axios from "axios";

export const TaskList = () => {
    const [tasks, setTasks] = useState<{
    [key: string]: { id: number; name: string; description: string, done: boolean };
  }>({});

  const handlePost = async () => {
    try {
      const res = await axios.get("http://localhost:2024/api/v1/tasks");
      setTasks(res.data);
    } catch (error) {
      console.error("Error al obtener los posts:", error);
    }
  };

  useEffect(() => {
    handlePost();
  }, [tasks]);

  const renderedPosts = Object.values(tasks).map(({ id, name, description }) => {
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
  });
  
  return (
    <div className="flex flex-wrap justify-between">
      {renderedPosts}
    </div>
  );
  
};