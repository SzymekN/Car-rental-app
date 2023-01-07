async function loadReports(currentPage=0){
    var reportsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/notifications/all","GET");
    var i=0,maxReports=30;
    var user,email;
    
    console.log(reportsInfo);
    if(Object.entries(reportsInfo).length!=0){
        for (i = currentPage* maxReports; i < (currentPage* maxReports)+ maxReports; i++) {
            if(i<Object.keys(reportsInfo).length){

                if(typeof reportsInfo[i].vehicle_id==="undefined"){
                let elem = document.createElement('tr');
                elem.append(tmp4.content.cloneNode(true));
                let td=document.createElement('td');
                var userInfo={}
                if(typeof reportsInfo[i].employee_id!=="undefined"){
                    
                    user= await getInfoWithBody("http://192.168.33.50:8200/api/v1/employees/info","POST",{id:reportsInfo[i].employee_id});
                    user= await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",{id:user.userId});
                    email=user.email;
                }
                else if(typeof reportsInfo[i].client_id!=="undefined"){
                    user= await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",{id:reportsInfo[i].client_id});
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
}

// async function reportDelete(){
//     var currentReport ={id:parseInt(localStorage.getItem("reportId"))};
//     await getInfoWithBody("http://192.168.33.50:8200/api/v1/reports","DELETE",currentReport);
//     document.location.href = "menage-reports.html";
// }

async function sendReport(){
    var descVal=document.getElementById("description").value;
    //const reportInfo={description:descVal};
    if(descVal.length!=0){
        var target;
        if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)==="employee-report.html")
            target="http://192.168.33.50:8200/api/v1/notifications/employee";
        else
            target="http://192.168.33.50:8200/api/v1/notifications/client";

        await getInfoWithBody(target,"POST",{description:descVal});
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