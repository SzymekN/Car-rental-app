async function loadNotificationRepairs(){
    var reportsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/notifications/all","GET");
    var i=0,maxReports=30,currentPage=0;
    var user,email;
    
    console.log(reportsInfo);
    if(Object.entries(reportsInfo).length!=0){
        for (i = currentPage* maxReports; i < (currentPage* maxReports)+ maxReports; i++) {
            if(i<Object.keys(reportsInfo).length){
                // to add when vehicle is in repair do not show
                
                //await getInfoWithBody("http://192.168.33.50:8200/api/v1/notifications")
                if(typeof reportsInfo[i].vehicle_id!=="undefined"){
                    console.log(reportsInfo[i])
                let elem = document.createElement('tr');
                elem.append(tmp5.content.cloneNode(true));
                let td=document.createElement('td');
                
                elem.querySelector("#notification_id").innerHTML=reportsInfo[i].id;
                elem.querySelector("#client_id").innerHTML=reportsInfo[i].client_id;
                elem.querySelector("#description").innerHTML=reportsInfo[i].description;
                elem.querySelector("#vehicle_id").innerHTML=reportsInfo[i].vehicle_id;
                
                elem.querySelector(".btn-danger").id=reportsInfo[i].id;
                elem.querySelector(".btn-success").id=reportsInfo[i].id;

                document.getElementById("tableBody").appendChild(elem);
            }
            }
        }
    }
}
function getReportId(idVal){
    localStorage.setItem("currentReport",idVal)
}
async function addRepair(){
    var costVal=document.getElementById("cost").value;
    var currentReportId=localStorage.getItem("currentReport");
    var notificationInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/notifications/info","POST",{id:parseInt(currentReportId)});
    const repairInfo={
        cost:parseInt(costVal),
        notification_id:parseInt(currentReportId),
        approved:1,
        vehicle_id:parseInt(notificationInfo.vehicle_id)
    }
    console.log(repairInfo)
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/repairs","POST",repairInfo);
    reload();
}
async function loadRepairs(idVal){
    var repairsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/repairs/all","GET");
    var i=0,maxReports=30,currentPage=0;
    if(Object.entries(repairsInfo).length!=0){
        for (i = currentPage* maxReports; i < (currentPage* maxReports)+ maxReports; i++) {
            if(i<Object.keys(repairsInfo).length){
                if(typeof repairsInfo[i].vehicle_id!=="undefined"){
                let elem = document.createElement('tr');
                elem.append(tmp6.content.cloneNode(true));
                let td=document.createElement('td');
                

                elem.querySelector("#id").innerHTML=repairsInfo[i].id;
                elem.querySelector("#cost").innerHTML=repairsInfo[i].cost;
                elem.querySelector("#approved").innerHTML="Tak";
                elem.querySelector("#vehicle_id").innerHTML=repairsInfo[i].vehicle_id;

                var notificationInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/notifications/info","POST",{id:repairsInfo[i].notification_id});

                elem.querySelector("#client_id").innerHTML=notificationInfo.client_id;
                elem.querySelector("#description").innerHTML=notificationInfo.description;

                elem.querySelector(".btn-danger").id=repairsInfo[i].id;
                document.getElementById("tableBody").appendChild(elem);
            }
            }
        }
    }
}

async function deleteRepair(idVal){
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/repairs","DELETE",{id:parseInt(idVal)});
    reload();
}