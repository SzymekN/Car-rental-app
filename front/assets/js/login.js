function login () {
  const login = {};
  var emailVal=document.getElementById('email').value;
  var passwordVal=document.getElementById('password').value
  if(emailVal.length!=0&&passwordVal!=0){
    Object.assign(login,{email:emailVal});
    Object.assign(login,{password:passwordVal});
    console.log(Object.values(login));
    Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/users/signin","POST",login)).then((data) => {
    localStorage.setItem("token",data.token)
    console.log(data)
    if(data.token)
      loginSuccess(data.role);
    }).catch( err => {
      console.log('error: '+ err);
      alert("Złe hasło lub email!");
      var currentLoc=(window.location.href.substring(window.location.href.lastIndexOf('/') + 1));
      if(currentLoc=="index.html")
        document.location.href="index.html";
      else if(currentLoc=="login.html")
        document.location.href = "login.html";
    });

  }
else{
  alert("Nie uzupełniono wszystkich pól!");
  var currentLoc=(window.location.href.substring(window.location.href.lastIndexOf('/') + 1));
    if(currentLoc=="index.html")
      document.location.href="index.html";
    else if(currentLoc=="login.html")
      document.location.href = "login.html";
}
}

async function loginSuccess(role){
  if(localStorage.getItem("token")){
    if(role=="client")
      document.location.href = "user-rent.html";
    else if(role=="employee")
      document.location.href = "employee-rent.html";
    else if(role=="accountant")
      document.location.href = "accountant-main.html";
    else if(role=="driver")
      document.location.href = "driver-main.html";
    else if(role=="admin")
      document.location.href = "menage-employees.html";
  }
}