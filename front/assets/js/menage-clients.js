async function loadClients(currentPage=0){
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
                elem.append(tmp3.content.cloneNode(true));
                let td=document.createElement('td');
                
                elem.querySelector("#id").innerHTML=clientsInfo[i].id;
                elem.querySelector("#name").innerHTML=clientsInfo[i].name;
                elem.querySelector("#surname").innerHTML=clientsInfo[i].surname;
                elem.querySelector("#email").innerHTML=userInfo.email;
                elem.querySelector("#phone_number").innerHTML=clientsInfo[i].phone_number;
                elem.querySelector(".btn-danger").id=clientsInfo[i].userId;
                elem.querySelector(".btn-success").id=clientsInfo[i].userId;

                
                document.getElementById("tableBody").appendChild(elem);
                }
            }
        }
    }
}
function getUserId(id){
    localStorage.setItem("userId",id);
}
async function userBlock(){
    console.log(parseInt(localStorage.getItem("userId")));
    var currentUser ={id:parseInt(localStorage.getItem("userId"))};
    await getInfoWithBody("http://192.168.33.50:8200/api/v1/users/block","PUT",currentUser);
    document.location.href = "menage-clients.html";
}

async function userUnblock(idVal){
    var currentUser ={id:parseInt(idVal)};
    Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/users/unblock","PUT",currentUser)).then(data=>{
        document.getElementById("password").innerText=data.password;
    });
}
