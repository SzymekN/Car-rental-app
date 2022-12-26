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

async function rentCar(){
    
    var start=new Date(JSON.stringify(localStorage.getItem("startDate")));
    var end=new Date(JSON.stringify(localStorage.getItem("endDate")));
    document.getElementById("startDate").textContent=formatDateOrder(start);
    document.getElementById("endDate").textContent=formatDateOrder(end);

    const value=await getCar();
    document.getElementById("dailyCost").textContent=parseInt(value);

    var rentCost=(end-start+1)*parseInt(value);
    document.getElementById("rentCost").textContent=rentCost;
    console.log(document.getElementById("adress").value)

    var additional=0;

    if(document.getElementById("adress").value=="Wypożyczalnia")
      document.getElementById("additionalCosts").textContent=0;
    else{
      document.getElementById("additionalCosts").textContent=20;
      additional=20;
    }
    document.getElementById("toPay").textContent=rentCost+additional;
}

function getCar(){
  const currentCar ={id:parseInt(localStorage.getItem("currentCar"))};
  console.log(typeof parseInt(localStorage.getItem("currentCar")))
var target="http://192.168.33.50:8200/api/v1/vehicles/single";
event.preventDefault();
    return new Promise(async (res, rej) => {                       // return a promise
      await fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(currentCar),
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
          var cost=loadCar(data);
          return res(cost);
      }).then(res.toString).catch( err => {
          return rej(err);                         // don't try again 
      });                                              // again until no more tries
  });
}

function loadCar(car){
  //alert("f");
    console.log(car);
    document.getElementById("currentCar").textContent=[car.brand,car.model].join(' ');
    document.getElementById("registrationNumber").textContent=car.registrationNumber;
    document.getElementById("dailyCost").textContent=car.dailyCost;
    document.getElementById("fuelConsumption").textContent=car.fuelConsumption;
    return car.dailyCost;
    
}