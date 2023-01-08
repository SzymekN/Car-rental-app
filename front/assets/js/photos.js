function getPhoto(brand,model){
    return "cars/"+brand+model+".jpg";
}

function sendPhotos(rentId,photos){
    const info={
        id:parseInt(rentId),
        img:photos
    }
    return new Promise(async (res, rej) => {                       
        await fetch(target, {method: httpMethod,mode: 'cors',body: info,
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