const menuList = ["Home", "Network", "Events"]

const NavMid: React.FC = () => {
    return (
        <div className="flex justify-around w-[20%] ml-[14%]">
            {menuList.map((data, key) => (
                <span key={key} className="hover:cursor-pointer hover:underline">{data}</span>
            ))}
        </div>
    )
}

export default NavMid;