async function pay(){
    var email=document.getElementById('email').value;
    console.log(document.body.contains(document.getElementById('email')))
   
    await createNewRent(email);
    //console.log(JSON.stringify(localStorage.getItem("startDate")));

    if(document.body.contains(document.getElementById('email'))&&email.length==0){
      alert("Nie wypełniono pola email!")
      document.location.href = "employee-checkout.html";
    }
    else if(email.length!=0)
      document.location.href = "employee-rent.html";
    else
      document.location.href="user-reservations.html";
  }
function createNewRent(emailVal){
    var rentInfo = {
        start_date: new Date(localStorage.getItem("startDate")),
        end_date: new Date(localStorage.getItem("endDate")),
        pickup_address: JSON.stringify(localStorage.getItem("pickupAdress")),
        vehicle_id: parseInt(localStorage.getItem("currentCar"))
    }
    if(emailVal.length!=0)
      Object.assign(rentInfo,{email:emailVal});
    console.log(rentInfo);
    var target="http://192.168.33.50:8200/api/v1/rentals/self";
    event.preventDefault();
        return new Promise(async (res, rej) => {                       
          await fetch(target, {method: "POST",mode: 'cors',body: JSON.stringify(rentInfo),
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
              return res(data);
          }).then(res.toString).catch( err => {
              return rej(err);                         
          });                                              
  });
}

async function loadRent(currentPage=0){
  var response=await getRent();
  console.log(response);
  var temp, item, a, i=0,maxCarsPage=30,status;
  var car,elem=0;
  
  //temp = document.getElementsByTagName("template")[0];
  
  //item = document.querySelector("#rentDiv");
  //console.log(item)
  if(Object.entries(response).length!=0){
      
    for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {
      if(i<Object.keys(response).length){
        //console.log(response[0].vehicle_id);
        car=await Promise.resolve(getRentCar(response[i].vehicle_id));
        //console.log(formatDateOrder(new Date(response[i].end_date)));
        if(response[i].end_date<new Date())
          status="Zakończone";
        else 
          status="W trakcie";
        let elem = document.createElement('div');
        if(i!=0)
          elem.style="border-top: 1px solid var(--bs-black);"
        elem.append(tmpl.content.cloneNode(true));
        //elem.content.getElementById("name").textContent="F";
        //console.log(elem.querySelector("#name"));
        elem.querySelector("#name").innerHTML=[car.brand,car.model].join(' ');
        elem.querySelector("#daily_cost").innerHTML=["Dzienny koszt:",car.dailyCost].join(' ')
        elem.querySelector("#fuel_consumption").innerHTML=["Spalanie:",car.fuelConsumption].join(' ')
        elem.querySelector("#status").innerHTML=["Status:",status].join("<br>");
        elem.querySelector("#start_date").innerHTML=["Rozpoczęcie wynajmu:",formatDateOrder(new Date(response[i].start_date))].join("<br>")
        elem.querySelector("#rent_cost").innerHTML=["Zapłacona kwota:",car.dailyCost].join("<br>")
        elem.querySelector("#end_date").innerHTML=["Zakończenie wynajmu:",formatDateOrder(new Date(response[i].end_date))].join("<br>")


        // elem.textContent=[car.brand,car.model].join(' ');
        // let p=a.querySelectorAll("h5");
        // p[0].textContent=["Dzienny koszt:",filteredCars[i].dailyCost].join(' ');
        // p[1].textContent=["Spalanie:",filteredCars[i].fuelConsumption].join(' ');
        
         //dla każdego następnego border u góry
        document.getElementById("rentGroup").appendChild(elem);
     
     
        
    }
  }
}
  else{
     let a=document.createElement("h4");//
     a.style.color='black';
     a.style.marginBottom='25px';
     a.style.marginLeft='15px';
     a.textContent="Nie wpożyczono jeszcze samochodu.";
     document.getElementById("rentGroup").appendChild(a);
  }

}
function getRent(){
var target="http://192.168.33.50:8200/api/v1/rentals/self";
event.preventDefault();
    return new Promise(async (res, rej) => {                       
      await fetch(target, {method: "GET",mode: 'cors',
      headers: {
        "Content-Type": "application/json",
        "Authorization":"Bearer "+localStorage.getItem("token")
      }}).then(async (r) => {   // fetch the resourse
        const data =  await r.json();
        // if(!r.ok)
        // {
        //   const error = (data && data.message) || r.status;
        //   return Promise.reject(error);
        // }
          return res(data);
      }).then(res.toString).catch( err => {
          if(err===400)
            return null;
          return rej(err);                         
      });                                              
});
}


function getRentCar(carId){
const currentCar ={id:carId};
console.log(typeof parseInt(localStorage.getItem("currentCar")))
var target="http://192.168.33.50:8200/api/v1/vehicles/single";
//event.preventDefault();
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
        //var cost=loadRentedCar(data);
        return res(data);
    }).then(res.toString).catch( err => {
        return rej(err);                         // don't try again 
    });                                              // again until no more tries
});
}

function loadRentedCar(car){
//alert("f");
  console.log(car);
  document.getElementById("currentCar").textContent=[car.brand,car.model].join(' ');
  document.getElementById("registrationNumber").textContent=car.registrationNumber;
  document.getElementById("dailyCost").textContent=car.dailyCost;
  document.getElementById("fuelConsumption").textContent=car.fuelConsumption;
  return car.dailyCost;
  
}