//TO DO: nie zezwalać na null
function chkRegister(target, times, delay) {

  var nameVal=document.getElementById('name').value;
  var surnameVal=document.getElementById('surname').value;
  var peselVal=document.getElementById('pesel').value;
  var phone_numberVal=document.getElementById('phone_number').value;
  var emailVal=document.getElementById('email').value;
  var passwordVal=document.getElementById('password').value;
  console.log(nameVal.length)
  if(nameVal.length!=0&&surnameVal.length!=0&&peselVal.length!=0&&peselVal.length!=0&&phone_numberVal.length!=0&&emailVal.length!=0&&passwordVal.length!=0){
  var login = {
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
  return new Promise((res, rej) => {                       // return a promise
      fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(login),
      headers: {
        "Content-Type": "application/json; charset=UTF-8",
        "Content-Length":"217"
      }}).then((r) => {   // fetch the resourse
          res(r);                                      // resolve promise if success
      }).then(res.toString).catch( err => {
          return rej(err);                         // don't try again 
      });                                              // again until no more tries
  });
}
  async function register () {
      var t = "http://192.168.33.50:8200/api/v1/users/signup";
      event.preventDefault();
    var response=await Promise.resolve(chkRegister(t, 3, 1000));
    if(response!=null)
      registrationSuccess();
}

  async function registrationSuccess(){  
    var myModal = new bootstrap.Modal(document.getElementById("success"),{});
    myModal.show();
  }
  function clickModal(){
    document.location.href = "login.html";
  }