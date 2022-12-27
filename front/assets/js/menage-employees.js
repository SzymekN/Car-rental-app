async function loadEmployees(currentPage=0){
    var employeesInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/employees/all","GET");
    var i=0,maxEmployees=30;

    if(Object.entries(employeesInfo).length!=0){
        for (i = currentPage* maxEmployees; i < (currentPage* maxEmployees)+ maxEmployees; i++) {
            if(i<Object.keys(employeesInfo).length){
                var currentUser ={id:employeesInfo[i].userId};
                var userInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",currentUser);

                let elem = document.createElement('tr');
                elem.append(tmp2.content.cloneNode(true));
                let td=document.createElement('td');
                
                elem.querySelector("#name").innerHTML=employeesInfo[i].name;
                elem.querySelector("#surname").innerHTML=employeesInfo[i].surname;
                elem.querySelector("#email").innerHTML=userInfo.email;
                elem.querySelector("#role").innerHTML=userInfo.role;
                elem.querySelector(".btn-danger").id=employeesInfo[i].id;

                elem.querySelector(".btn-warning").id=employeesInfo[i].id;
                
                document.getElementById("tableBody").appendChild(elem);
            }
        }
    }
}
function getEmployeeID(id){
    localStorage.setItem("employeeId",id);
}
async function employeeDelete(){
    var currentEmployee ={id:parseInt(localStorage.getItem("employeeId"))};
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/employees","DELETE",currentEmployee);
    document.location.href = "menage-employees.html";
}

// function getEmployees(){
//     var target="http://192.168.33.50:8200/api/v1/employees/all";
//     event.preventDefault();
//       return new Promise(async (res, rej) => {                       
//         await fetch(target, {method: "GET",mode: 'cors',
//         headers: {
//           "Content-Type": "application/json",
//           "Authorization":"Bearer "+localStorage.getItem("token")
//         }}).then(async (r) => {   
        
//           const data =  await r.json();
//           if(!r.ok)
//           {
//             const error = (data && data.message) || r.status;
//             return Promise.reject(error);
//           }
//             return res(data);
//         }).then(res.toString).catch( err => {
//             return rej(err);                        
//         });                                              
//     });
// }