async function loadReports(currentPage=0){
    var reportsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/notifications/all","GET");
    var i=0,maxReports=30;
    var user,email;
    
    console.log(reportsInfo);
    if(Object.entries(reportsInfo).length!=0){
        for (i = currentPage* maxReports; i < (currentPage* maxReports)+ maxReports; i++) {
            if(i<Object.keys(reportsInfo).length){
                //var currentReport ={id:reportsInfo[i].id};
                //var reportInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",currentReport);

                let elem = document.createElement('tr');
                elem.append(tmp4.content.cloneNode(true));
                let td=document.createElement('td');
                var userInfo={}
                //console.log(typeof reportsInfo[i].employee_id1!=="undefined")
                if(typeof reportsInfo[i].employee_id!=="undefined")
                    Object.assign(userInfo,{id:reportsInfo[i].employee_id});
                else if(typeof reportsInfo[i].client_id!=="undefined")
                    Object.assign(userInfo,{id:reportsInfo[i].client_id});
                //Object.assign(userInfo,{id:reportsInfo[i].employee_id});

                if(Object.keys(userInfo).length!=0){
                    user= await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",userInfo);
                    email=user.email;
                }
                else 
                    email="undefined";

                elem.querySelector("#id").innerHTML=reportsInfo[i].id;
                elem.querySelector("#email").innerHTML=email;
                elem.querySelector("#message").innerHTML=reportsInfo[i].description;
                elem.querySelector(".btn-danger").id=reportsInfo[i].id;
                
                document.getElementById("tableBody").appendChild(elem);
            }
        }
    }
}
// function getReportID(id){
//     localStorage.setItem("reportId",id);
// }
async function reportDelete(){
    var currentReport ={id:parseInt(localStorage.getItem("reportId"))};
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/reports","DELETE",currentReport);
    document.location.href = "menage-reports.html";
}

async function sendReport(){
    const reportInfo={};
    var descVal=document.getElementById("description").value;
    if(descVal.length!=0){
        await getInfoWithBody("http://192.168.33.50:8200/api/v1//notifications","POST",{description:descVal});
        alert("Pomyślnie wysłano zgłoszenie");
        reload();
    }
    else
        alert("Wiadomość jest pusta")
    
}

async function deleteReport(idVal){
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/notifications","DELETE",{id:parseInt(idVal)});
    reload();
}