import api from "../utils/api.ts";
import type {GroupType} from "../types/Groups.ts";

export const fetchGroups = () => new Promise<GroupType[]>((resolve, reject) =>
    api.get(`/groups`).then((res) => resolve(res.data.data)).catch(
        err => reject(err)
    )
);

export const fetchGroup = (id: number) => new Promise<GroupType>((resolve, reject) =>
    api.get(`/groups/${id}`).then((res) => resolve(res.data.data)).catch((err) => reject(err))
)

export const createGroup = (group: GroupType) => new Promise<GroupType>((resolve, reject) =>
    api.post(`/groups`, {...group}).then((res) => resolve(res.data.data)).catch((err) => reject(err))
)

export const updateGroup = (group: GroupType) => new Promise<GroupType>((resolve, reject) =>
    api.patch(`/groups/${group.ID}`, group).then((res) => resolve(res.data.data)).catch((err) => reject(err))
)

export const expenseRepays = (groupId: number) => new Promise<{From: string, To: string, Amount: number}[]>((resolve, reject) =>
    api.get(`/groups/${groupId}/repays`).then((res) => resolve(res.data.data)).catch((err) => reject(err))
)

export const addUserToGroup = (groupId: number, userIds: number[]) => new Promise((resolve, reject) => {
    api.post(`/groups/${groupId}/add_users`, {user_ids: userIds}).then((res) => {
        resolve(res.data)
    }).catch((err) => {
        reject(err)
    })
})