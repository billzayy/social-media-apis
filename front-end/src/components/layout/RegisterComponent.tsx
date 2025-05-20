import { useState } from "react";
import { RegisterAddition, RegisterReq } from "@/types/Users";
import { useNavigate } from "react-router-dom";

import Steps from "@/components/ui/steps";
import BasicComponent from "./Register/RegisterBasic";
import AdditionComponent from "./Register/RegisterAddition";
import RegisterFinal from "./Register/RegisterFinal";

var stepList: number[] = [1,2,3]

const RegisterComponent: React.FC = () => {
    const [selectedNum, useSelectedNum] = useState<number>(1)
    const [basicReq, setBasicReq] = useState<RegisterReq | undefined>()
    const [addition, setAddition] = useState<RegisterAddition| undefined>()

    const navigate = useNavigate()

    return (
        <div className="flex justify-center items-center mb-5">
            <div>
                <div>
                    <Steps number={stepList} selectedNum={selectedNum} />
                    {selectedNum == 2 ? <AdditionComponent moveToStep={useSelectedNum} addition={setAddition} />:
                            selectedNum == 3 ? <RegisterFinal basicReq={basicReq} addition={addition}/> :
                            <BasicComponent moveToStep={useSelectedNum} reqDispatch={setBasicReq}/>
                        } 
                </div>
                <div className="flex justify-center items-center mt-6 text-sm">
                    <div className="text-gray-500">You're already have an account?</div>
                    <div
                        onClick={() => {navigate("/login")}}
                        className="ml-2 text-amber-400 underline hover:cursor-pointer">Log in!</div>
                </div>
            </div>
        </div>
    )
}

export default RegisterComponent;