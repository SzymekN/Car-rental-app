function getPhoto(brand,model){
  return "cars/"+brand+model+".jpg";
}

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

async function sendPhotos(rentId,photos){
const prom = await imageUploaded(photos)

  const info={
      rental_id:parseInt(rentId),
      img:prom
  }

  var target;
  if(window.location.href.substring(window.location.href.lastIndexOf('/') + 1)==="user-checkout.html")
    target="http://192.168.33.50:8200/api/v1/rentals/save-image-before";
  else
    target="http://192.168.33.50:8200/api/v1/rentals/save-image-after"

  return new Promise(async (res, rej) => {                       
    await fetch(target, {method: 'POST' ,mode: 'cors',body: JSON.stringify(info),
    headers: {
      "Content-Type": "application/json",
      "Authorization":"Bearer "+localStorage.getItem("token")
    }}).then(async (r) => {   
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