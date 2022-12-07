function chkLogin(target, times, delay) {
  const login = {
    email: document.getElementById('email').value,
    password: document.getElementById('password').value
  }
  return new Promise((res, rej) => {                       // return a promise
     //(function rec(i) {                                    // recursive IIFE
      fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(login),
      headers: {
        "Content-Type": "application/json; charset=UTF-8",
        "Content-Length":"217"
      }}).then((r) => {   // fetch the resourse
          res(r);                                      // resolve promise if success
      }).then(res.toString).catch( err => {
          //if (times === 0)                             // if number of tries reached
              return rej(err);                         // don't try again
          //setTimeout(() => rec(--times), delay )       // otherwise, wait and try 
      });                                              // again until no more tries
      //})(times);

  });
}
async function login () {
    var t = "http://192.168.33.50:8200/api/v1/users/signin";
    event.preventDefault();
const fetchPromise=chkLogin(t, 3, 1000).then((response) => response.json())
.then((data) => {
  console.log(data)
  localStorage.setItem("token",data.token)
  //console.log(`localStorage set with token value: ${data.token}`)
  loginSuccess();
}).catch( err => {
    console.log('error: '+ err)
});
fetchPromise;
}


async function loginSuccess(){
  if(localStorage.getItem("token")!="undefined"){
    document.location.href = "user-menu.html";
  }
}