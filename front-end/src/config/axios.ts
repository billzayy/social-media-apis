import axios from "axios";

var url: string = "http://localhost:3000/auth"

interface ResponseAPI { 
    statusCode: number,
    data: any,
    message: string
}

export const LoginAxios = async (username: string, password: string): Promise<ResponseAPI> => { 
    try {
        const queryParams = `?userName=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`;
        const response = await axios.post(`${url}/login${queryParams}`,{
            headers: {
                'Content-Type': 'application/json',
            },
        });
        
        const resp: ResponseAPI = {
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }

        sessionStorage.setItem("local", JSON.stringify({ "id": response.data.data.ID, "Token": response.data.data.Token }))

        return resp
    } catch (error) {
        console.log('Error logging in:', error);
        return {statusCode: 404, data: error, message:"Failure"}
    }
}