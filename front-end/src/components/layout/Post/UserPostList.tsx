import { GetPostAxios, ResponseAPI } from "@/config/axios";
import UserPost from "./UserPost";
import { useEffect, useState } from "react";
import { toast } from "sonner";

const UserPostList: React.FC = () => {
    const [postData, setPostData] = useState<ResponseAPI | undefined>()

    useEffect(() => {
        let isCancelled = false;
        async function fetchPostData() {
            const resp = await GetPostAxios();

            if (!isCancelled) {
                if (resp.statusCode != 200) {
                    toast.error(`Error get post data !`, {
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
            {postData != undefined ? postData.data.map((v: any) => (
                <div
                    className="" key={v.postId}>
                    <UserPost
                        id={v.postId}
                        user={v.userId}
                        content={v.content}
                        createdAt={v.createdAt}
                        likes={v.likes == undefined ? 0 : v.likes}
                        comments={v.comments == undefined ? 0 : v.comments}
                        shares={v.shares == undefined ? 0 : v.shares}
                        media={v.media == undefined? undefined : v.media}
                    />
                </div>
            )) : <div className="mt-10">Loading ...</div>}
        </div>
    )
}

export default UserPostList;