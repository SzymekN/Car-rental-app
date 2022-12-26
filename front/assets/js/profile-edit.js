function edit(){
    const profile=JSON.parse(localStorage.getItem("profile"));
    const changeData = {};
    var nameVal=document.getElementById("name").value;
    var surnameVal=document.getElementById("surname").value;
    var phone_numberVal=document.getElementById("phone_number").value
    var emailVal=document.getElementById("email").value
    if(nameVal!=profile.name)
        Object.assign(changeData,{name:nameVal});
    if(surnameVal!=profile.surname)
        Object.assign(changeData,{surname:surnameVal});
    if(phone_numberVal!=profile.phone_number)
        Object.assign(changeData,{phone_number:phone_numberVal});
    if(emailVal!=profile.email)
        Object.assign(changeData,{email:emailVal});
    //console.log(changeData)
    if(Object.keys(changeData).length==0)
      alert("Nie zmieniono żadnej wartości!");
    else{
      Promise.ressolve(editData(changeData));
      document.location.href = "car-settings.html";
    }
   
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
              return res(data);
          }).then(res.toString).catch( err => {
              return rej(err);                         
          });                                              
  });
}
async function loadProfileInfo(){
  var profile=await getProfileInfo();
  localStorage.setItem("profile",JSON.stringify(profile));
  document.getElementById("name").value=profile.name;
  document.getElementById("surname").value=profile.surname;
  document.getElementById("email").value=profile.email;
  document.getElementById("phone_number").value=profile.phone_number;
}
function getProfileInfo(){
  var target="http://192.168.33.50:8200/api/v1/clients/profileInfo";
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