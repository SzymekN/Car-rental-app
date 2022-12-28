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
                
                elem.querySelector("#id").innerHTML=employeesInfo[i].id;
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


async function addEmployee(){
  var nameVal=document.getElementById('name').value;
  var surnameVal=document.getElementById('surname').value;
  var peselVal=document.getElementById('pesel').value;
  var salaryVal=document.getElementById('salary').value;
  var emailVal=document.getElementById('email').value;
  var passwordVal=document.getElementById('password').value;
  console.log(nameVal.length)
  if(nameVal.length!=0&&surnameVal.length!=0&&peselVal.length!=0&&peselVal.length!=0&&salaryVal.length!=0&&emailVal.length!=0&&passwordVal.length!=0){
  var employeeData = {
    name: nameVal,
    surname: surnameVal,
    pesel: peselVal,
    salary: salaryVal,
    user:{
      email: emailVal,
      password: passwordVal,
      role: "employee"
    }
  }
  var response=Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/employees","POST",employeeData));
  console.log(repsonse);
}else{
  alert("Wszystkie pola muszą być uzupełnione!")
  return null;
}
    
    
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