async function selfDelete(){
    Promise.resolve(deleteRequest(document.getElementById('password').value)).then(cars=>{
        document.location.href = "index.html";    
});
}
function deleteRequest(input){
    var target="http://192.168.33.50:8200/api/v1/clients/self";
    var data={password:input};
    event.preventDefault();
        return new Promise(async (res, rej) => {                       
          await fetch(target, {method: "DELETE",mode: 'cors',body: JSON.stringify(data),
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