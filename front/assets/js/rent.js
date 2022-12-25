function padTo2Digits(num) {
    return num.toString().padStart(2, '0');
}

function formatDateOrder(date) {
    return [
      padTo2Digits(date.getDate()),
      padTo2Digits(date.getMonth() + 1),
      date.getFullYear(),
    ].join('-');
}

function rentCar(){
    
    var start=new Date(JSON.stringify(localStorage.getItem("startDate")));
    var end=new Date(JSON.stringify(localStorage.getItem("endDate")));
    //console.log(end);
    document.getElementById("startDate").textContent=formatDateOrder(start);
    document.getElementById("endDate").textContent=formatDateOrder(end);
    getCar();

}

function getCar(){
  const currentCar ={id:parseInt(localStorage.getItem("currentCar"))};
  console.log(typeof parseInt(localStorage.getItem("currentCar")))
var target="http://192.168.33.50:8200/api/v1/vehicles/single";
event.preventDefault();
    const getData=new Promise(async (res, rej) => {                       // return a promise
      await fetch(target, {method: "POST",mode: 'cors',body: currentCar,
      headers: {
        "Content-Type": "application/json",
        "Authorization":"Bearer "+localStorage.getItem("token")
      }}).then(async (r) => {   // fetch the resourse
        // const isJson = r.headers.get('content-type')?.includes('application/json')
        const data =  await r.json();
        if(!r.ok)
        {
          const error = (data && data.message) || r.status;
          return Promise.reject(error);
        }
          loadCar(data);
          return res(data);
      }).then(res.toString).catch( err => {
          return rej(err);                         // don't try again 
      });                                              // again until no more tries
  });

}

function loadCar(car){
  alert("f");
    console.log(car);
    document.getElementById("registrationNumber").textContent=formatDateOrder(car.registrationNumber);

    
}