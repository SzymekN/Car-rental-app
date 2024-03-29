function getInfoWithoutBody(target,httpMethod){
      return new Promise(async (res, rej) => {                       
        await fetch(target, {method: httpMethod,mode: 'cors',
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

 function getInfoWithBodyWithoutToken(target,httpMethod,httpBody){
  return new Promise(async (res, rej) => {                       
    await fetch(target, {method: httpMethod,mode: 'cors',body: JSON.stringify(httpBody),
    headers: {
      "Content-Type": "application/json"
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

 function   getInfoWithBody(target,httpMethod,httpBody){
        return new Promise(async (res, rej) => {                       
          await fetch(target, {method: httpMethod,mode: 'cors',body: JSON.stringify(httpBody),
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

 function reload(){
  document.location.href=window.location.href.substring(window.location.href.lastIndexOf('/') + 1);
 }