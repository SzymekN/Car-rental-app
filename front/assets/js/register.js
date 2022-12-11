function chkRegister(target, times, delay) {
    var login = {
      name: document.getElementById('name').value,
      surname: document.getElementById('surname').value,
      pesel: document.getElementById('pesel').value,
      phone_number: document.getElementById('phone_number').value,
      user:{
        email: document.getElementById('email').value,
        password: document.getElementById('password').value,
      }
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
  const fetchPromise=chkRegister(t, 3, 1000).then((response) => response.json())
  .then((data) => {
    console.log(data)
    localStorage.setItem("token",data.token)
    //console.log(`localStorage set with token value: ${data.token}`)
    registrationSuccess();
  }).catch( err => {
      console.log('error: '+ err)
  });
  fetchPromise;
}

  async function registrationSuccess(){
    if(localStorage.getItem("token")!="undefined"){
      document.location.href = "user-menu.html";
    }
  }