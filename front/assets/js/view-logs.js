async function loadLogs(currentPage=0){
    var logsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/logs/all","GET");
    console.log(logsInfo)
    var i=0,maxLogs=30;

    if(Object.entries(logsInfo).length!=0){
        for (i = currentPage* maxLogs; i < (currentPage* maxLogs)+ maxLogs; i++) {
            if(i<Object.keys(logsInfo).length){
                let elem = document.createElement('tr');
                elem.append(tmp4.content.cloneNode(true));
                let td=document.createElement('td');
                
                elem.querySelector("#timestamp").innerHTML=logsInfo[i].timestamp;
                elem.querySelector("#message").innerHTML=logsInfo[i].value;
                
                document.getElementById("tableBody").appendChild(elem);
            }
        }
    }
}