

let nowId = new Date().getTime();

export function GetNewId(): number {
    let id = new Date().getTime();
    if (id == nowId){
        id++
    }
    nowId = id
    return id;

}