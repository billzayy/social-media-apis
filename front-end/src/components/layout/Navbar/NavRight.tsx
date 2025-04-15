import { faSearch, faBell as DarkBell, faBookmark as DarkBookMark } from '@fortawesome/free-solid-svg-icons'
import { faBell as LightBell, faBookmark as LightBookMark } from "@fortawesome/free-regular-svg-icons"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Link } from 'react-router-dom'
import { IconComponents } from '@/components/ui/Icon'

const NavRight: React.FC = () => {
    return (
        <div className="flex justify-center items-center w-[20%]">
            <IconComponents defaultIcon={faSearch} margin='mr-8' />
            <IconComponents hoverIcon={DarkBookMark} defaultIcon={LightBookMark} margin='mr-8'/>
            <IconComponents hoverIcon={DarkBell} defaultIcon={LightBell} margin='mr-8'/>
            <DropdownMenu>
                <DropdownMenuTrigger>
                    <Avatar className='size-12 hover:cursor-pointer'>
                        <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                        <AvatarFallback>CN</AvatarFallback>
                    </Avatar>
                </DropdownMenuTrigger>
                <DropdownMenuContent>
                    <DropdownMenuLabel>My Account</DropdownMenuLabel>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem className='hover:cursor-pointer'>Profile</DropdownMenuItem>
                    <DropdownMenuItem className='hover:cursor-pointer'>Setting</DropdownMenuItem>
                    <DropdownMenuItem
                        onClick={() => {
                            sessionStorage.removeItem("local")
                        }}
                        className='hover:cursor-pointer'>
                        <Link to={"/login"}>Log Out</Link>
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>

        </div>        
    )
}

export default NavRight;