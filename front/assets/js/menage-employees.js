async function loadEmployees(currentPage=0){
    var response=await getEmployees();
    console.log(response);
    var i=0,maxEmployees=30;

    if(Object.entries(response).length!=0){
        for (i = currentPage* maxEmployees; i < (currentPage* maxEmployees)+ maxEmployees; i++) {
            if(i<Object.keys(response).length){
                let elem = document.createElement('tr');
                elem.append(tmp2.content.cloneNode(true));
                let td=document.createElement('td');
                
                // td.innerHTML=response[i].id;
                // elem.appendChild(td);
                // document.getElementById("tableBody").appendChild(elem);
                // td.innerHTML=response[i].name;
                // elem.appendChild(td);
                // document.getElementById("tableBody").appendChild(elem);
                
                console.log(JSON.stringify(response[i].userId))
                elem.querySelector("#name").innerHTML=response[i].name;
                elem.querySelector("#surname").innerHTML=response[i].surname;
                elem.querySelector("#email").innerHTML=response[i].email;//
                elem.querySelector("#role").innerHTML=response[i].role;//
                elem.querySelector(".btn-danger").id=response[i].id;

                elem.querySelector(".btn-warning").id=response[i].id;
                
                document.getElementById("tableBody").appendChild(elem);
            }
        }
    }
}
function getEmployees(){
    var target="http://192.168.33.50:8200/api/v1/employees/all";
    event.preventDefault();
      return new Promise(async (res, rej) => {                       
        await fetch(target, {method: "GET",mode: 'cors',
        headers: {
          "Content-Type": "application/json",
          "Authorization":"Bearer "+localStorage.getItem("token")
        }}).then(async (r) => {   
        
          const data =  await r.json();
          if(!r.ok)
          {
            const error = (data && data.message) || r.status;
            return Promise.reject(error);
          }
            return res(data);
        }).then(res.toString).catch( err => {
            return rej(err);                        
        });                                              
    });
}