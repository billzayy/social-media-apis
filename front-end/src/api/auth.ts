import { publicAxios } from "@/config/axios";
import { ResponseAPI } from "@/types/ResponseAPI";
import { RegisterReq } from "@/types/Users";

export const LoginAxios = async (username: string, password: string): Promise<ResponseAPI> => { 
    try {
        const queryParams = `?userName=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`;
        const response = await publicAxios.post(`/auth/login${queryParams}`,{
            headers: {
                'Content-Type': 'application/json',
            },
        });
        
        const resp: ResponseAPI = {
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }

        localStorage.setItem("id", response.data.data.UserId)
        localStorage.setItem("token", resp.data.Token)
        sessionStorage.setItem("cookie", JSON.stringify({ "Token": document.cookie }))
        
        console.log(document.cookie)

        return resp
    } catch (error) {
        console.log('Error logging in:', error);
        return {statusCode: 404, data: error, message:"Failure"}
    }
}

export const RegisterAxios = async (req: RegisterReq | undefined): Promise<ResponseAPI> => { 
    try {
        if (req === undefined) { 
            return {statusCode: 403, data: "Request Error : undefined", message:"Failure"} 
        }
        const response = await publicAxios.post('/auth/register', req, {
            headers: {
               'Content-Type': 'application/json' 
            }
        })

        const resp: ResponseAPI = {
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }
        
        return resp
    } catch (error) {
        return {statusCode: 404, data: error, message:"Failure"} 
    }
}