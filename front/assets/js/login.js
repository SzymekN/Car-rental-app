async function login () {
    var t = "http://192.168.33.50:8200/api/v1/users/signin";
    event.preventDefault();
const fetchPromise=chkLogin(t, 3, 1000).then((response) => response.json())
.then((data) => {
  console.log(data)
  localStorage.setItem("token",data.token)
  getAllCars()
 
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
    document.location.href = "user-menu.html";
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

function getAllCars() {
  // const login = {
  //   email: document.getElementById('email').value,
  //   password: document.getElementById('password').value
  // }
  var target="http://192.168.33.50:8200/api/v1/vehicles/all";
  const getData=new Promise(async (res, rej) => {                       // return a promise
    await fetch(target, {method: "GET",mode: 'cors',
    headers: {
      "Content-Type": "application/json; charset=UTF-8",
      "Content-Length":"217",
      "Authorization":"Bearer "+localStorage.getItem("token")
    }}).then(async (r) => {   // fetch the resourse
      // const isJson = r.headers.get('content-type')?.includes('application/json')
      const data =  await r.json();
      if(!r.ok)
      {
        const error = (data && data.message) || r.status;
        return Promise.reject(error);
      }
        //res(r);                                      // resolve promise if success
        return res(data);
    }).then(res.toString).catch( err => {
        return rej(err);                         // don't try again 
    });                                              // again until no more tries
});

getData.then(data=>{
  localStorage.setItem("allCars",JSON.stringify(data))
  console.log(JSON.stringify(data))
  loginSuccess();
  makeFilters(data);
}).catch(err=>console.log(err));
}

function makeFilters(data){
  //const data = new Map(Object.entries(JSON.parse(jsonData)));
  const brand = [...new Set(data.map(item => item.brand))];
  const model = [...new Set(data.map(item => item.model))];
  const type = [...new Set(data.map(item => item.type))]; 
  const color = [...new Set(data.map(item => item.color))];
  const map1= new Map();
  map1.set('brand',brand);
  map1.set('model',model);
  map1.set('type',type);
  map1.set('color',color);
  console.log(map1.get('brand'));
  localStorage.setItem('allFilters',JSON.stringify(Array.from(map1.entries())));
}