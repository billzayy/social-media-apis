import { Button } from "@/components/ui/button"
import { RegisterAxios } from "@/api/auth"
import { RegisterAddition, RegisterReq } from "@/types/Users"
import { useEffect, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { toast } from "sonner"

const RegisterFinal: React.FC<{
    // Dispatch Action for set Request Basic Register Data
    basicReq: RegisterReq | undefined
    // Dispatch Action for set Request Addition Register Data
    addition: RegisterAddition | undefined
}> = ({ basicReq, addition }) => { 
    const req: RegisterReq | undefined = basicReq
    ? { ...basicReq, Addition: addition ?? undefined }
        : undefined;
    
    const [registerStatus, setRegisterStatus] = useState<boolean>(false)

    const navigate = useNavigate()

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

    useEffect(() => {
        if (registerStatus) {
            navigate("/login")
        }
    }, [registerStatus, navigate])

    return (
        <div>
            <div className="font-bold text-xl my-5">All set !</div>
            <Button
                onClick={() => {
                    fetchRegister()
                }}
                asChild className="mt-4 w-[100%] bg-amber-500 text-black hover:cursor-pointer hover:text-white">
                { 
                    registerStatus ?
                        <Link to={"/login"}>Go to Login</Link> : <div>Go to Login</div>
                }
            </Button>
        </div>
    )
}

export default RegisterFinal