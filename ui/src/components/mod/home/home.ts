import router from "../../router/router";


export const PageHome = "/home"

export function GotoHome() {
    router.push(PageHome).then(r => {
    })
}
