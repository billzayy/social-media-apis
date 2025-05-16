import { Button } from "@/components/ui/button"
import { RegisterAxios } from "@/config/axios"
import { RegisterAddition, RegisterReq } from "@/types/Users"
import { useState } from "react"
import { Link } from "react-router-dom"
import { toast } from "sonner"

const RegisterFinal: React.FC<{
    basicReq: RegisterReq | undefined
    addition: RegisterAddition | undefined
}> = ({ basicReq, addition }) => { 
    const req: RegisterReq | undefined = basicReq
    ? { ...basicReq, Addition: addition ?? undefined }
        : undefined;
    
    const [registerStatus, setRegisterStatus] = useState<boolean>(false)

    async function fetchRegister() { 
        const resp = await RegisterAxios(req)

        if (resp.statusCode != 201) {
            toast.error(`Failed to register !`, {
                description: resp.data.message,
                position: "top-right"
            })
        } else { 
            setRegisterStatus(true)
        }
    }
    return (
        <div>
            <div className="font-bold text-xl my-5">All set !</div>
            <div className="mb-3">Welcome to the Social Media ! Have fun </div>
            <Button
                onClick={() => {
                    fetchRegister()
                }}
                asChild className="mt-4 w-[100%] bg-amber-500 text-black hover:cursor-pointer hover:text-white">
                { 
                    registerStatus ?
                        <Link to={"/"}>Go to Homepage</Link> : <div></div>
                }
            </Button>
        </div>
    )
}

export default RegisterFinal