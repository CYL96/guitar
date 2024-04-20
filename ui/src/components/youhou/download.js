// ==UserScript==
// @name         吉他谱下载
// @namespace    http://tampermonkey.net/
// @version      2024-04-19
// @description  吉他谱下载
// @author       cyl96
// @match        *://*.fox4.cn/*
// @match        *://*.gtpso.com/*
// @match        *://*.cwguitar.cn/*
// @icon         data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==
// @grant        GM_xmlhttpRequest
// ==/UserScript==

let parseInfo = {
    dlUrl: window.location.origin,
    nameInput: '',
    picList: [],
    ServerUrl: 'http://127.0.0.1:55002'
}
var newContainer = document.createElement('div');

// 设置容器的样式，将其放置在页面顶部
newContainer.style.height = '10vh'
document.body.insertBefore(newContainer, document.body.firstChild);

function httpPostReqByHost(host, api, para, onOk) {
    const data = JSON.stringify(para);
    const config = {
        method: 'post',
        url: host + api,
        headers: {
            'Content-Type': 'application/json'
        },
        onload: function (response) {
            onOk(JSON.parse(response.responseText))
        },
        onerror: function (error) {
            alert(error)
        },
        data: data
    };

    return GM.xmlHttpRequest(config);
}


const url = window.location.href;
let div = document.createElement('div');
div.style.position = 'fixed';
div.style.top = '0';
div.style.left = '0';
div.style.padding = '5px';
div.style.width = '100%';
div.style.maxHeight = '5vh';
div.style.display = 'flex'
div.style.flexWrap = 'warp'
div.style.alignItems = 'center'
div.style.justifyContent = 'center'
div.style.background = '#2D9FC2FF'
div.style.border = '1px solid'


const myInput = document.createElement('input');
myInput.style.width = '200px'
myInput.style.height = '30px'
myInput.style.verticalAlign = 'middle'
myInput.value = parseInfo.ServerUrl
myInput.addEventListener('input', function () {
    parseInfo.ServerUrl = myInput.value
});

const nameInput = document.createElement('input')
nameInput.style.width = '500px'
nameInput.style.height = '30px'
nameInput.style.marginLeft = '5px'
nameInput.style.verticalAlign = 'middle'
nameInput.value = parseInfo.nameInput
nameInput.addEventListener('input', function () {
    parseInfo.nameInput = nameInput.value
});


let button = document.createElement("button")
button.style.width = '100px'
button.style.marginLeft = '5px'
button.innerText = '解析'
button.style.height = '30px'
button.onclick = function () {
    ParseNowUrlInfo()
}

let dlbutton = document.createElement("button")
dlbutton.style.width = '100px'
dlbutton.style.marginLeft = '5px'
dlbutton.innerText = '下载'
dlbutton.style.height = '30px'
dlbutton.onclick = function () {
    if (parseInfo.nameInput === '' || parseInfo.picList.length === 0) {
        alert("未识别到相关图片信息")
        return
    }
    dlbutton.innerText = '下载中...'
    dlbutton.disabled = true
    let req = {
        name: parseInfo.nameInput,
        picList: parseInfo.picList
    }
    httpPostReqByHost(
        parseInfo.ServerUrl,
        '/api/DownloadGuitar',
        req,
        function (resp) {
            if (resp.state !== 0) {
                alert(resp.msg)
            } else {
                alert("下载成功")
            }
        })
    dlbutton.innerText = '下载'
    dlbutton.disabled = false

}

let picLabel = document.createElement("label")
picLabel.style.width = '100px'
picLabel.style.height = '30px'
picLabel.style.marginLeft = '5px'
picLabel.style.verticalAlign = 'middle'
picLabel.style.textAlign = 'center'

div.appendChild(myInput)
div.appendChild(button)
div.appendChild(dlbutton)
div.appendChild(nameInput)
div.appendChild(picLabel)

document.body.appendChild(div)

const GetFox4Info = function () {
    parseInfo.nameInput = ''
    parseInfo.picList = []
    picLabel.innerText = ''
    //  获取标题
    let title = document.getElementsByClassName("post-title")
    if (title) {
        parseInfo.nameInput = title[0].innerText
        nameInput.value = parseInfo.nameInput
    }
    // 获取图片
    let content = document.getElementsByClassName("post-content")
    if (content) {
        for (let i = 0; i < content[0].children.length; i++) {
            let item = content[0].children[i]
            parseInfo.picList.push(parseInfo.dlUrl + item.getAttribute("src"))
        }
        picLabel.innerText = '已识别：' + parseInfo.picList.length + " 张"
    }
}


const GetgtpsoInfo = function () {

    //  获取标题
    let title = document.getElementsByClassName("text-center")
    if (title) {
        parseInfo.nameInput = title[0].innerText
        nameInput.value = parseInfo.nameInput
    }
    // 获取图片
    let content = document.getElementsByClassName("img-fluid")
    if (content) {
        for (let i = 0; i < content.length; i++) {
            parseInfo.picList.push(content[i].getAttribute("src"))
        }
        picLabel.innerText = '已识别：' + parseInfo.picList.length + " 张"
    }
}

const GetcwguitarInfo = function () {

    //  获取标题
    let title = document.getElementsByClassName("pt_10 text-ellipsis")
    if (title) {
        parseInfo.nameInput = title[0].innerText
        nameInput.value = parseInfo.nameInput
    }
    // 获取图片
    let content = document.getElementsByClassName(" thumbnail mb_0")
    if (content) {
        for (let i = 0; i < content[1].children.length; i++) {
            parseInfo.picList.push(content[1].children[i].getAttribute("src"))
        }
        picLabel.innerText = '已识别：' + parseInfo.picList.length + " 张"
    }
}


const ParseNowUrlInfo = function () {
    let url = window.location.href;

    parseInfo.nameInput = ''
    parseInfo.picList = []
    picLabel.innerText = ''

    switch (true) {
        case url.indexOf("fox4.cn") !== -1:
            GetFox4Info()
            break;
        case url.indexOf("gtpso.com") !== -1:
            GetgtpsoInfo()
            break;
        case url.indexOf("cwguitar.cn") !== -1:
            GetcwguitarInfo()
            break;

    }
    console.log(parseInfo)

}
ParseNowUrlInfo()