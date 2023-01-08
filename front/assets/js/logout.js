function logout () {
  if(localStorage.getItem("token")!=null){
    Promise.resolve(getInfoWithoutBody("http://192.168.33.50:8200/api/v1/users/signout","GET")).then((data) => {
      console.log(data)
      localStorage.removeItem("token")
      logoutAction();
    }).catch( err => {
      console.log('error: '+ err)
      alert("Wrong token!")
  });
}else{
  alert("Log in first!");
  logoutAction();
  }
}
function logoutAction(){
    if(localStorage.getItem("token")==null){
      document.location.href = "index.html";
      //localStorage.clear();
    }else(console.log('error: token still exists!'))
  }