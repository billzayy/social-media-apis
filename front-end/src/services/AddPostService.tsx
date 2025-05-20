import { AddPostAxios } from "@/api/post";
import { toast } from "sonner";

export async function AddPostService(userId: string | null, content: string) {
	if (!content) {
		toast.error("Failed to create post", {
			description: "Content is empty!",
			position: "top-right",
		});
		return;
	}

	if (!userId) {
		toast.error("Failed to create post!", {
			description: "Please login to create post!",
			position: "top-right",
		});
		return;
	}

	try {
		const resp = await AddPostAxios(userId, content);

		if (resp.statusCode === 201) {
			toast.success("Created", {
				description: "Created",
				position: "top-right",
			});
		}
	} catch (error) {
		toast.error("Error", {
			description: `${error}`,
			position: "top-right",
		});
	}
}
