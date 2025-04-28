import { GetPostAxios, ResponseAPI } from "@/config/axios";
import UserPost from "./UserPost";
import { useEffect, useState } from "react";
import { toast } from "sonner";

const UserPostList: React.FC = () => {
    const [postData, setPostData] = useState<ResponseAPI>()

    useEffect(() => {
        let isCancelled = false;
        async function fetchData() {
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

        fetchData();
        return () => {
            isCancelled = true;
        };
    }, []);

    
    return (
        <div>
            {postData != undefined ? postData?.data.Posts.map((data: any) => (
                <div className="" key={data.PostId}>
                    <UserPost
                        avatar={data.Author.ProfilePicture}
                        content={data.Content}
                        createdAt={data.CreatedAt}
                        likes={data.Likes == undefined ? 0 : data.Likes}
                        comments={data.Comments == undefined? 0: data.Comments}
                        media={data.Media == undefined? undefined : data.Media}
                    />
                </div>
            )) : <div>Hello</div>}
        </div>
    )
}

export default UserPostList;