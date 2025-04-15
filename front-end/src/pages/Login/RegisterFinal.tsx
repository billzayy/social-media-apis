import { Button } from "@/components/ui/button"
import { Link } from "react-router-dom"

const RegisterFinal: React.FC = () => { 
    return (
        <div>
            <div className="font-bold text-xl my-5">All set !</div>
            <div className="mb-3">Welcome to the Social Media ! Have fun </div>
            <Button asChild className="mt-4 w-[100%] bg-amber-500 text-black hover:cursor-pointer hover:text-white">
                <Link to={"/"}>Go to Homepage</Link>
            </Button>
        </div>
    )
}

export default RegisterFinal