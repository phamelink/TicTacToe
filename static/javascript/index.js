function sendIndex(index) {
    const XHR = new XMLHttpRequest(),
        FD = new FormData();
    FD.append("index", index)
    XHR.open('POST', '192.168.1.129:8000')
    XHR.send(FD)
    location.reload()
}

function reset() {
    const XHR = new XMLHttpRequest(),
        FD = new FormData();
    FD.append("index", "-1")
    XHR.open('POST', '192.168.1.129:8000')
    XHR.send(FD)
    location.reload()
}

function hint() {
    const XHR = new XMLHttpRequest(),
        FD = new FormData();
    FD.append("index", "-2")
    XHR.open('POST', '192.168.1.129:8000')
    XHR.send(FD)
    location.reload()
}