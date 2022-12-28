async function loadReports(currentPage=0){
    var reportsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/reports/all","GET");
    var i=0,maxReports=30;

    if(Object.entries(reportsInfo).length!=0){
        for (i = currentPage* maxReports; i < (currentPage* maxReports)+ maxReports; i++) {
            if(i<Object.keys(reportsInfo).length){
                //var currentReport ={id:reportsInfo[i].id};
                //var reportInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",currentReport);

                let elem = document.createElement('tr');
                elem.append(tmp4.content.cloneNode(true));
                let td=document.createElement('td');
                
                elem.querySelector("#id").innerHTML=reportsInfo[i].id;
                elem.querySelector("#email").innerHTML=reportsInfo[i].email;
                elem.querySelector("#message").innerHTML=reportsInfo[i].message;

                elem.querySelector(".btn-danger").id=reportsInfo[i].id;

                
                document.getElementById("tableBody").appendChild(elem);
            }
        }
    }
}
function getReportID(id){
    localStorage.setItem("reportId",id);
}
async function reportDelete(){
    var currentReport ={id:parseInt(localStorage.getItem("reportId"))};
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/reports","DELETE",currentReport);
    document.location.href = "menage-reports.html";
}
