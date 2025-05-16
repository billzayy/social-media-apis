export interface Users { 
    UserId: string,
    FullName: string,
    ProfilePicture: string
}

export interface RegisterReq{
    UserName : string,
    Email : string,
    FirstName : string,
    SurName : string,
    Password : string,
    Location: string,
    Age: number,
    Addition?: RegisterAddition | undefined
}

export interface RegisterAddition { 
    Bio?: string | null,
    Url?: string[]
}