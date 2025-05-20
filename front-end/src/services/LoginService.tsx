import { LoginAxios } from "@/api/auth";
import { NavigateFunction } from "react-router-dom";
import { toast } from "sonner";

export async function LoginService(userName:string, password: string, navigate: NavigateFunction)  {
	const result = await LoginAxios(userName, password);

	if (result?.message == "Failure") {
		toast.error(`Login failed !`, {
			description: result.data.message,
			position: "top-right",
		});
	} else {
		toast.success("Login successful!", {
			position: "top-right",
		});

		navigate("/"); // Navigate to the dashboard or target page
	}
}