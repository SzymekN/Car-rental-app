async function selfDelete(){
  const data = {}
  Object.assign(data,{password:document.getElementById('password').value});
  console.log(JSON.stringify(data))
  Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/clients/self","DELETE",data)).then(cars=>{
      document.location.href = "index.html";    
});
}
