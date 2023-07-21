import service from "@/requests/api";

interface loginData {
    username: string,
    password: string,
    modules: string[],
}

export function login(data: loginData){
    return service.user({
        url: "/login",
        method: "post",
        data
    })
}
