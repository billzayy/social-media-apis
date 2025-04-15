import Navbar from "@/components/layout/Navbar";
import Post from "@/components/layout/Post";

const HomePage: React.FC = () => { 
    return (
        <div className=" ">
            <div className="bg-white text-black flex justify-center items-center">
                <Navbar />
            </div>
            <Post/>
        </div>
    )
}

export default HomePage;