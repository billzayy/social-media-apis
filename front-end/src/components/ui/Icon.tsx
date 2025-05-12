import { IconProp } from "@fortawesome/fontawesome-svg-core"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { useState } from "react"

export const IconComponents: React.FC<{
    name: string,
    data : number,
    hoverIcon?: IconProp,
    defaultIcon: IconProp,
    color?: string,
    margin? : string
}> = ({name, data, hoverIcon, defaultIcon, color, margin }) => {    
    const [hover, setHover] = useState<boolean>(false)

    return (
        <FontAwesomeIcon
            onMouseEnter={() => { setHover(true) }}
            onMouseLeave={() => { setHover(false) }}
            className={`text-xl ${margin} hover:cursor-pointer text-${color}-400`}
            icon={
                name != "Like" && data == 0 ?
                    hover ?
                        hoverIcon || defaultIcon
                        :
                        defaultIcon
                    :
                    name == "Like" && data != 0 ?
                        hoverIcon || defaultIcon
                        :
                    defaultIcon}>
        </FontAwesomeIcon>
    )
}