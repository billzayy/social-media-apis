import { Button } from "@/components/ui/button";
import { RegisterAddition } from "@/types/Users";
import { useState } from "react";

const AdditionComponent: React.FC<{
    // Dispatch Action for set step need to move
    moveToStep: React.Dispatch<React.SetStateAction<number>>
     // Dispatch Action for set Request Addition Register Data
    addition: React.Dispatch<React.SetStateAction<RegisterAddition | undefined>>
}> = ({ moveToStep, addition }) => { 
    const [bio, setBio] = useState<string>("")
    const [urls, setUrls] = useState<string[]>([])

    const req: RegisterAddition = { // Request Addition Register data
        Bio: bio || null,
        Url: urls
    }

    return (
        <div>
            <div className="text-left font-bold text-xl mb-5">Additional Information</div>
            <textarea
                onChange={(e) => {
                    setBio(e.target.value)
                }}
                name="" id="" className="border border-gray-500 w-full h-[100px] px-1 resize-none" placeholder="Bio"></textarea>
            <br />
            <input
                onChange={(e) => { 
                    setUrls(prevUrls => {
                        const updated = [...prevUrls];
                        updated[0] = e.target.value;
                        return updated;
                      });
                }}
                required placeholder="Social Media URL #1" className="border border-black border-b-white text-white my-2.5 w-[100%] py-1 focus:outline-none" />
            <br />
            <input
                onChange={(e) => { 
                    setUrls(prevUrls => {
                        const updated = [...prevUrls];
                        updated[1] = e.target.value;
                        return updated;
                      });
                }}
                placeholder="Social Media URL #2" className="border border-black border-b-white text-white my-2.5 w-[100%] py-1 focus:outline-none" />
            <br />
            <input
                onChange={(e) => { 
                    setUrls(prevUrls => {
                        const updated = [...prevUrls];
                        updated[2] = e.target.value;
                        return updated;
                      });
                }}
                placeholder="Social Media URL #3" className="border border-black border-b-white text-white my-2.5 w-[100%] py-1 focus:outline-none" />
            <br />
            <Button
                onClick={() => {
                    moveToStep(3);
                    addition(req)
                }}
                className="mt-4 w-[100%] bg-amber-500 text-black hover:cursor-pointer hover:text-white"
            >CONTINUE</Button>
        </div>
    )
}

export default AdditionComponent;