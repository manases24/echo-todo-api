import { Link } from "react-router-dom";

export const Header = () => {
  return (
    <div>
      <nav className="bg-gray-50 dark:bg-gray-700">
        <div className="max-w-screen-xl px-4 py-3 mx-auto">
          <div className="flex items-center">
            <ul className="flex flex-row font-medium mt-0 space-x-8 rtl:space-x-reverse text-sm">
              <li className="mr-6">
                <Link className="text-blue-500 hover:text-blue-800" to="/">
                  Home
                </Link>
              </li>
              <li className="mr-6">
                <Link className="text-blue-500 hover:text-blue-800" to="tasks">
                  Tasks
                </Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </div>
  );
};
