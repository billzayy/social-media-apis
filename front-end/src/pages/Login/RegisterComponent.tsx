import Steps from "@/components/ui/steps";
import { useState } from "react";
import BasicComponent from "./RegisterBasic";
import AdditionComponent from "./RegisterAddition";
import RegisterFinal from "./RegisterFinal";

var stepList: number[] = [1,2,3]

const RegisterComponent: React.FC<{ dispatch: React.Dispatch<React.SetStateAction<string>> }> = ({ dispatch }) => {
    const [selectedNum, useSelectedNum] = useState(1)

    return (
        <div className="flex justify-center items-center mb-5">
            <div>
                <div>
                    <Steps number={stepList} selectedNum={selectedNum} />
                    {selectedNum == 2 ? <AdditionComponent dispatch={useSelectedNum} />:
                            selectedNum == 3 ? <RegisterFinal/> :
                            <BasicComponent dispatch={useSelectedNum} />
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