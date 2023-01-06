async function selfDelete(){
  Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/clients/self","DELETE",document.getElementById('password').value)).then(cars=>{
      document.location.href = "index.html";    
});
}