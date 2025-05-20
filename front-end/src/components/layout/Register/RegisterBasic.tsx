import { Button } from "@/components/ui/button";
import { RegisterReq } from "@/types/Users";
import React, { useState } from "react";

const BasicComponent: React.FC<{
    // Dispatch Action for set step need to move
    moveToStep: React.Dispatch<React.SetStateAction<number>>
    // Dispatch Action for set Request Basic Register Data
    reqDispatch : React.Dispatch<React.SetStateAction<RegisterReq | undefined>> 
}> = ({ moveToStep, reqDispatch }) => { 
    
    const [userName, setUserName] = useState<string>("")
    const [email, setEmail] = useState<string>("")
    const [firstName, setFirstName] = useState<string>("")
    const [surName, setSurName] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const [city, setCity] = useState<string>("")
    const [age, setAge] = useState<number>(0)

    // Request Basic Register Data
    const req: RegisterReq = { 
        UserName: userName,
        Email: email,
        FirstName: firstName,
        SurName: surName,
        Password: password,
        Location: city,
        Age: age
    }

    return (
        <div className="text-left pl-4">
            <div className="mt-1 mb-3 text-2xl">Basic Information</div>
            <form action="">
                <div>
                    <input onChange={(e) => {
                        setUserName(e.target.value)
                    }} required type="text" placeholder="Login" className="border border-black border-b-white text-white my-2.5 w-[92%] py-1 focus:outline-none" />
                </div>
                <div>
                    <input onChange={(e) => {
                        setEmail(e.target.value)
                    }} required type="email" placeholder="Email" className="border border-black border-b-white text-white my-2.5 w-[92%] py-1 focus:outline-none" />
                </div>
                <div className="flex">
                    <div>
                        <input onChange={(e) => {
                            setFirstName(e.target.value)
                        }} required type="text" placeholder="Name" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                    <div>
                        <input onChange={(e) => {
                            setSurName(e.target.value)
                        }} required type="text" placeholder="Surname" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                </div>
                <div>
                    <input onChange={(e) => {
                            setPassword(e.target.value)
                    }} required type="password" placeholder="Password" className="border border-black border-b-white text-white my-2.5 w-[92%] py-1 focus:outline-none" />
                </div>
                <div className="flex">
                    <div>
                        <input onChange={(e) => {
                            setCity(e.target.value)
                        }} required type="text" placeholder="City" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                    <div>
                        <input onChange={() => {
                            setAge(18)
                        }} required type="text" placeholder="Age" className="border border-black border-b-white text-white my-2.5 w-[85%] py-1 focus:outline-none" />
                    </div>
                </div>
            </form>
            <Button
                onClick={() => {
                    moveToStep(2);
                    reqDispatch(req)
                }}
                className="mt-4 w-[94%] bg-amber-500 text-black hover:cursor-pointer hover:text-white"
            >CONTINUE</Button>
        </div>
    )
}

export default BasicComponent