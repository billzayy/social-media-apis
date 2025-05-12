import { IconComponents } from "@/components/ui/Icon";
import { Interacts } from "@/types/interacts";
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"
import {
    faBookmark as LightBookMark,
    faMessage as LightMessage,
    faShareFromSquare as LightShare,
    faThumbsUp as LightThumbsUp
} from "@fortawesome/free-regular-svg-icons";

import {
    faBookmark as DarkBookMark,
    faMessage as DarkMessage,
    faShareFromSquare as DarkShare,
    faThumbsUp as DarkThumbsUp
} from "@fortawesome/free-solid-svg-icons";
import { PostReq } from "@/types/Posts";
import { renderFormattedText } from "@/components/utils/FormattedText";
import { useEffect, useState } from "react";
import { AddLikes, CheckLikes, RemoveLikes } from "@/config/axios";
import { toast } from "sonner";

const UserPost: React.FC<PostReq> = ({ id, user, content, createdAt, likes, comments, shares, media }) => {
    const [likeData, setLikeData] = useState<number>(0)
    const [likeCheck, setLikeCheck] = useState<boolean>(false)
    const [clicked, setClicked] = useState<boolean>(false)

    var interactList: Interacts[] = [
        { name: "Like",data: likes, defaultIcon: LightThumbsUp, hoverIcon: DarkThumbsUp,color: "green"},
        { name: "Comment",data: comments, defaultIcon: LightMessage, hoverIcon: DarkMessage,color: "amber"},
        { name: "Share",data: shares, defaultIcon: LightShare, hoverIcon: DarkShare,color: "blue"}
    ]

    const date = new Date(createdAt)

    useEffect(() => { 
        let isCancelled = false;
        async function fetchCheckLike() {
            const resp = await CheckLikes(user.ID, id);

            if (!isCancelled) {
                if (resp.statusCode != 200) {
                    toast.error(`Error get like data !`, {
                        description: resp.data.message,
                        position: "top-right"
                    })
                } else { 
                    if (resp.data == true) {
                        setClicked(true)
                    } else { 
                        setClicked(false)
                    }
                }
            }
        }

        fetchCheckLike();
        return () => {
            isCancelled = true;
        };
    },[])

    return (
        <div className="w-full bg-white py-4 px-10 shadow-md my-10 border">
            <div id="Header" className="flex items-center">
                <Avatar className='size-14 hover:cursor-pointer mr-5'>
                    <AvatarImage src={user.profilePicture} alt="@shadcn" />
                    <AvatarFallback>{user.fullName[0]}</AvatarFallback>
                </Avatar>
                <div>
                    <div className="hidden">{user.ID}</div>
                    <div className="font-bold">{user.fullName}</div>
                    <div className="text-gray-400 text-sm font-light">Junior Software Engineer at Google</div>
                    <div className="text-gray-400 text-sm font-light">{date.toLocaleDateString()}</div>
                </div>
            </div>
            
            {
                media != undefined ?
                   <div id="content" className="my-5 w-full">
                        {renderFormattedText(content)}
                        <img src={media[0].Url} alt="" className="my-4"/> 
                    </div> :
                    <div id="content" className="my-5 w-full">
                        {renderFormattedText(content)}
                    </div>
            }
            
            <div id="interact " className="flex justify-between items-center">
                <div className="flex justify-center items-center">
                    {interactList.map((data, key) => (
                        <div
                            onClick={() => {
                                if (data.name == "Like" && !clicked) { 
                                    AddLikes(user.ID, id)
                                    setClicked(true)
                                    setLikeData(likes+1)
                                }

                                if (data.name == "Like" && clicked) { 
                                    RemoveLikes(user.ID, id)
                                    setClicked(false)
                                    setLikeData(likes-1)
                                }
                            }}
                            className="mr-5 flex items-center" key={key}>
                            <IconComponents name={data.name} data={data.data} defaultIcon={data.defaultIcon} hoverIcon={data.hoverIcon} color={data.color} margin="mr-2"/>
                            <p className="text-sm">{data.data}</p>
                        </div>
                    ))}
                </div>
                <div>
                    <IconComponents name="Shares" data={likeData} defaultIcon={LightBookMark} hoverIcon={DarkBookMark} color="amber"/>
                </div>
            </div>
        </div>
    )
}

export default UserPost;