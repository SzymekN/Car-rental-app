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
        if(newVal==confirmVal){
            Object.assign(updateData,{old_password:oldVal});
            Object.assign(updateData,{new_password:newVal});
            Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/clients/update/password","PUT",updateData)).then((data) => {
                passwordChangeSuccess();  
            }).catch( err => {
                  console.log('error: '+ err);
                  alert("Niepoprawne stare hasło!");
                  document.location.href=window.location.href.substring(window.location.href.lastIndexOf('/') + 1);
                });
            
        }
        else{
            alert("Hasła nie są takie same!")
            document.location.href = "user-password.html";
        }
    }
    else{
        alert("Wszystkie pola muszą być uzupełnione!")
        document.location.href = "user-password.html";
    }
}
function passwordChangeSuccess(){
    alert("Pomyślnie zmieniono hasło.\nZaloguj się ponownie.")
    logout();
}