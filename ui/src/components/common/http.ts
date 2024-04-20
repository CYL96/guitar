import axios from "axios";
import {runConfig} from "./config";
import {ElMessage} from "element-plus";
import {MessageErr} from "./msg";

export interface ServerResponse {
    state: number,
    msg: string,
    data: any
}

export async function httpPostReq(api: string, para: any) {
    return httpPostReqByHost(runConfig.server, api, para)
}


export async function httpPostReqCommonNotify(api: string, para: any) {
    return httpPostReqByHost(runConfig.server, api, para).then(
        res => {
            if (res.state != 0) {
                MessageErr(res.msg)
            }
            return res
        }
    )
}


export async function httpPostReqByHost(host: string, api: string, para: any) {
    const data = JSON.stringify(para);

    const config = {
        method: 'post',
        url: host + api,
        headers: {
            'Content-Type': 'application/json'
        },
        data: data
    };
    let resp: ServerResponse = {
        state: -1,
        msg: "失败",
        data: null,
    }
    try {
        await axios(config)
            .then(function (response) {
                resp.state = response.data.state;
                resp.msg = response.data.msg;
                resp.data = response.data.data;
            })
            .catch(function (error) {
                resp.state = 1;
                resp.msg = "无法访问目标主机";
                resp.data = {};
            });
    } catch (err) {
        resp.state = 1;
        resp.msg = "无法访问目标主机";
        resp.data = {};

    }
    return resp
}