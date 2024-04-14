// import React from "react";

// import { TaskData } from "../types";

// interface TaskResultProps {
//   task: TaskData | null; 
// }

// export const TaskResult: React.FC<TaskResultProps> = ({ id, name, description, done }) => {


//   if (!id) {
//     // Si no hay tarea creada, mostrar un mensaje indicando que no hay tarea
//     return <div>No task created yet</div>;
//   }

//   // Si hay tarea creada, mostrar los detalles de la tarea
//   return (
//      <div
//         className="flex-shrink-0 container mx-auto bg-white shadow-md rounded-lg w-1/3 p-4 mb-4 border border-pink-500"
//         style={{ width: "30%", marginBottom: "20px" }}
//         key={id}
//       >
//         <div className="border-b border-pink-500 mb-4 pb-2">
//           <h3 className="text-xl font-bold">{name}</h3>
//           <p className="text-xl">{description}</p>
//           <p>Done: {done ? "Yes" : "No"}</p>
//         </div>
//       </div>
//   );
// };
