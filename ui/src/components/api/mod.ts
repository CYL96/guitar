import {httpPostReqCommonNotify} from "../common/http";

const ApiDownloadGuitarModInfo = '/api/DownloadGuitarModInfo'

export function ApiDownloadGuitarModInfoServer() {
    return httpPostReqCommonNotify(ApiDownloadGuitarModInfo, {})
}

const ApiGetGuitarModInfo = '/api/GetGuitarModInfo'

export function ApiGetGuitarModInfoServer() {
    return httpPostReqCommonNotify(ApiGetGuitarModInfo, {})
}

export interface GetGuitarModInfoResult {
    name: string,
    picList: string[],
}


const ApiGetGuitarModList = '/api/GetGuitarModList'

export function ApiGetGuitarModListServer() {
    return httpPostReqCommonNotify(ApiGetGuitarModList, {})
}

export interface GetGuitarModListResult {
    mod_id: number
    mod_name: string,
    mod_url: string,
}