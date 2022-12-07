function chkLogout(target, times, delay) {

    return new Promise((res, rej) => {                       // return a promise
        fetch(target, {method: "GET",mode: 'cors',body: JSON.stringify(login),
        headers: {
          "Content-Type": "application/json; charset=UTF-8",
          "Content-Length":"217",
          "Authorization":"Bearer "+localStorage.getItem("token")
        }}).then((r) => {   // fetch the resourse
            res(r);                                      // resolve promise if success
        }).then(res.toString).catch( err => {
            return rej(err);                         // don't try again 
        });                                              // again until no more tries
    });
  }
  async function logout () {
      var t = "http://192.168.33.50:8200/api/v1/users/signout";
      event.preventDefault();
  if(localStorage.getItem("token")!=null){
  const fetchPromise=chkLogout(t, 3, 1000).then((response) => response.json())
  .then((data) => {
    console.log(data)
    localStorage.removeItem("token")
    logoutAction();
  }).catch( err => {
      console.log('error: '+ err)
      alert("Wrong token!")
  });
  fetchPromise;
}else{
  alert("Log in first!");
  logoutAction();
  }

}

  async function logoutAction(){
    if(localStorage.getItem("token")==null){
      document.location.href = "index.html";
    }else(console.log('error: token still exists!'))
  }