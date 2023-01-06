// zawsze daje error
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
      Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/clients/self","PUT",changeData)).then((data) => {
        alert("Pomyślnie zmieniono dane.");
        reload();
    }).catch( err => {
        console.log('error: '+ err);
        alert("Wprowadzono złe dane!");
        reload();
      });;
      
    }
}

async function loadProfileInfo(){
  var profile=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/clients/profileInfo","GET");
  localStorage.setItem("profile",JSON.stringify(profile));
  document.getElementById("name").value=profile.name;
  document.getElementById("surname").value=profile.surname;
  document.getElementById("email").value=profile.email;
  document.getElementById("phone_number").value=profile.phone_number;
}
