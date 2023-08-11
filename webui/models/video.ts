import { IUser } from "./user";

export interface IVideo {
    id: number;
    title: string;
    description: string;
    like: number;
    dislike: number;
    url: string;
    videoId: string;
    user: IUser;
}