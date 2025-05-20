import React from "react"

const Steps: React.FC<{ number:number[], selectedNum: number }> = ({number, selectedNum}) => { 
    return (
        <div className="flex items-center justify-center -mt-5">
            {number.map((value, key)=> (
                value != number[number.length - 1] ?
                    <div key={key} className="flex justify-center items-center">
                        <NumberBox number={value} selected={selectedNum == value} />
                        <Line/>
                    </div> :
                    <NumberBox key={key} number={value} selected={selectedNum == value} />
            ))}
        </div>
    )
}

const NumberBox: React.FC<{ number: number, selected: boolean }> = ({ number, selected }) => {
    return (
        <div className={`${selected ? "bg-amber-400 text-black" : ""} border border-amber-400 rounded-full px-3 py-1`}>{number}</div>)
}

const Line = () => { 
    return(<hr className="w-34 h-0.5 mx-0 bg-amber-400 border-0 rounded-sm md:my-10 dark:bg-gray-700"/>)
}

export default Steps;