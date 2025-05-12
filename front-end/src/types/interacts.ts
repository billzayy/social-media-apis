import { IconDefinition } from "@fortawesome/fontawesome-svg-core";

export interface Interacts { 
    name: string,
    data: number,
    defaultIcon: IconDefinition,
    hoverIcon: IconDefinition,
    color? : string
}