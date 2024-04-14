import React, { useState } from "react";
import axios from "axios";

import { TaskData } from "../types";

export const TaskCreate: React.FC = () => {
  const [task, setTask] = useState<TaskData>({
    name: "",
    description: "",
    done: false,
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value, type } = e.target;
    // Manejar diferentes tipos de campos de entrada
    const newValue =
      type === "checkbox" ? (e.target as HTMLInputElement).checked : value;
    setTask({
      ...task,
      [name]: newValue,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await axios.post("http://localhost:2024/api/v1/task", task);
      console.log("Task created successfully:", task);
      setTask({
        name: "",
        description: "",
        done: false,
      });
    } catch (error) {
      console.error("Error al enviar la solicitud:", error);
    }
  };

  return (
    <>
      <div className="mt-4">
        <form onSubmit={handleSubmit}>
          <label htmlFor="name" className="block">
            <span className="text-gray-700">Task name</span>
            <input
              type="text"
              id="name"
              name="name"
              value={task.name}
              onChange={handleChange}
              placeholder="Task name"
              className="form-input mt-1 block w-1/2 border-current outline-pink-500 rounded-lg border border-pink-500"
            />
          </label>
          <label htmlFor="description" className="block mt-4">
            <span className="text-gray-700">Description</span>
            <textarea
              id="description"
              name="description"
              value={task.description}
              onChange={handleChange}
              placeholder="Task description"
              className="form-textarea mt-1 block w-1/2 border-current outline-pink-500 rounded-lg border border-pink-500"
            ></textarea>
          </label>
          <label htmlFor="done" className="block mt-4">
            <input
              type="checkbox"
              id="done"
              name="done"
              checked={task.done}
              onChange={handleChange}
              className="form-checkbox"
            />
            <span className="ml-2 text-gray-700">Task done</span>
          </label>
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded my-2"
          >
            Submit
          </button>
        </form>
      </div>
    </>
  );
};
