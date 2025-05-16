import { GetPostAxios, ResponseAPI } from "@/config/axios";
import UserPost from "./UserPost";
import { useEffect, useState } from "react";
import { toast } from "sonner";
import { PostData } from "@/types/Posts";

const UserPostList: React.FC = () => {
    const [postData, setPostData] = useState<ResponseAPI | undefined>()

    useEffect(() => {
        let isCancelled = false;
        async function fetchPostData() {
            const resp = await GetPostAxios();

            if (!isCancelled) {
                if (resp.statusCode != 200) {
                    toast.error(`Failed to get data !`, {
                        description: "result.data.message",
                        position: "top-right"
                    })
                } else { 
                    setPostData(resp)
                }
            }
        }


        fetchPostData();

        return () => {
            isCancelled = true;
        };
    }, []);

    
    return (
        <div>
            {postData !== undefined ? postData.data.PostList.map((v: PostData) => (
                <div
                    className="" key={v.PostId}>
                    <UserPost
                        id={v.PostId}
                        user={v.Author}
                        content={v.Content}
                        createdAt={v.CreatedAt}
                        likes={v.Likes == undefined ? 0 : v.Likes}
                        comments={v.Comments == undefined ? 0 : v.Comments}
                        shares={v.Shares == undefined ? 0 : v.Shares}
                        media={v.Media == undefined? undefined : v.Media}
                    />
                </div>
            )) : <div className="mt-10">Loading ...</div>}
        </div>
    )
}

export default UserPostList;