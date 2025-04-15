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

var interactList: Interacts[] = [
    { data: 15, defaultIcon: LightThumbsUp, hoverIcon: DarkThumbsUp,color: "green"},
    { data: 20, defaultIcon: LightMessage, hoverIcon: DarkMessage,color: "amber"},
    { data:1, defaultIcon: LightShare, hoverIcon: DarkShare,color: "blue"}
]

const UserPost: React.FC = () => { 
    return (
        <div className="w-full bg-white py-4 px-10 shadow-md my-10 border">
            <div id="Header" className="flex items-center">
                <Avatar className='size-14 hover:cursor-pointer mr-5'>
                    <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                    <AvatarFallback>CN</AvatarFallback>
                </Avatar>
                <div>
                    <div className="font-bold">BillZay</div>
                    <div className="text-gray-400 text-sm font-light">Junior Software Engineer at Google</div>
                    <div className="text-gray-400 text-sm font-light">25 Nov at 12:24 PM</div>
                </div>
            </div>
            
            <div id="content" className="my-5">
                Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.
            </div>
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