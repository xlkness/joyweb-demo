export interface LoginFormInt {
    username: string
    password: string
    modules: string[],
}

export class LoginData {
    ruleForm: LoginFormInt={
        username: "",
        password: "",
        modules: [],
    }
}