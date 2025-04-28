import LoginComponent from "./LoginComponent"
import Background from "../../assets/Bg-login.png"
import Logo from "../../assets/Logo-Login.png"
import { useState } from "react"
import RegisterComponent from "./RegisterComponent"

const LoginPage: React.FC = () => {
    const [section, useSection] = useState("login")
    return (
        <div className="flex bg-black text-white w-screen h-screen items-center justify-center">
            <div className="picture">
                <img src={Background} alt="Background Image" className="max-w-screen max-h-[125vh]"/>
            </div>
            <div className="h-screen w-full flex justify-center items-center">
                <div className="text-center">
                    <div className="flex justify-center my-5 mx">
                        <img src={Logo} alt="logo" className="w-[70%]" />
                    </div>
                    {section != "register" ?
                        <LoginComponent dispatch={useSection} /> :
                        <RegisterComponent dispatch={useSection} />} 
                </div>
            </div>
        </div>
    )
}

export default LoginPage