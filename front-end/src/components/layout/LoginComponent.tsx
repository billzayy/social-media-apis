import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { LoginService } from "../../services/LoginService";

const LoginComponent: React.FC = () => { 
    const [userName, setUserName] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const navigate = useNavigate()

    return (
        <div className="block">
            <div>Welcome, login to your account!</div>
            <form action="" className="my-5">
                <div>
                    <input
                        value={userName}
                        onChange={(e) => {setUserName(e.target.value)}}
                        type="text" placeholder="UserName" className="border border-black border-b-white text-white my-3.5 ml-2.5 w-[80%] py-1 focus:outline-none" />
                </div>
                <div>
                    <input
                        value={password}
                        onChange={(e) => {setPassword(e.target.value)}}
                        type="password" placeholder="Password" className="border border-black border-b-white text-white my-3.5 ml-2.5 w-[80%] py-1 focus:outline-none" />
                </div>
            </form>

            <div className="flex justify-around mb-3">
                <div className="flex items-center">
                    <Switch className="h-5 mr-1.5 bg-gradient-to-r from-amber-400 hover:cursor-pointer"/>
                    <p>Remember me</p>
                </div>
                <div className="text-gray-500 hover:cursor-pointer font-light underline">Forgot Password?</div>
            </div>

            <Button
                onClick={() => {LoginService(userName, password, navigate)}}
                size="default"
                className="w-[80%] my-4 bg-amber-400 text-black h-[60%] hover:cursor-pointer hover:text-white hover:bg-gray-700">
                Login
            </Button>
            <div className="flex justify-center mt-2 text-sm">
                <p className="text-gray-500 font-thin">Don't have account yet?</p>
                <p
                    className="text-amber-400 ml-2 font-light hover:cursor-pointer"
                    onClick={() => {navigate("/register")}}
                >Sign Up</p>
            </div>
        </div>
    )
}

export default LoginComponent;