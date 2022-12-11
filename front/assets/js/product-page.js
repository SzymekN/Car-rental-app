//wywolanie na poczatku 
// if(localStorage.getItem("currentPage"))
//     getFilterCars(localStorage.getItem("currentPage"));
// else
//     getFilterCars();

function getFilterCars(currentPage=0){

    var filters,maxCarsPage=30;
    
    if(localStorage.getItem("filters"))
        filters=localStorage.getItem("filters");
    else
        filters="NULL";
    
    
    var cars=filterCars(filters);
        
   
    //var cars = [filters, '1', '3', '7' ];
    
    let temp, item, a, i;

    temp = document.getElementsByTagName("template")[0];
    item = temp.content.querySelector("div");
    if(cars!=null){
    for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {
      a = document.importNode(item, true);
      let elem=a.querySelectorAll("h3");
      elem[0].textContent=[cars[i].brand,cars[i].model].join(' ');
      let p=a.querySelectorAll("h5");
      p[0].textContent=["Dzienny koszt:",cars[i].dailyCost].join(' ');
      p[1].textContent=["Spalanie:",cars[i].fuelConsumption].join(' ');
      document.getElementById("cardGroup").appendChild(a);
    }}
    else{
        let a=document.createElement("h4");//
        a.style.color='white';
        a.textContent="Brak samochodów dla wybranej kategorii.";
        
        document.getElementById("cardGroup").appendChild(a);
    }
    }


function filterChange(){
    document.getElementById("cardGroup").remove();
    localStorage.setItem("filters","filters");
    document.location.href = "user-rent.html";
}

function filterCars(filters){
    var jsonObj=JSON.parse(localStorage.getItem("allCars"));
    if(jsonObj)
        console.log(jsonObj[0].id);
    return jsonObj;
    //const map = new Map(Object.entries(JSON.parse(json)));
}