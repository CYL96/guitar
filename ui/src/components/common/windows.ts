// @ts-ignore
import router from "../components/router/router";

export function NewWindow(path: string) {
    let url = window.location.protocol + '//' + window.location.host + router.resolve(path).href;
    console.log(url)
    window.open(url, '_blank');
}