async function getClient(idVal){

  var client= await getInfoWithBody("http://192.168.33.50:8200/api/v1/clients/info","POST",{id:parseInt(idVal)})
  localStorage.setItem("currentClient",JSON.stringify(client));
  //console.log(employee);
  var userInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",{id:parseInt(client.userId)});
  localStorage.setItem("currentUser",JSON.stringify(userInfo));
  console.log(client.phone_number)
  document.getElementById("modalName").value=client.name;
  document.getElementById("modalSurname").value=client.surname;
  document.getElementById("modalEmail").value=userInfo.email;
  document.getElementById("modalPhonenumber").value=client.phone_number;
}
async function loadClientsEmp(currentPage=0){
    var clientsInfo=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/clients/all","GET");
    var i=0,maxClients=30;
    console.log(clientsInfo)
    if(Object.entries(clientsInfo).length!=0){
        for (i = currentPage* maxClients; i < (currentPage* maxClients)+ maxClients; i++) {
            if(i<Object.keys(clientsInfo).length){
                var currentUser ={id:clientsInfo[i].userId};
                var userInfo=await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/info","POST",currentUser);
                if(userInfo.role==="client"){
                    let elem = document.createElement('tr');
                    elem.append(tmp7.content.cloneNode(true));
                    let td=document.createElement('td');
                    
                    elem.querySelector("#id").innerHTML=clientsInfo[i].id;
                    elem.querySelector("#name").innerHTML=clientsInfo[i].name;
                    elem.querySelector("#surname").innerHTML=clientsInfo[i].surname;
                    elem.querySelector("#email").innerHTML=userInfo.email;
                    elem.querySelector("#phone_number").innerHTML=clientsInfo[i].phone_number;
                    //console.log(clientsInfo[i])
                    elem.querySelector(".btn-warning").id=clientsInfo[i].id;
                    //elem.querySelector(".btn-success").id=clientsInfo[i].userId;

                    
                    document.getElementById("tableBody").appendChild(elem);
                }
            }
        }
    }
   
}

async function editClient(){
    const clientId=localStorage.getItem("currentClientId");
    var client=JSON.parse(localStorage.getItem("currentClient"));
    var user=JSON.parse(localStorage.getItem("currentUser"));
    var clientData={};
    var userData={};
  
    
    var nameVal=document.getElementById("modalName").value;
    var surnameVal=document.getElementById("modalSurname").value;
    var emailVal=document.getElementById("modalEmail").value;
    var numberVal=document.getElementById("modalPhonenumber").value;
   
    //console.log(client.name)
    
    
    if(nameVal!==client.name)
        Object.assign(clientData,{name:nameVal});
    if(surnameVal!==client.surname)
        Object.assign(clientData,{surname:surnameVal});
    if(emailVal!==user.email)
        Object.assign(userData,{email:emailVal});
    if(numberVal!=client.phone_number)
        Object.assign(clientData,{phone_number:numberVal});
   
    
  
    if(Object.keys(clientData).length==0&&Object.keys(userData).length==0)
      alert("Nie zmieniono żadnej wartości!");
    else{
      if(Object.keys(clientData).length!=0){
        if(!isNaN(parseInt(clientId)))
            Object.assign(clientData,{id:parseInt(clientId)})
         console.log(clientData)
          Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/clients","PUT",clientData)).then((data) => {
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
          console.log(userData)
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