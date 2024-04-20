import {ElMessage, MessageHandler} from "element-plus";

var nowMessage:MessageHandler = null
export const MessageErr = (err: string) => {
    if (nowMessage != null){
        nowMessage.close()
    }
    nowMessage = ElMessage({
        showClose: true,
        message: err,
        type: 'error',
    })
}