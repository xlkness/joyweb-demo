import service from "@/requests/api";

interface homeView {
    project: string,
}

interface commandServerData {
    project: string,
    name: string,
    desc: string,
    addr: string,
    envs: string[],
}

interface deleteCommandServer {
    project: string,
    name: string,
}

interface envData {
    project: string,
    name: string,
    desc: string,
    index: number,
}

interface deleteEnv {
    project: string,
    name: string,
}

interface commandlistData {
    project: string,
    name: string,
}

interface execCommandFieldInfo {
    name: string,
    value: string,
}
interface execCommandBaseInfo {
    name: string,
    desc: string,
    fields: execCommandFieldInfo[],
}
interface execCommand {
    project: string,
    command_server_name: string,
    env: string,
    base_info: execCommandBaseInfo,
}

interface deleteDislikeCommand {
    project: string,
    id: string,
}

interface deleteExecHistory {
    project: string,
    name: string,
    index: number,
}

export function gmtool(data: homeView) {
    return service.gmtool({
        url: "/gmtool",
        method: "get",
        params: data,
    })
}

export function queryCmdServerCommandList(data: commandlistData) {
    return service.gmtool({
        url: "/commandlist",
        method: "get",
        params: data,
    })
}

export function execCmdServerCommand(data: execCommand) {
    return service.gmtool({
        url: "/execcommand",
        method: "post",
        data: data,
    })
}

export function addCmdServer(data: commandServerData) {
    return service.gmtool({
        url: "/addcommandserver",
        method: "post",
        data,
    })
}

export function editCmdServer(data: commandServerData) {
    return service.gmtool({
        url: "/editcommandserver",
        method: "post",
        data,
    })
}

export function deleteCmdServer(data: deleteCommandServer) {
    return service.gmtool({
        url: "/deletecommandserver",
        method: "post",
        data,
    })
}

export function addEnv(data: envData) {
    return service.gmtool({
        url: "/addenv",
        method: "post",
        data,
    })
}

export function editEnv(data: envData) {
    return service.gmtool({
        url: "/editenv",
        method: "post",
        data,
    })
}

export function deleteEnv(data: deleteEnv) {
    return service.gmtool({
        url: "/deleteenv",
        method: "post",
        data,
    })
}

export function likeCommand(data: execCommand) {
    return service.gmtool({
        url: "/likecommand",
        method: "post",
        data: data,
    })
}

export function dislikeCommand(data: deleteDislikeCommand) {
    return service.gmtool({
        url: "/dislikecommand",
        method: "post",
        data: data,
    })
}

export  function deleteExecHistory(data: deleteExecHistory) {
    return service.gmtool({
        url: "/deleteexechistory",
        method: "post",
        data: data
    })
}