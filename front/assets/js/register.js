
function chkRegister(target) {
  var nameVal=document.getElementById('name').value;
  var surnameVal=document.getElementById('surname').value;
  var peselVal=document.getElementById('pesel').value;
  var phone_numberVal=document.getElementById('phone_number').value;
  var emailVal=document.getElementById('email').value;
  var passwordVal=document.getElementById('password').value;
  console.log(nameVal.length)
  if(nameVal.length!=0&&surnameVal.length!=0&&peselVal.length!=0&&peselVal.length!=0&&phone_numberVal.length!=0&&emailVal.length!=0&&passwordVal.length!=0){
  var regData = {
    name: nameVal,
    surname: surnameVal,
    pesel: peselVal,
    phone_number: phone_numberVal,
    user:{
      email: emailVal,
      password: passwordVal,
    }
  }
}else{
  alert("Wszystkie pola muszą być uzupełnione!")
  return null;
}
  return getInfoWithBody(target,"POST",regData);
}
  async function register () {
      var t = "http://192.168.33.50:8200/api/v1/users/signup";
      event.preventDefault();
    Promise.resolve(chkRegister(t)).then((data) => {
      registrationSuccess();
    }).catch( err => {
        console.log('error: '+ err);
        alert("Złe dane");
      });      
}

  async function registrationSuccess(){  
    var myModal = new bootstrap.Modal(document.getElementById("success"),{});
    myModal.show();
  }
  function clickModal(){
    var currentLoc=(window.location.href.substring(window.location.href.lastIndexOf('/') + 1))
    if(currentLoc=="employee-create.html")
      document.location.href="employee-create.html";
    else if(currentLoc=="register.html")
      document.location.href = "login.html";
  }