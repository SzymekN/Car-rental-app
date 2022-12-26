function pay(){
    createNewReservation();
    document.location.href = "user-reservations.html";
    
  }
function createNewReservation(){
    const rentInfo = {
        start_date: localStorage.getItem("startDate"),
        end_date: localStorage.getItem("endDate"),
        pickup_address: localStorage.getItem("pickupAdress"),
        vehicle_id: parseInt(localStorage.getItem("currentCar"))
    }

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
              var cost=loadCar(data);
              return res(cost);
          }).then(res.toString).catch( err => {
              return rej(err);                         
          });                                              
  });
}