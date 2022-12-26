function edit(){
    //createNewReservation();

    const changeData = {};
    var nameVal=document.getElementById("name").value;
    var surnameVal=document.getElementById("surname").value;
    var phone_numberVal=document.getElementById("phone_number").value
    var emailVal=document.getElementById("email").value
    if(nameVal)
        rentData=Object.assign(changeData,{name:nameVal});
    if(surnameVal)
         rentData=Object.assign(changeData,{surname:surnameVal});
    if(phone_numberVal)
         rentData=Object.assign(changeData,{phone_number:phone_numberVal});
    if(emailVal)
         rentData=Object.assign(changeData,{email:emailVal});
    //console.log(changeData)
    if(Object.keys(changeData).length==0)
      alert("Nie zmieniono żadnej wartości!");
   
}
function editData(data){
    var target="http://192.168.33.50:8200/api/v1/clients/self";
    event.preventDefault();
        return new Promise(async (res, rej) => {                       
          await fetch(target, {method: "PUT",mode: 'cors',body: JSON.stringify(data),
          headers: {
            "Content-Type": "application/json",
            "Authorization":"Bearer "+localStorage.getItem("token")
          }}).then(async (r) => {   // fetch the resourse
            const data =  await r.json();
            if(!r.ok)
            {
              const error = (data && data.message) || r.status;
              return Promise.reject(error);
            }
              var cost=loadCar(data);
              return res(cost);
          }).then(res.toString).catch( err => {
              return rej(err);                         
          });                                              
  });
}