import { Button } from "@/components/ui/button";
import React from "react";

const BasicComponent: React.FC<{ dispatch: React.Dispatch<React.SetStateAction<number>> }> = ({dispatch}) => { 
    return (
        <div className="text-left pl-4">
            <div className="mt-1 mb-3 text-2xl">Basic Information</div>
            <form action="">
                <div>
                    <input required type="text" placeholder="Login" className="border border-black border-b-white text-white my-2.5 w-[92%] py-1 focus:outline-none" />
                </div>
                <div>
                    <input required type="email" placeholder="Email" className="border border-black border-b-white text-white my-2.5 w-[92%] py-1 focus:outline-none" />
                </div>
                <div className="flex">
                    <div>
                        <input required type="text" placeholder="Name" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                    <div>
                        <input required type="text" placeholder="Surname" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                </div>
                <div>
                    <input required type="password" placeholder="Password" className="border border-black border-b-white text-white my-2.5 w-[92%] py-1 focus:outline-none" />
                </div>
                <div className="flex">
                    <div>
                        <input required type="text" placeholder="City" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                    <div>
                        <input required type="text" placeholder="Age" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                </div>
            </form>
            <Button
                onClick={() => { dispatch(2) }}
                className="mt-4 w-[94%] bg-amber-500 text-black hover:cursor-pointer hover:text-white"
            >CONTINUE</Button>
        </div>
    )
}

export default BasicComponent