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
    // id: string,
    avatar: string,
    content: string,
    createdAt: string,
    likes: number,
    comments: number,
    media?: PostMedia[]
}
