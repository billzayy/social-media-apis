import Background from "../../assets/Bg-login.png"
import Logo from "../../assets/Logo-Login.png"
import RegisterComponent from "../../components/layout/RegisterComponent"

const RegisterPage: React.FC = () => {
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
                    <RegisterComponent />
                </div>
            </div>
        </div>
    )
}

export default RegisterPage