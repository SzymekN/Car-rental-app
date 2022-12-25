//async
function login () {
    var t = "http://192.168.33.50:8200/api/v1/users/signin";
    event.preventDefault();
const fetchPromise=chkLogin(t, 3, 1000).then((response) => response.json())
.then((data) => {
  localStorage.setItem("token",data.token)
  if(data.token)
    loginSuccess();
  else
    alert("Wrong credentials!");
 
}).catch( err => {
    console.log('error: '+ err)
});
fetchPromise;
}

function chkLogin(target, times, delay) {
  const login = {
    email: document.getElementById('email').value,
    password: document.getElementById('password').value
  }
  return new Promise((res, rej) => {                       // return a promise
      fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(login),
      headers: {
        "Content-Type": "application/json; charset=UTF-8",
        "Content-Length":"217"
      }}).then((r) => {   // fetch the resourse
          res(r);                                     // resolve promise if success
      }).then(res.toString).catch( err => {
              return rej(err);                         // don't try again
      });                                              // again until no more tries
   
  });
}

async function loginSuccess(){
  if(localStorage.getItem("token")){//!="undefined"
    document.location.href = "user-rent.html";
  }
}

//async function getAllCars(){
  // var res="http://192.168.33.50:8200/api/v1/vehicles/all";
  // //event.preventDefault();
  // const fetchPromise=await getCars(res, 3, 1000).then(async response =>{
  //   const data=await response.body.json();
   // const map = new Map(Object.entries(JSON.stringify(data))); //We first convert the string to an object and then to an array, because we can’t parse a JSON string to a Map directly. 
    // localStorage.setItem("allCars",await data)
    // console.log(await data)
    // loginSuccess()
  // }).catch( err => {
  //     console.log('error: '+ err)
  // });
  // fetchPromise;
//}

