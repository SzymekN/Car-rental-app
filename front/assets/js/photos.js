function getPhoto(brand,model){
    return "cars/"+brand+model+".jpg";
}

let base64String = "";
     
async function imageUploaded(file) {
    var image = file[0];
 
    var reader = new FileReader();
    console.log("next");
     
    return new Promise((resolve, reject) =>{
    reader.onload = function () {
        base64String = reader.result.replace("data:", "")
            .replace(/^.+,/, "");
 
            console.log("To się robi za późno");
          imageBase64Stringsep = base64String;
          
          // alert(imageBase64Stringsep);
          console.log(base64String);
          reader.readAsDataURL(image);
          resolve(base64String)
      }
    });
    prom.then((resolve) => {
      console.log("result: " + resolve);
    });
    
}
    
 
function displayString() {
    console.log("Base64String about to be printed");
    alert(base64String);
}

async function sendPhotos(rentId,photos){
  // const fd = new FormData()
    // let prom = new Promise((resolve, reject) =>{

  let prom = imageUploaded(photos)
  prom.then(resolve =>{
    console.log(resolve);

  });
  console.log("A to za wcześniue");

    const info={
        id:parseInt(rentId),
        img:base64String
    }
    console.log(JSON.stringify(info))
    console.log(typeof(bity))
    console.log(base64String)

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