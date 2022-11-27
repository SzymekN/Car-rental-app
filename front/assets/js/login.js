function chk(target, times, delay) {
  const login = {
    username: "asd",
    password: "asd"
  }
  return new Promise((res, rej) => {                       // return a promise
      (function rec(i) {                                   // recursive IIFE
          fetch(target, {method: "POST",mode: 'no-cors',body: JSON.stringify(login),
          headers: {
            "Content-Type": "application/json; charset=UTF-8",
            "Date":"Sun, 27 Nov 2022 12:08:09 GMT",
            "Content-Length":"217"
          }}).then((r) => {   // fetch the resourse
              res(r);                                      // resolve promise if success
          }).catch( err => {
              if (times === 0)                             // if number of tries reached
                  return rej(err);                         // don't try again

              setTimeout(() => rec(--times), delay )       // otherwise, wait and try 
          });                                              // again until no more tries
      })(times);

  });
}
async function postData () {
    /*const login = {
      username: "asd",
      password: "asd"
    }

    const response = await fetch("//172.19.0.4:8200/api/v3/operators/signin", {
      method: "POST",
      mode:'cors',
      body: JSON.stringify(login),
      headers: {
        "Content-Type": "application/json; charset=UTF-8",
        "Date":"Sun, 27 Nov 2022 12:08:09 GMT",
        "Content-Length":"217"
      }
    })
   
    if (!response.ok) {
      throw new Error(`Request failed with status ${response.status}`)
    }
    alert("f")
    console.log("Request successful!")
    */
   
    var t = "http://172.19.0.7:8200/api/v3/operators/signin";
    event.preventDefault();
chk(t, 3, 1000).then( image => {
    console.log('success')
}).catch( err => {
    console.log('error: '+ err)
});
  }