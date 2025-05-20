import { privateAxios, publicAxios } from "@/config/axios";
import { ResponseAPI } from "@/types/ResponseAPI";

export const CheckLikes = async (userId: string, postId: string): Promise<ResponseAPI> => { 
    try {
        var token = localStorage.getItem("token");

				if (token === "") {
					return {
						statusCode: 404,
						data: "Not found Token",
						message: "Failure",
					};
				}
        const response = await publicAxios.post('/api/v1/interact/check-like', 
            {
                UserId: userId,
                PostId: postId
            }, 
            {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization' : `bearer ${token}`
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
        var token = localStorage.getItem("token")

        if (token === "") { 
            return { statusCode: 404, data: "Not found Token", message: "Failure" }; 
        }

        const response = await privateAxios.post('/api/v1/interact/add-like', 
            {
                UserId: userId,
                PostId: postId
            }, 
            {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${token}`
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
        const response = await privateAxios.delete('/api/v1/interact/delete-like', {
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