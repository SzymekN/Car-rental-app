async function pay(){
    var emailVal;
    var rentInfo = {
      start_date: new Date(localStorage.getItem("startDate")),
      end_date: new Date(localStorage.getItem("endDate")),
      pickup_address: JSON.stringify(localStorage.getItem("pickupAdress")),
      vehicle_id: parseInt(localStorage.getItem("currentCar"))
  }
    if(document.body.contains(document.getElementById('email'))){
      emailVal=document.getElementById('email').value;
      if(email.length==0){
        alert("Nie wypełniono pola email!")
        document.location.href = "employee-checkout.html";
      }
      else{
        Object.assign(rentInfo,{email:emailVal});
      }
    }
    Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/rentals/self","POST",rentInfo)).then((data) => {
    alert("Pomyślnie zarezerwowano pojazd")
    if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)=="employee-checkout.html"){
      document.location.href = "employee-rent.html";
    }
    else
      document.location.href="user-reservations.html";
    }).catch( err => {
        console.log('error: '+ err);
        alert("Rezerwacja samochodu nie powiodła się");

      });
  }


async function loadRent(currentPage=0){
  var response=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/rentals/self","GET");
    var temp, item, a, i=0,maxCarsPage=30,status;
    var car,elem=0;
    
    console.log(Object.values(response))
    if(Object.entries(response).length!=0){
        
      for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {
        if(i<Object.keys(response).length){
  
          const currentCar ={id:response[i].vehicle_id};
          console.log(typeof parseInt(localStorage.getItem("currentCar")))
          var car =await getInfoWithBody("http://192.168.33.50:8200/api/v1/vehicles/single","POST",currentCar);
          if(response[i].end_date<new Date())
            status="Zakończone";
          else 
            status="W trakcie";
          let elem = document.createElement('div');
          if(i!=0)
            elem.style="border-top: 1px solid var(--bs-black);"
          elem.append(tmpl.content.cloneNode(true));
          elem.querySelector("#name").innerHTML=[car.brand,car.model].join(' ');
          elem.querySelector("#daily_cost").innerHTML=["Dzienny koszt:",car.dailyCost].join(' ')
          elem.querySelector("#fuel_consumption").innerHTML=["Spalanie:",car.fuelConsumption].join(' ')
          elem.querySelector("#status").innerHTML=["Status:",status].join("<br>");
          elem.querySelector("#start_date").innerHTML=["Rozpoczęcie wynajmu:",formatDateOrder(new Date(response[i].start_date))].join("<br>")
          elem.querySelector("#rent_cost").innerHTML=["Zapłacona kwota:",car.dailyCost].join("<br>")
          elem.querySelector("#end_date").innerHTML=["Zakończenie wynajmu:",formatDateOrder(new Date(response[i].end_date))].join("<br>")

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

function loadRentedCar(car){
  console.log(car);
  document.getElementById("currentCar").textContent=[car.brand,car.model].join(' ');
  document.getElementById("registrationNumber").textContent=car.registrationNumber;
  document.getElementById("dailyCost").textContent=car.dailyCost;
  document.getElementById("fuelConsumption").textContent=car.fuelConsumption;
  return car.dailyCost;
}