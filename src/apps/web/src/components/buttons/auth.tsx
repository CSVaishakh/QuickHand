import { Link } from "@tanstack/react-router"

interface AuthButtoProps {
    name: string
    path: string
}

const AuthButton = ( props:  AuthButtoProps ) => {
    return (
      <Link to={props.path}>
        <button className="px-4 py-1 border-2 rounded-xl border-black bg-white font-semibold text-xl text-blue-400 hover:bg-blue-400 hover:text-white ">
          {props.name}
            </button>
      </Link>
  )
}

export default AuthButton