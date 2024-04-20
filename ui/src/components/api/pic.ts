import {httpPostReqCommonNotify} from "../common/http.ts";


// 请求列表
const ApiGetGuitarPicList = '/api/GetGuitarPicList'

export function ApiGetGuitarPicListServer(req: GetGuitarPicListReq) {
    return httpPostReqCommonNotify(ApiGetGuitarPicList, req)
}

export interface GetGuitarPicListReq {
    name: string
    search: string
}

export interface GetGuitarPicListResult {
    id: number
    class_name: string
    picList: string[]
}

export function CopyGuitarPicListResult(item: GetGuitarPicListResult): GetGuitarPicListResult {
    let result = NewGetGuitarPicListResult()
    result.id = item.id
    result.class_name = item.class_name
    item.picList.forEach((value) => {
            result.picList.push(value)
        }
    )
    return result
}

export function NewGetGuitarPicListResult(): GetGuitarPicListResult {
    return {
        id: 0,
        class_name: "",
        picList: []
    }
}

// 删除图片
const ApiDeleteGuitarClass = '/api/DeleteGuitarClass'

export function ApiDeleteGuitarClassServer(name: string) {
    return httpPostReqCommonNotify(ApiDeleteGuitarClass, {name: name})
}


// 修改名称
const ApiRenameGuitarClass = '/api/RenameGuitarClass'

export function ApiRenameGuitarClassServer(req :ApiRenameGuitarClassServerReq) {
    return httpPostReqCommonNotify(ApiRenameGuitarClass, req)
}

export interface ApiRenameGuitarClassServerReq {
    name: string
    oldName: string
}
export function NewApiRenameGuitarClassServerReq(): ApiRenameGuitarClassServerReq {
    return {
        name: "",
        oldName: ""
    }
}