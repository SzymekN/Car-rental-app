function updatePassword(){
    const updateData = {};
    var oldVal=document.getElementById("old_password").value;
    var newVal=document.getElementById("new_password").value;
    var confirmVal=document.getElementById("confirm_password").value;
   
    if(oldVal&&newVal&&confirmVal){
        if(newVal==oldVal){
            alert("Nowe hasło musi być inne od starego!")
            document.location.href = "user-password.html";
            return;
        }
        if(newVal==confirmVal)
            rentData=Object.assign(updateData,{new_password:newVal});
        else{
            alert("Hasła nie są takie same!")
            document.location.href = "user-password.html";
            return;
        }
    }
    else{
        alert("Wszystkie pola muszą być uzupełnione!")
        document.location.href = "user-password.html";
        return;
    }
    console.log(changeData)
  
}

function update(data){
    var target="http://192.168.33.50:8200/api/v1/clients/update/password";
    event.preventDefault();
        return new Promise(async (res, rej) => {                       
          await fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(data),
          headers: {
            "Content-Type": "application/json",
            "Authorization":"Bearer "+localStorage.getItem("token")
          }}).then(async (r) => {   // fetch the resourse
            const data =  await r.json();
            if(!r.ok)
            {
              const error = (data && data.message) || r.status;
              return Promise.reject(error);
            }
              var cost=loadCar(data);
              return res(cost);
          }).then(res.toString).catch( err => {
              return rej(err);                         
          });                                              
  });
}