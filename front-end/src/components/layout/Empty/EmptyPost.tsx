import { Skeleton } from "@/components/ui/skeleton";

var interactList : number[] = [1,2,3]

const EmptyPost: React.FC = () => { 
    return (
			<div className="w-full bg-white py-4 px-10 shadow-md my-10 border">
				<div id="Header" className="flex justify-between">
					<div className="flex items-center">
						<Skeleton className="size-14 hover:cursor-pointer mr-4 rounded-full" />
						<div className="w-[300px]">
							<div className="font-bold">
								<Skeleton className="size-3 h-3 w-full m-1.5" />
							</div>
							<div className="text-gray-400 text-sm font-light">
								<Skeleton className="size-3 h-3 w-full m-1.5" />
							</div>
							<div className="text-gray-400 text-sm font-light">
								<Skeleton className="size-3 h-3 w-full m-1.5" />
							</div>
						</div>
					</div>
					<div className="w-8 my-3">
						<Skeleton className="size-5 h-8 rounded-3xl" />
					</div>
				</div>

				<div id="content" className="my-5 w-full">
					<Skeleton className="size-30 h-20 w-full" />
				</div>

				<div id="interact " className="flex justify-between items-center">
					<div className="flex justify-center items-center">
						{interactList.map((key) =>(
                            <Skeleton key={key} className="size-5 mx-1 w-10"/>
                        ))}
					</div>
					<div className="w-6 my-3 mr-3">
						<Skeleton className="size-5 w-full" />
					</div>
				</div>
			</div>
		);
}

export default EmptyPost;