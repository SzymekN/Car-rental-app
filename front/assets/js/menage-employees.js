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
  var salaryVal=parseInt(document.getElementById('salary').value);
  var emailVal=document.getElementById('email').value;
  var passwordVal=document.getElementById('password').value;

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
  await Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/employees","POST",employeeData));
  document.location.href="menage-employees.html";
}else{
  alert("Wszystkie pola muszą być uzupełnione!")
  return null;
}
}

async function getEmployee(idVal){
  
  if(document.getElementById("roleList").innerHTML.length)
    document.getElementById("roleList").innerHTML="";
  var employee= await getInfoWithBody("http://192.168.33.50:8200/api/v1/employees/info","POST",{id:parseInt(idVal)})
  localStorage.setItem("currentEmployee",JSON.stringify(employee));
  var userInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",{id:parseInt(employee.userId)});
  localStorage.setItem("currentUser",JSON.stringify(userInfo));
  document.getElementById("role").innerText=userInfo.role;
  
  document.getElementById("modalName").value=employee.name;
  document.getElementById("modalSurname").value=employee.surname;
  document.getElementById("modalEmail").value=userInfo.email;

  a=document.createElement("li");
  if(userInfo.role==="employee")
    a.classList.add("disabled");
  else{
    a.id="employee";
    a.addEventListener('click', function handleClick(event) {
      localStorage.setItem("role",this.id);
      document.getElementById("role").innerText=this.id;
    });
    }
  a.appendChild(document.createTextNode("employee"));
  a.classList.add("dropdown-item");
  document.getElementById("roleList").appendChild(a);

  a=document.createElement("li");
  if(userInfo.role==="admin")
    a.classList.add("disabled");
  else{
   a.id="admin";
   a.addEventListener('click', function handleClick(event) {
    localStorage.setItem("role",this.id);
    document.getElementById("role").innerText=this.id;
  });
  }
  a.appendChild(document.createTextNode("admin"));
  a.classList.add("dropdown-item");
  document.getElementById("roleList").appendChild(a);
}

async function editEmployee(){
  const employeeId=localStorage.getItem("employeeId");
  var employee=JSON.parse(localStorage.getItem("currentEmployee"));
  var user=JSON.parse(localStorage.getItem("currentUser"));
  var employeeData={};
  var userData={};

  
  var nameVal=document.getElementById("modalName").value;
  var surnameVal=document.getElementById("modalSurname").value;
  var emailVal=document.getElementById("modalEmail").value;
  var roleVal=document.getElementById("role").innerText;

  if(nameVal!==employee.name)
      Object.assign(employeeData,{name:nameVal});
  if(surnameVal!==employee.surname)
      Object.assign(employeeData,{surname:surnameVal});
  if(emailVal!==user.email)
      Object.assign(userData,{email:emailVal});
  if(roleVal!==user.role)
      Object.assign(userData,{role:roleVal});
  console.log(employeeData)
  console.log(userData)

  if(Object.keys(employeeData).length==0&&Object.keys(userData).length==0)
    alert("Nie zmieniono żadnej wartości!");
  else{
    if(Object.keys(employeeData).length!=0){
      if(!isNaN(parseInt(employeeId)))
          Object.assign(employeeData,{id:parseInt(employeeId)})
        Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/employees","PUT",employeeData)).then((data) => {
        alert("Pomyślnie zmieniono dane.");
        reload();
      }).catch( err => {
        console.log('error: '+ err);
        alert("Wprowadzono złe dane!");
        reload();
      });
    }
    else if(Object.keys(userData).length!=0){
      if(!isNaN(parseInt(user.id)))
        Object.assign(userData,{id:parseInt(user.id)})
      Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/users","PUT",userData)).then((data) => {
        alert("Pomyślnie zmieniono dane.");
        reload();
      }).catch( err => {
        console.log('error: '+ err);
        alert("Wprowadzono złe dane!");
        reload();
      });
    }
  }
}