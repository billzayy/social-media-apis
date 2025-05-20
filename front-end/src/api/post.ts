import { privateAxios, publicAxios } from "@/config/axios";
import { ResponseAPI } from "@/types/ResponseAPI";

export const GetPostAxios = async (): Promise<ResponseAPI> => { 
    try {
        const response = await publicAxios.get('/api/v1/post/get-post', {
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
        return {statusCode: 404, data: err, message: "Failure"}
    }
}

export const AddPostAxios = async (userId: string, content: string): Promise<ResponseAPI> => { 
    try {
        const response = await privateAxios.post("/api/v1/post/add-post",
          {
            userId: userId,
            content: content
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );

        const resp: ResponseAPI = {
            statusCode: response.data.statusCode,
            data: response.data.data,
            message: response.data.message
        }

        return resp

    } catch (err) {
        return { statusCode: 500, data: err, message: "Failure" };
    }
}