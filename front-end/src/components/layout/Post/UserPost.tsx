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

const UserPost: React.FC<PostReq> = ({avatar, content, createdAt, likes, comments, media}) => { 
    var interactList: Interacts[] = [
        { data: likes, defaultIcon: LightThumbsUp, hoverIcon: DarkThumbsUp,color: "green"},
        { data: comments, defaultIcon: LightMessage, hoverIcon: DarkMessage,color: "amber"},
        { data:1, defaultIcon: LightShare, hoverIcon: DarkShare,color: "blue"}
    ]

    const date = new Date(createdAt)

    return (
        <div className="w-full bg-white py-4 px-10 shadow-md my-10 border">
            <div id="Header" className="flex items-center">
                <Avatar className='size-14 hover:cursor-pointer mr-5'>
                    <AvatarImage src={avatar} alt="@shadcn" />
                    <AvatarFallback>CN</AvatarFallback>
                </Avatar>
                <div>
                    <div className="font-bold">BillZay</div>
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
                        <div className="mr-5 flex items-center" key={key}>
                            <IconComponents defaultIcon={data.defaultIcon} hoverIcon={data.hoverIcon} color={data.color} margin="mr-2"/>
                            <p className="text-sm">{data.data}</p>
                        </div>
                    ))}
                </div>
                <div>
                    <IconComponents defaultIcon={LightBookMark} hoverIcon={DarkBookMark} color="amber"/>
                </div>
            </div>
        </div>
    )
}

export default UserPost;