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
  
    Promise.resolve(getCar()).then(data => {   
        var value=loadCar(data);
        var dateDiff=(end.getTime()-start.getTime())/(1000*3600*24)+1
      document.getElementById("dailyCost").textContent=parseInt(value);

      var rentCost=dateDiff*parseInt(value);
      document.getElementById("rentCost").textContent=rentCost;
      var additional=0;
      if(document.getElementById("adress").value=="Wypożyczalnia")
        document.getElementById("additionalCosts").textContent=0;
      else{
        document.getElementById("additionalCosts").textContent=20;
        additional=20;
      }
      document.getElementById("toPay").textContent=rentCost+additional;
    }).catch( err => {
        console.log(err)                     
    });                                            
    
}

function getCar(){
  const currentCar ={id:parseInt(localStorage.getItem("currentCar"))};
    return getInfoWithBody("http://192.168.33.50:8200/api/v1/vehicles/single","POST",currentCar);
}

function loadCar(car){
    console.log(car);
    document.getElementById("currentCar").textContent=[car.brand,car.model].join(' ');
    document.getElementById("registrationNumber").textContent=car.registrationNumber;
    document.getElementById("dailyCost").textContent=car.dailyCost;
    document.getElementById("fuelConsumption").textContent=car.fuelConsumption;
    document.getElementById("carImage").src=getPhoto(car.brand,car.model);
    return car.dailyCost;
    
}
function  endRent(idVal){
  console.log(idVal);
}