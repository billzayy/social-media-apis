import CreatePost from "./Post/CreatePost";
import UserPostList from "./Post/UserPostList";

const Post: React.FC = () => { 
    return (
        <div className="mx-[11%] my-[3.5%] h-screen flex justify-between mt-26">
            <div id="left" className={`w-2/3 mr-6`}>
                <CreatePost/>
                <UserPostList/>
            </div>
            <div id="right" className={`border border-green-400 w-1/3`}>
                <div>Content1</div>
                <div>Content1</div>
                <div>Content1</div>
                <div>Content1</div>
            </div>
        </div>
    )
}

export default Post;