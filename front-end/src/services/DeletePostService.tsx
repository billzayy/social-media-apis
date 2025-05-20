import { RemovePostAxios } from "@/api/post";
import { toast } from "sonner";

export async function DeletePostService(userId:string, postId:string) {
	if (userId === "" || userId === null) {
		toast.error("Failed to delete post", {
			description: "userId is empty!",
			position: "top-right",
		});
	}

	if (postId === "" || postId === null) {
		toast.error("Failed to delete post", {
			description: "postId is empty!",
			position: "top-right",
		});
	}

	try {
		const resp = await RemovePostAxios(userId, postId);

		if (resp.statusCode === 200) {
			toast.success("Deleted!", {
				description: "Delete post successful!",
				position: "top-right",
			});
			location.reload();
		} else {
			toast.error("Failed to delete post", {
				description: resp.data,
				position: "top-right",
			});
		}
	} catch (error) {
		toast.error("Failed to delete post", {
			description: `${error}`,
			position: "top-right",
		});
	}
}