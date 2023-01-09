function getPhoto(brand,model){
  return "cars/"+brand+model+".jpg";
}

// let base64String = "";
// var reader = new FileReader();
// reader.onload = function () {
//         console.log("LOADED")
// }
async function imageUploaded(file) {
  var image = file;

  console.log("next");
  return new Promise((onSuccess, onError) => {
    try {
      const reader = new FileReader() ;
      reader.onload = function(){ onSuccess(this.result.replace("data:", "")
      .replace(/^.+,/, "")) } ;
      reader.readAsDataURL(image);
    } catch(e) {
      onError(e);
    }
  });
}
  

function displayString() {
  console.log("Base64String about to be printed");
  alert(base64String);
}

async function sendPhotos(rentId,photos){
// const fd = new FormData()
  // let prom = new Promise((resolve, reject) =>{

const prom = await imageUploaded(photos)
console.log("A to za wcześniue");
console.log(prom);

  const info={
      rental_id:parseInt(rentId),
      img:prom
  }
  // console.log(JSON.stringify(info))
  // console.log(typeof(bity))
  // console.log(prom)

  return new Promise(async (res, rej) => {                       
    await fetch("http://192.168.33.50:8200/api/v1/rentals/save-image", {method: 'POST' ,mode: 'cors',body: JSON.stringify(info),
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