import { GetPostAxios } from "@/config/axios";
import { NavigateFunction } from "react-router-dom";
import { toast } from "sonner";

export async function PostService(navigate: NavigateFunction) {
    const result = await GetPostAxios();

    if (result?.message == "Failure") { 
      toast.error(`Login failed !`, {
        description: result.data.message,
        position: "top-right"
      })
    }else {
        toast.success("Login successful!", {
          position: "top-right"
        });

        navigate("/"); // Navigate to the dashboard or target page
    }
}
