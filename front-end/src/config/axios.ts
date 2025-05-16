import { RegisterReq } from './../types/Users';
import axios from "axios";

var url: string = "http://localhost:3000"

export interface ResponseAPI { 
    statusCode: number,
    data: any,
    message: string
}

export const LoginAxios = async (username: string, password: string): Promise<ResponseAPI> => { 
    try {
        const queryParams = `?userName=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`;
        const response = await axios.post(`${url}/auth/login${queryParams}`,{
            headers: {
                'Content-Type': 'application/json',
            },
        });
        
        const resp: ResponseAPI = {
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }

        localStorage.setItem("user", JSON.stringify({"userId" : response.data.data.UserId}))
        sessionStorage.setItem("local", JSON.stringify({"Token": response.data.data.Token }))

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
        const response = await axios.post(`${url}/auth/register`, req, {
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

export const GetPostAxios = async (): Promise<ResponseAPI> => { 
    try {
        const response = await axios.get(`${url}/api/v1/post/get-post`, {
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

    } catch (err) {
        console.log("Error get data in:", err);
        return {statusCode: 404, data: err, message: "Failure"}
    }
}

export const CheckLikes = async (userId: string, postId: string): Promise<ResponseAPI> => { 
    try {
        const response = await axios.post(`${url}/api/v1/interact/check-like`, 
            {
                UserId: userId,
                PostId: postId
            }, 
            {
                headers: {
                    'Content-Type': 'application/json'
                }
            }
        );

        const resp: ResponseAPI = { 
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }

        return resp
    } catch (err) {
        return {statusCode: 404, data: err, message: "Failure"} 
    } 
} 

export const AddLikes = async (userId:string, postId: string): Promise<ResponseAPI> => { 
    try {
        const response = await axios.post(`${url}/api/v1/interact/add-like`, 
            {
                UserId: userId,
                PostId: postId
            }, 
            {
                headers: {
                    'Content-Type': 'application/json'
                }
            }
        );

        const resp: ResponseAPI = { 
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }

        return resp
    } catch (err) {
        console.log("Error Add Like in:", err);
        return {statusCode: 500, data: err, message: "Failure"} 
    }
}

export const RemoveLikes = async (userId: string, postId: string): Promise<ResponseAPI> => { 
    try {
        const response = await axios.delete(`${url}/api/v1/interact/delete-like`, {
            headers: {
                'Content-Type': 'application/json'
            },
            data: {
                userId: userId,
                postId: postId
            }
        });

        const resp: ResponseAPI = { 
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        };

        return resp;
    } catch (err) {
        console.log("Error in RemoveLikes:", err);
        return { statusCode: 500, data: err, message: "Failure" }; 
    }
};