function queryStatus() {
    fetch("/status")
        .then((response) => {
            return response.text()
        })
        .then((txt) => {
            let rootObj = document.getElementById("tableHolder");
            let prevTable = document.getElementById("statusTable");
            if (prevTable != null) {
                rootObj.removeChild(prevTable);
            }

            rootObj.innerHTML = txt;
        })
}

setInterval(queryStatus, 500)