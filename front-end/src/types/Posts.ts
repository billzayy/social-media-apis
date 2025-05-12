import { Users } from "./Users"

export interface Post { 
    ID: string,
    UserId: string,
    Content: string,
    CreatedAt: string,
}

export interface PostLikes {
    UserId: string,
    PostId: string,
    DateLike: string
}

export interface PostComments { 
    ID: string,
    UserId: string,
    PostId: string,
    Comment: string,
    SentDate: string
}

export interface PostMedia { 
    PostId: string,
    Type: string,
    Url: string
}

export interface PostReq { 
    id: string,
    user: Users,
    content: string,
    createdAt: string,
    likes: number,
    comments: number,
    shares: number,
    media?: PostMedia[]
}
