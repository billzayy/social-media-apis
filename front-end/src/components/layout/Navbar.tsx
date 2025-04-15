import NavMid from "./Navbar/NavMid";
import NavRight from "./Navbar/NavRight";

const Navbar: React.FC = () => { 
    return (
        <div className="fixed top-0 right-0 left-0 w-full flex justify-between items-center bg-white py-4 px-10 text-lg shadow-md z-10">
            <div>Logo</div>
            <NavMid/>
            <NavRight/> 
        </div>
    )
}

export default Navbar;