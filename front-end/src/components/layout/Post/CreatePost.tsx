import { faCalendar, faCirclePlay } from "@fortawesome/free-regular-svg-icons";
import { faImage, faPen, IconDefinition } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"

interface Options{
    icon: IconDefinition,
    name: string
}

var OptionsList: Options[] = [
    { icon: faImage, name: "Photo" },
    { icon: faCirclePlay, name: "Video" },
    { icon: faCalendar, name: "Event" },
    {icon: faPen, name:"Write an Article"}
]

const CreatePost : React.FC = () => {
    return (
        <div className="bg-white shadow-xl px-5 py-4 border">
            <div id="up" className="flex items-center py-8">
                <Avatar className='size-13 hover:cursor-pointer'>
                    <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                    <AvatarFallback>CN</AvatarFallback>
                </Avatar>
                <input type="text" className="ml-2 focus:outline-none" placeholder="Post something?" />
            </div> 
            <div className="my-2 border-t border-gray-300"></div>
            <div id="down" className="flex items-center justify-between px-2">
                {OptionsList.map((data, key) => (
                    <div key={key} className="flex items-center hover:cursor-pointer mt-3 mb-2">
                        <FontAwesomeIcon icon={data.icon} className="mr-2.5 text-blue-400" />
                        <p className="">{data.name}</p>
                   </div> 
                ))}
            </div>
        </div>
    )
}

export default CreatePost;