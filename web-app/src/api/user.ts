import api from "../utils/api.ts";
import type {UserType} from "../types/user.ts";

export const fetchUsers = () => new Promise<UserType[]>((resolve, reject) =>
    api.get(`/users`).then((res) => resolve(res.data.data)).catch(
        err => reject(err)
    )
);

export const searchUsers = (query: string) => new Promise<UserType[]>((resolve, reject) =>
    api.get(`/users/search?search_term=${query}`).then((res) => resolve(res.data.data)).catch(
        err => reject(err)
    )
);

export const fetchCurrentUser = () => new Promise<UserType>((resolve, reject) =>
    api.get(`users/me`).then((res) => resolve(res.data.data)).catch(
        err => reject(err)
    )
);
