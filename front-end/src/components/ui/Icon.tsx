import { IconProp } from "@fortawesome/fontawesome-svg-core"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { useState } from "react"

export const IconComponents: React.FC<{
    hoverIcon?: IconProp,
    defaultIcon: IconProp,
    color?: string,
    margin? : string
}> = ({ hoverIcon, defaultIcon, color, margin }) => {    
    const [hover, setHover] = useState<boolean>(false)

    return (
        <FontAwesomeIcon
            onMouseEnter={() => {setHover(true)}}
            onMouseLeave={() => {setHover(false)}}   
            className={`text-xl ${margin} hover:cursor-pointer text-${color}-400`}
            icon={hover ? hoverIcon || defaultIcon : defaultIcon}>
        </FontAwesomeIcon>
    )
}