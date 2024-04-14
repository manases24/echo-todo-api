import React from 'react'
import ReactDOM from 'react-dom/client'

import { Link, RouterProvider, createBrowserRouter } from 'react-router-dom';

import App from './App.tsx'
import './index.css'


const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>
  },
  // {
  //   path: "search",
  //   element: <TaskList/>
  // },
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    
    <RouterProvider router={router} />
  </React.StrictMode>,
)
