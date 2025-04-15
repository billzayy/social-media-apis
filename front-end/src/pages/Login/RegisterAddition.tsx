import { Button } from "@/components/ui/button";

const AdditionComponent: React.FC<{ dispatch: React.Dispatch<React.SetStateAction<number>> }> = ({dispatch}) => { 
    return (
        <div>
            <div className="text-left font-bold text-xl mb-5">Additional Information</div>
            <textarea name="" id="" className="border border-gray-500 w-full h-[100px] px-1 resize-none" placeholder="Bio"></textarea>
            <br />
            <input required placeholder="Social Media URL #1" className="border border-black border-b-white text-white my-2.5 w-[100%] py-1 focus:outline-none" />
            <br />
            <input placeholder="Social Media URL #2" className="border border-black border-b-white text-white my-2.5 w-[100%] py-1 focus:outline-none" />
            <br />
            <input placeholder="Social Media URL #3" className="border border-black border-b-white text-white my-2.5 w-[100%] py-1 focus:outline-none" />
            <br />
            <Button
                onClick={() => { dispatch(3) }}
                className="mt-4 w-[100%] bg-amber-500 text-black hover:cursor-pointer hover:text-white"
            >CONTINUE</Button>
        </div>
    )
}

export default AdditionComponent;