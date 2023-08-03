import service from "@/requests/api";

interface userBaseInfoData {
    username: string,
    password: string,
}

interface groupListData {
    name: string,
    is_admin: boolean,
    group: string,
}

interface createUserData {
    base_info: userBaseInfoData,
    group_list: groupListData[],
}

interface deleteUserData {
    username: string,
}

interface loginData {
    base_info: userBaseInfoData,
}

export function createUser(data: createUserData){
    return service.user({
        url: "/createuser",
        method: "post",
        data
    })
}

export function editUser(data: createUserData){
    return service.user({
        url: "/edituser",
        method: "post",
        data
    })
}

export function deleteUser(data: deleteUserData){
    return service.user({
        url: "/deleteuser",
        method: "post",
        data
    })
}

export function listUsers(data){
    return service.user({
        url: "/listusers",
        method: "get",
        params: data
    })
}

export function login(data: loginData){
    return service.user({
        url: "/login",
        method: "post",
        data
    })
}

export function logout() {
    return service.user({
        url: "/logout",
        method: "get",
    })
}