import router from "../../router/router";


export const PageDown="/down"

export function GotoDown(){
    router.push(PageDown).then(r => {})
}