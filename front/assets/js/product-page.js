//wywolanie na poczatku 
if(localStorage.getItem("currentPage"))
    getFilterCars(localStorage.getItem("currentPage"));
else
    getFilterCars();

function getFilterCars(currentPage=0){

    var filters,maxCarsPage=30;
    
    if(localStorage.getItem("filters"))
        filters=localStorage.getItem("filters");
    else
        filters="NULL";
    
    
    var cars=filterCars(filters);
        
   
    //var cars = [filters, '1', '3', '7' ];
    //
    let temp, item, a, i;

    temp = document.getElementsByTagName("template")[0];
    item = temp.content.querySelector("div");
    
    for (i = currentPage*maxCarsPage; i < (currentPage*maxCarsPage)+maxCarsPage; i++) {
      a = document.importNode(item, true);
      let elem=a.querySelectorAll("h3");
      elem[0].textContent=cars[i];
      document.getElementById("cardGroup").appendChild(a);
    }
}

function filterChange(){
    document.getElementById("cardGroup").remove();
    localStorage.setItem("filters","filters");
    document.location.href = "user-rent.html";
}

function filterCars(filters){
   
    //const map = new Map(Object.entries(JSON.parse(json)));
}