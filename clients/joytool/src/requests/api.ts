import axios from 'axios'
import {ElMessageBox} from "element-plus";

// 创建axios实例
const userService=axios.create({
    baseURL: "/user",
    timeout: 10000,
    headers: {
        "Content-type": "application/json;charset=utf-8",
        "Cache-Control": 'no-cache',
    }
})

const gmtoolService=axios.create({
    baseURL: "/gmtoolapi",
    timeout: 10000,
    headers: {
        "Content-type": "application/json;charset=utf-8",
        "Cache-Control": 'no-cache',
    }
})

// axios.post('192.168.1.60:8888/login', {username: "likun", password: "likun1"}).then((res)=>{
//     console.log(res)
// }).catch((err)=>{
//     console.log('login error')
//     console.log(err)
// })

// 请求拦截
userService.interceptors.request.use((config)=>{
    config.headers=config.headers || {}
    if(localStorage.getItem('token')){
        config.headers["x-token"]=localStorage.getItem('token') || ""
    } else {
        console.log("not found local storage token")
    }
    return config
})
// 响应拦截
userService.interceptors.response.use((res)=>{
    const code:number=res.data.code
    if(code!=200) {
        console.log("interceptor err code", res)
        ElMessageBox.alert("请求服务器成功，但是逻辑错误:" + res.data.message, "服务器错误码" + code, {
            type: "warning",
            confirmButtonText: '知道了',
        })
        return Promise.reject(res.data)
    }

    return res.data
}, (err)=>{
    console.log(err)
    const code:number = err.response && err.response.status || -1
    const message: string = err.response && err.response.data.message || err
    ElMessageBox.alert(message, "请求服务器返回http错误码-" + code, {
        type: "error",
        confirmButtonText: '知道了',
    })
    return Promise.reject(err)
})

// 请求拦截
gmtoolService.interceptors.request.use((config)=>{
    config.headers=config.headers || {}
    if(localStorage.getItem('token')){
        config.headers["x-token"]=localStorage.getItem('token') || ""
    } else {
        // console.log("not found local storage token")
    }
    return config
})
// 响应拦截
gmtoolService.interceptors.response.use((res)=>{
    const code:number=res.data.code
    if(code!=200) {
        console.log("interceptor err code", res)
        ElMessageBox.alert("请求服务器成功，但是逻辑错误:" + res.data.message, "服务器错误码" + code, {
            type: "warning",
            confirmButtonText: '知道了',
        })
        return Promise.reject(res.data)
    }

    return res.data
}, (err)=>{
    console.log("err", err)
    const code:number = err.response && err.response.status || -1
    const message: string = err.response && err.response.data.message || err
    ElMessageBox.alert(message, "请求服务器返回http错误码-" + code, {
        type: "error",
        confirmButtonText: '知道了',
    })
    return Promise.reject(err)
})

export default {
    user: userService,
    gmtool: gmtoolService,
}
