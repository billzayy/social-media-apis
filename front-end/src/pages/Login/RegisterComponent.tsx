import Steps from "@/components/ui/steps";
import { useState } from "react";
import BasicComponent from "./RegisterBasic";
import AdditionComponent from "./RegisterAddition";
import RegisterFinal from "./RegisterFinal";
import { RegisterAddition, RegisterReq } from "@/types/Users";

var stepList: number[] = [1,2,3]

const RegisterComponent: React.FC<{ dispatch: React.Dispatch<React.SetStateAction<string>> }> = ({ dispatch }) => {
    const [selectedNum, useSelectedNum] = useState<number>(1)
    const [basicReq, setBasicReq] = useState<RegisterReq | undefined>()
    const [addition, setAddition] = useState<RegisterAddition| undefined>()

    return (
        <div className="flex justify-center items-center mb-5">
            <div>
                <div>
                    <Steps number={stepList} selectedNum={selectedNum} />
                    {selectedNum == 2 ? <AdditionComponent dispatch={useSelectedNum} addition={setAddition} />:
                            selectedNum == 3 ? <RegisterFinal basicReq={basicReq} addition={addition}/> :
                            <BasicComponent dispatch={useSelectedNum} reqDispatch={setBasicReq}/>
                        } 
                </div>
                <div className="flex justify-center items-center mt-6 text-sm">
                    <div className="text-gray-500">You're already have an account?</div>
                    <div
                        onClick={() => {dispatch("login")}}
                        className="ml-2 text-amber-400 underline hover:cursor-pointer">Log in!</div>
                </div>
            </div>
        </div>
    )
}

export default RegisterComponent;