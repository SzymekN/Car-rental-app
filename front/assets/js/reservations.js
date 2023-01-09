async function pay(){
  if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)!=="employee-checkout.html")
    
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
        console.log(rentInfo)
        Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/rentals/rent-for-user","POST",rentInfo)).then((data) => {
        alert("Pomyślnie zarezerwowano pojazd");
        document.location.href = "employee-rent.html";
    
    }).catch( err => {
        console.log('error: '+ err);
        alert("Rezerwacja samochodu nie powiodła się");

      });
    }}
    else{
      if(document.getElementById('formFileMultiple').files.length!=0){
        Promise.resolve(getInfoWithBody("http://192.168.33.50:8200/api/v1/rentals/self","POST",rentInfo)).then((data) => {
          for(var i=0;i<document.getElementById('formFileMultiple').files.length;i++)
            sendPhotos(data.id,document.getElementById('formFileMultiple').files[i]);
        alert("Pomyślnie zarezerwowano pojazd")
        document.location.href="user-reservations.html";
        }).catch( err => {
          console.log('error: '+ err);
          alert("Rezerwacja samochodu nie powiodła się");
        });
        
      }
      else
        alert("Nie dodano zdjęć!");
    }
  }
  


async function loadRent(currentPage=0){
  var response;
  if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)==="user-reservations.html")
    response=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/rentals/self","GET");
  else
    response=await getInfoWithoutBody("http://192.168.33.50:8200/api/v1/rentals/get-active","GET");
    console.log(response[0]);
    var temp, item, a, i=0,maxCarsPage=30,status;
    var car,elem=0;
    
    if(Object.entries(response).length!=0){
        
      for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {
        if(i<Object.keys(response).length){
          
          const currentCar ={id:response[i].vehicle_id};
          console.log(typeof parseInt(localStorage.getItem("currentCar")))
          var car =await getInfoWithBody("http://192.168.33.50:8200/api/v1/vehicles/single","POST",currentCar);
          var start=new Date(response[i].start_date);
          var end=new Date(response[i].end_date);
          var dateDiff=(end.getTime()-start.getTime())/(1000*3600*24)+1
          if(new Date(response[i].end_date)<new Date())
            status="Zakończone";
          else 
            status="W trakcie";
          let elem = document.createElement('div');
          if(i!=0)
            elem.style="border-top: 1px solid var(--bs-black);"
          
          var paid;
          if(dateDiff*car.dailyCost<0)
            paid=0;
          else
            paid=dateDiff*car.dailyCost;
          elem.append(tmpl.content.cloneNode(true));
          elem.querySelector("#name").innerHTML=[car.brand,car.model].join(' ');
          elem.querySelector("#daily_cost").innerHTML=["Dzienny koszt:",car.dailyCost].join(' ')
          elem.querySelector("#fuel_consumption").innerHTML=["Spalanie:",car.fuelConsumption].join(' ')
          elem.querySelector("#status").innerHTML=["Status:",status].join("<br>");
          elem.querySelector("#start_date").innerHTML=["Rozpoczęcie wynajmu:",formatDateOrder(new Date(response[i].start_date))].join("<br>")
          elem.querySelector("#rent_cost").innerHTML=["Zapłacona kwota:",paid].join("<br>")
          elem.querySelector("#end_date").innerHTML=["Zakończenie wynajmu:",formatDateOrder(new Date(response[i].end_date))].join("<br>")
          elem.querySelector("img").src=getPhoto(car.brand,car.model);
          if(new Date()>=new Date(response[i].end_date)){
            elem.querySelector("#end").disabled=true;
            elem.querySelector("#crash").disabled=true;
          }
          else if(new Date(response[i].start_date)>new Date()){
            elem.querySelector("#end").disabled=true;
            elem.querySelector("#crash").disabled=true;
          }
          // to avoid cannot set properties do disabled
          else{
            elem.querySelector("#end").id=response[i].id;
            elem.querySelector("#crash").id=car.id;
          }
          // border for every other element than first
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

async function endRent(){
  var idVal=localStorage.getItem("currentRentId");
  console.log(idVal);
  for(var i=0;i<document.getElementById('formFileMultiple').files.length;i++)
    sendPhotos(idVal,document.getElementById('formFileMultiple').files[i]);
  await getInfoWithBody("http://192.168.33.50:8200/api/v1/rentals/end","POST",{id:parseInt(idVal)});
  reload();
}

async function reportDamage(){
  var descVal=document.getElementById("description").value;
  var idVal=localStorage.getItem('currentCarId');
  console.log({vehicle_id:parseInt(idVal),description:descVal});
  if(descVal.length!=0){
      if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)==="employee-rent.html")
        await getInfoWithBody("http://192.168.33.50:8200/api/v1/notifications/employee","POST",{vehicle_id:parseInt(idVal),description:descVal});
      else
        await getInfoWithBody("http://192.168.33.50:8200/api/v1/notifications/client","POST",{vehicle_id:parseInt(idVal),description:descVal});
      alert("Pomyślnie wysłano zgłoszenie");
      reload();
  }
  else
      alert("Wiadomość jest pusta")

}